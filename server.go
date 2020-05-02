package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Server http.Server
	games  map[string]*Game
}

func (s *Server) Start(games map[string]*Game) error {
	r := mux.NewRouter()
	r.HandleFunc("/move", s.handleMove)
	r.HandleFunc("/game_state", s.handleGetState)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	s.Server = http.Server{
		Addr:    ":80",
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(r),
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
		gh = NewGame(gameID)
		s.games[gameID] = gh
	}

	writeJSON(rw, gh)
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

	err := gh.ProcessMove(request.Move)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	writeJSON(rw, gh)
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
