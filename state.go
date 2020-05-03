package main

import "fmt"

// Move represents what the user will send to the server
type Move struct {
	Player     `json:"player"`
	Action     `json:"action"`
	Tile       int `json:"tile"`
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

	IsTransitioning bool `json:"is_transitioning"`

	PlayerMap         map[Player]*PlayerState `json:"player_map"`
	DiscardedTiles    []int                   `json:"discarded_tiles"`
	RemainingTiles    []int                   `json:"remaining_tiles"`
	LastDiscardedTile *int                    `json:"last_discarded_tile"`
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
		DiscardedTiles:    []int{},
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

	case Eat, EatLeft, EatRight:
		ps.Eat(m.Tile, m.Action)

	case Pong:
		ps.Pong(m.Tile)

	case Gong:
		gs.RemainingTiles = ps.Gong(m.Tile, gs.RemainingTiles)

	case InnerGong:
		gs.RemainingTiles = ps.InnerGong(m.Tile, gs.RemainingTiles)

	case Discard:
		// only player's turn can call discard
		if gs.PlayerTurn != m.Player {
			return fmt.Errorf("not your turn to discard")
		}
		ps.Discard(m.Tile)
		// trigger update of all players
		for k, v := range gs.PlayerMap {
			v.ResetStatus()
			if k != m.Player {
				v.UpdateStatus(m.Tile)
			}
		}

	case Call:
		// reset everything
		// move to next player
		if m.Player != gs.Starter { // check for "diao zhng"
			gs.Starter = gs.Starter.next()
			if gs.Starter == P0 {
				gs.PrevailingWind.next()
			}
		}
	}

	gs.stateTransit(m.Action, m.Player, &m.Tile)

	return nil
}

// Reflects state of the game
// IsTransitioning: a discard just happened
// * only valid moves to players are draw/eat/pong/gong/call
//
// !IsTransitioning: a player just took a tile either by:
// draw/eat/ping/gong.
// * valid next moves are discard/inner_gong/call
func (gs *GameState) stateTransit(action Action, player Player, tile *int) {
	switch action {
	case Eat, EatLeft, EatRight, Pong, Gong, InnerGong:
		gs.LastDiscardedTile = nil
		gs.IsTransitioning = false
		gs.PlayerTurn = player

	case Draw:
		gs.IsTransitioning = false
		if gs.LastDiscardedTile != nil {
			gs.DiscardedTiles = append(gs.DiscardedTiles, *gs.LastDiscardedTile)
			gs.LastDiscardedTile = nil
		}
		gs.PlayerTurn = player

	case Discard:
		gs.IsTransitioning = true
		if gs.LastDiscardedTile != nil {
			gs.DiscardedTiles = append(gs.DiscardedTiles, *gs.LastDiscardedTile)
		}
		gs.LastDiscardedTile = tile
		gs.PlayerTurn = player.next()
	}

	// ensures turn order correctness
	gs.TurnNumber += 1
}
