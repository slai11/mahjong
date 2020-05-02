package main

import (
	"fmt"
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

// Move represents what the user will send to the server
type Move struct {
	Player     `json:"player"`
	Action     `json:"action"`
	Tile       `json:"tile"`
	TurnNumber int `json:"turn_number"`
}

// GameState represents the mahjong table
type GameState struct {
	// anchor player, prevailingWind changes when PlayerTurn == starter
	Starter        Player `json:"starter"`
	PrevailingWind `json:"prevailing_wind"`

	// next player who can make a move, can jump
	PlayerTurn Player `json:"player_turn"`

	// To ensure correct order of moves, no one can move
	// if two players click Call at the same time, the first request received by
	// the server will hold
	TurnNumber int `json:"turn_number"`

	// transitioning between players
	// possible actions
	// if false, either a discard or Call
	// discard will set val to true
	//
	// if true, either a draw, eat/pong/gong or Call
	// draw/eat/pong/gong will move to next player and set to false
	IsTransitioning bool `json:"is_transitioning"`

	PlayerMap         map[Player]*PlayerState `json:"player_map"`
	DiscardedTiles    []Tile                  `json:"discarded_tiles"`
	RemainingTiles    []Tile                  `json:"remaining_tiles"`
	LastDiscardedTile *Tile                   `json:"last_discarded_tile"`
}

func NewGameState() GameState {
	tiles := initSet()
	pMap := make(map[Player]*PlayerState)
	for _, p := range []Player{P0, P1, P2, P3} {
		handSize := 13
		if p == P0 {
			handSize = 14
		}

		ps, leftover := NewPlayerState(handSize, tiles)
		pMap[p] = ps
		tiles = leftover
	}

	return GameState{
		RemainingTiles:    tiles,
		PlayerMap:         pMap,
		Starter:           P0,
		PlayerTurn:        P0,
		PrevailingWind:    0,
		IsTransitioning:   false,
		DiscardedTiles:    []Tile{},
		LastDiscardedTile: nil,
	}
}

// Draw -> {Discard, Gong}
// Discard -> {Draw, Eat, Pong, Gong, Call}
func (gs *GameState) NextTurn(m Move) error {
	if gs.TurnNumber != m.TurnNumber {
		return fmt.Errorf("turn over")
	}

	ps, _ := gs.PlayerMap[m.Player]
	switch m.Action {
	case Draw:
		gs.RemainingTiles = ps.Draw(gs.RemainingTiles)
		gs.IsTransitioning = false
		if gs.LastDiscardedTile != nil {
			gs.DiscardedTiles = append(gs.DiscardedTiles, *gs.LastDiscardedTile)
			gs.LastDiscardedTile = nil
		}

	case Eat, EatLeft, EatRight:
		ps.Eat(m.Tile, m.Action)
		gs.LastDiscardedTile = nil
		gs.IsTransitioning = false

	case Pong:
		ps.Pong(m.Tile)
		gs.LastDiscardedTile = nil
		gs.IsTransitioning = false

	case Gong:
		gs.RemainingTiles = ps.Gong(m.Tile, gs.RemainingTiles)
		gs.LastDiscardedTile = nil
		gs.IsTransitioning = false

	case InnerGong:
		gs.RemainingTiles = ps.InnerGong(m.Tile.Suit, m.Tile.Value, gs.RemainingTiles)
		gs.LastDiscardedTile = nil
		gs.IsTransitioning = false

	case Discard:
		ps.Discard(m.Tile.ID)
		// trigger update of all players
		for k, v := range gs.PlayerMap {
			if k != m.Player {
				v.UpdateStatus(m.Tile)
			}
		}

		if gs.LastDiscardedTile != nil {
			gs.DiscardedTiles = append(gs.DiscardedTiles, *gs.LastDiscardedTile)
		}
		gs.LastDiscardedTile = &m.Tile
		gs.IsTransitioning = true

	case Call:
		// reset everything
		// move to next player
		if m.Player != gs.Starter {
			gs.Starter = gs.Starter.next()
			if gs.Starter == P0 {
				gs.PrevailingWind.next()
			}
		}

	}

	gs.PlayerTurn = m.Player
	gs.TurnNumber += 1
	return nil
}
