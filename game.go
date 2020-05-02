package main

import (
	"time"
)

// Game contains detail of gamestate and metadata like identifiers
type Game struct {
	GameState `json:"game_state"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewGame(id string) *Game {
	return &Game{
		GameState: NewGameState(),
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (g *Game) ProcessMove(m Move) error {
	g.UpdatedAt = time.Now()
	return g.GameState.NextTurn(m)
}
