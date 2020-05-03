package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Server http.Server
	games  map[string]*GameHolder
}

func (s *Server) Start(games map[string]*GameHolder) error {
	r := mux.NewRouter()

	// not the most REST-compliant but its straightforward
	r.Use(CORS)

	r.HandleFunc("/move", s.handleMove)
	r.HandleFunc("/game_state", s.handleGetState)
	r.HandleFunc("/player_select", s.handlePlayerSelect)
	//headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	//originsOk := handlers.AllowedOrigins([]string{"*"})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	s.Server = http.Server{
		Addr: ":80",
		//Handler: handlers.CORS(originsOk, headersOk, methodsOk)(r),
		Handler: r,
	}

	s.games = games

	go func() {
		for range time.Tick(86400 * time.Minute) {
			s.cleanupOldGames()
		}
	}()

	return s.Server.ListenAndServe()
}

// GET game_state
// creates a new game state if 404
func (s *Server) handleGetState(rw http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	gameID := q.Get("game_id")

	gh, ok := s.games[gameID]
	if !ok {
		// TODO set number of games that we can concurrently manage
		gh = &GameHolder{g: NewGame(gameID)}
		s.games[gameID] = gh
	}

	writeJSON(rw, gh.Get())
}

// POST /move
func (s *Server) handleMove(rw http.ResponseWriter, req *http.Request) {
	var request struct {
		GameID string `json:"game_id"`
		Index  int    `json:"index"`
		Move   `json:"move"`
	}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(rw, "Error decoding", 400)
		return
	}

	gh, ok := s.games[request.GameID]
	if !ok {
		http.Error(rw, "No such game", 404)
		return
	}

	err := gh.Update(request.Move)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	writeJSON(rw, gh)
}

// POST  player_select
func (s *Server) handlePlayerSelect(rw http.ResponseWriter, req *http.Request) {
	var request struct {
		GameID    string `json:"game_id"`
		Index     int    `json:"index"`
		Selection int    `json:"selection"`
	}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(rw, "Error decoding", 400)
		return
	}

	gh, ok := s.games[request.GameID]
	if !ok {
		http.Error(rw, "No such game", 404)
		return
	}

	err := gh.SelectPlayer(request.Selection)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	writeJSON(rw, gh.Get())
}

func (s *Server) cleanupOldGames() {
	for _, v := range s.games {
		fmt.Println(v)
		//if v.UpdatedAt()
	}
}

func writeJSON(rw http.ResponseWriter, resp interface{}) {
	j, err := json.Marshal(resp)
	if err != nil {
		http.Error(rw, "unable to marshal response: "+err.Error(), 500)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			fmt.Println("options ok")
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}
