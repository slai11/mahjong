package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Server http.Server
	games  map[string]*GameHolder
	mu     sync.RWMutex
}

func (s *Server) Start(games map[string]*GameHolder) error {
	r := mux.NewRouter()

	// not the most REST-compliant but its straightforward
	r.Use(CORS)
	r.HandleFunc("/move", s.handleMove)
	r.HandleFunc("/game_state", s.handleGetState)
	r.HandleFunc("/player", s.handlePlayerSelect)

	s.Server = http.Server{
		Addr:    ":80",
		Handler: r,
	}

	s.games = games

	go func() {
		for range time.Tick(30 * time.Minute) {
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

	s.mu.RLock()
	defer s.mu.RUnlock()
	gh, ok := s.games[gameID]
	if !ok {
		log.WithFields(log.Fields{
			"method": "handleGetState",
		}).Info("Not found, creating new game: ", gameID)

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

	// read lock on the gamesMap but internally, "move" will acquire its own
	// write lock for updating game state
	s.mu.RLock()
	defer s.mu.RUnlock()
	gh, ok := s.games[request.GameID]
	if !ok {
		log.WithFields(log.Fields{
			"method": "handleMove",
		}).Error("Game does not exist: ", request.GameID)

		http.Error(rw, "No such game", 404)
		return
	}

	err := gh.Update(request.Move)
	if err != nil {
		log.WithFields(log.Fields{
			"method": "handleMove",
		}).Error("Error with update: ", err)

		http.Error(rw, err.Error(), 400)
		return
	}

	writeJSON(rw, gh.Get())
}

// GET player
// not locking here because it is unlikely to be deleted since its new
func (s *Server) handlePlayerSelect(rw http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	gameID := q.Get("game_id")

	gh, ok := s.games[gameID]
	if !ok {
		log.WithFields(log.Fields{
			"method": "handlePlayerSelect",
		}).Error("Game not found: ", gameID)

		http.Error(rw, "No such game", 404)
		return
	}

	playerID, err := gh.SelectPlayer()
	if err != nil {
		log.WithFields(log.Fields{
			"method": "handlePlayerSelect",
		}).Error(err)

		http.Error(rw, err.Error(), 400)
		return
	}

	writeJSON(rw, map[string]int{"assigned_number": playerID})
}

func (s *Server) cleanupOldGames() {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.WithFields(log.Fields{
		"method": "cleanupOldGames",
	}).Info("Locking map to clean up. Number of games: ", len(s.games))

	for k, v := range s.games {
		// game inactive for an hour
		if v.g.UpdatedAt.Add(1 * time.Hour).Before(time.Now()) {
			log.WithFields(log.Fields{
				"method": "cleanupOldGames",
			}).Warn("Deleting inactive game: ", k)

			delete(s.games, k)
		}
	}
	log.WithFields(log.Fields{
		"method": "cleanupOldGames",
	}).Info("Done with clean up")
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

// Referenced https://asanchez.dev/blog/cors-golang-options/
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}
