package main

import "fmt"

// Move represents what the user will send to the server
type Move struct {
	Player     `json:"player"`
	Action     `json:"action"`
	Tile       int `json:"tile"`
	TurnNumber int `json:"turn_number"`
}

type LastWinningHand struct {
	Player      `json:"player"`
	WinningTile int   `json:"winning_tile"`
	Hand        []int `json:"hand"`
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
	LastWinningHand   `json:"last_winning_hand"`
	LastWinningTurn   int `json:"last_winning_turn"`
}

func NewGameState() GameState {
	// raw game state: prevailing wind
	gs := GameState{
		PrevailingWind:  0,
		TurnNumber:      0,
		LastWinningTurn: -1,
	}
	gs.resetBoard(P0)
	return gs
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
		remaining, err := ps.Draw(gs.RemainingTiles)
		if err != nil {
			return nil
		}
		gs.RemainingTiles = ps.repairHand(remaining)
		ps.updateInnerGMap()

	case Eat, EatLeft, EatRight:
		ps.Eat(m.Tile, m.Action)

	case Pong:
		ps.Pong(m.Tile)

	case Gong:
		gs.RemainingTiles = ps.Gong(m.Tile, gs.RemainingTiles)
		gs.RemainingTiles = ps.repairHand(gs.RemainingTiles)

	case InnerGong:
		gs.RemainingTiles = ps.InnerGong(m.Tile, gs.RemainingTiles)
		gs.RemainingTiles = ps.repairHand(gs.RemainingTiles)

	case Discard:
		// only player's turn can call discard
		if gs.PlayerTurn != m.Player {
			return fmt.Errorf("not your turn to discard")
		}

		if err := ps.Discard(m.Tile); err != nil {
			return err
		}

		// trigger update of all players
		for k, v := range gs.PlayerMap {
			v.ResetStatus()
			if k != m.Player {
				v.UpdateStatus(m.Tile)
			}
		}

	case Call:
		// move to next player
		if m.Player != gs.Starter { // check for "diao zhng"
			gs.Starter = gs.Starter.next()
			// advance wind if need
			if gs.Starter == P0 {
				gs.PrevailingWind = gs.PrevailingWind.next()
			}
		}

		// record winner
		gs.recordWinner(ps, m)

		// display winning hand only for next turn
		gs.LastWinningTurn = gs.TurnNumber + 1

		// reset everything
		gs.resetBoard(gs.Starter)
	}

	gs.stateTransit(m.Action, m.Player, &m.Tile)

	if err := gs.validateTurn(); err != nil {
		return err
	}

	return nil
}

func (gs *GameState) recordWinner(ps *PlayerState, m Move) {
	// record winner
	whand := LastWinningHand{Hand: ps.Hand, Player: m.Player}
	// winning tile is either last discarded or last drawn
	if gs.LastDiscardedTile == nil {
		whand.WinningTile = ps.LastDrawnTile
	} else {
		whand.WinningTile = *gs.LastDiscardedTile
		// consolidate winners hand
		if whand.WinningTile != -1 {
			whand.Hand = append(whand.Hand, whand.WinningTile)
		}
	}

	for _, d := range ps.Displayed {
		whand.Hand = append(whand.Hand, d...)
	}
	gs.LastWinningHand = whand
}

// resets board to with `start` being the dealer
// only prevailing wind and turn number is not set here
func (gs *GameState) resetBoard(dealer Player) {
	tiles := initSet()
	pMap := make(map[Player]*PlayerState)
	for _, p := range []Player{P0, P1, P2, P3} {
		handSize := 13
		if p == dealer {
			handSize = 14
		}

		ps, leftover := NewPlayerState(handSize, tiles)
		pMap[p] = ps
		tiles = leftover
	}

	gs.Starter = dealer
	gs.PlayerTurn = dealer
	gs.IsTransitioning = false
	gs.DiscardedTiles = []int{}
	gs.PlayerMap = pMap
	gs.RemainingTiles = tiles
	gs.LastDiscardedTile = nil
}

// Reflects state of the game
// IsTransitioning: a discard just happened
// * only valid moves to players are draw/eat/pong/gong/call
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

// potentially expensive but only run after a move is performed
// ensures no duplicate tiles and no shortage/surplus
func (gs *GameState) validateTurn() error {
	checkMap := make(map[int]bool)
	for _, t := range gs.RemainingTiles {
		if _, ok := checkMap[t]; ok {
			return fmt.Errorf("game state invalid: duplicate @ remaining tiles")
		}
		checkMap[t] = true
	}

	for _, t := range gs.DiscardedTiles {
		if _, ok := checkMap[t]; ok {
			return fmt.Errorf("game state invalid: duplicate @ discarded tile")
		}
		checkMap[t] = true
	}

	for _, v := range gs.PlayerMap {
		for _, t := range v.Hand {
			if _, ok := checkMap[t]; ok {
				return fmt.Errorf("game state invalid")
			}
			checkMap[t] = true
		}
		for _, s := range v.Displayed {
			for _, t := range s {
				if _, ok := checkMap[t]; ok {
					return fmt.Errorf("game state invalid")
				}
				checkMap[t] = true
			}
		}
	}

	if gs.LastDiscardedTile != nil {
		if _, ok := checkMap[*gs.LastDiscardedTile]; ok {
			return fmt.Errorf("game state invalid: duplicate @ last discarded tile")
		}
		checkMap[*gs.LastDiscardedTile] = true
	}

	if len(checkMap) != 148 {
		return fmt.Errorf("game state invalid: insufficent tiles")
	}

	return nil
}
