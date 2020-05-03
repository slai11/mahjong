package main

import (
	"fmt"
	"sync"
	"time"
)

// Game contains detail of gamestate and metadata like identifiers
type Game struct {
	GameState        `json:"game_state"`
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	PlayersAvailable map[int]bool
}

func NewGame(id string) *Game {
	return &Game{
		GameState: NewGameState(),
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		PlayersAvailable: map[int]bool{
			0: true,
			1: true,
			2: true,
			3: true,
		},
	}
}

func (g *Game) ProcessMove(m Move) error {
	g.UpdatedAt = time.Now()
	return g.GameState.NextTurn(m)
}

func (g *Game) ProcessSelection(m int) error {
	g.UpdatedAt = time.Now()
	avail, ok := g.PlayersAvailable[m]
	if !ok || !avail {
		return fmt.Errorf("Player Not Available")
	}

	g.PlayersAvailable[m] = false
	return nil
}

// context-free holder that prevents access when writing
type GameHolder struct {
	g  *Game
	mu sync.RWMutex
}

func (gh *GameHolder) Update(m Move) error {
	gh.mu.Lock()
	defer gh.mu.Unlock()
	return gh.g.ProcessMove(m)
}

func (gh *GameHolder) Get() *Game {
	gh.mu.RLock()
	defer gh.mu.RUnlock()
	return gh.g
}

func (gh *GameHolder) SelectPlayer(selection int) error {
	gh.mu.Lock()
	defer gh.mu.Unlock()
	return gh.g.ProcessSelection(selection)
}
