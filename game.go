package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Game contains detail of gamestate and metadata like identifiers
type Game struct {
	GameState        `json:"game_state"`
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	PlayersAvailable []int     `json:"players_available"`
}

func NewGame(id string) *Game {
	p := []int{0, 1, 2, 3}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(p), func(i, j int) { p[i], p[j] = p[j], p[i] })
	return &Game{
		GameState:        NewGameState(),
		ID:               id,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		PlayersAvailable: p,
	}
}

func (g *Game) ProcessMove(m Move) error {
	g.UpdatedAt = time.Now()
	return g.GameState.NextTurn(m)
}

func (g *Game) ProcessSelection() (int, error) {
	g.UpdatedAt = time.Now()

	if len(g.PlayersAvailable) == 0 {
		return -1, fmt.Errorf("Player Not Available")
	}
	pos := g.PlayersAvailable[0]
	g.PlayersAvailable = g.PlayersAvailable[1:]
	return pos, nil
}

// context-free holder that prevents access when writing
type GameHolder struct {
	g  *Game
	mu sync.Mutex
}

func (gh *GameHolder) Update(m Move) error {
	gh.mu.Lock()
	defer gh.mu.Unlock()
	return gh.g.ProcessMove(m)
}

func (gh *GameHolder) Get() *Game {
	gh.mu.Lock()
	defer gh.mu.Unlock()
	return gh.g
}

func (gh *GameHolder) SelectPlayer() (int, error) {
	gh.mu.Lock()
	defer gh.mu.Unlock()
	return gh.g.ProcessSelection()
}
