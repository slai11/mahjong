package main

import "fmt"

type PlayerState struct {

	// displayed tiles on board
	Displayed [][]int `json:"displayed"`

	// starts off with 13/14
	// hand should only be visible to player
	Hand []int `json:"hand"`

	// most recent tile that entered hand from drawing
	// or repairing
	LastDrawnTile int `json:"last_drawn_tile"`

	CanEat      bool `json:"can_eat"`
	CanEatLeft  bool `json:"can_eat_left"`
	CanEatRight bool `json:"can_eat_right"`
	CanPong     bool `json:"can_pong"`
	CanGong     bool `json:"can_gong"`

	// only accounts for Hand + Displayed triplets
	InnerGongMap map[int]int `json:"inner_gong_map"`
}

func NewPlayerState(size int, tiles []int) (*PlayerState, []int) {
	hand := tiles[0:size]
	tiles = tiles[size:]

	ps := PlayerState{
		Hand:         hand,
		Displayed:    [][]int{},
		InnerGongMap: make(map[int]int),
	}

	tiles = ps.RepairHand(tiles)
	ps.updateInnerGMap()

	// starting hand does not have a last drawn tile
	ps.LastDrawnTile = -1
	return &ps, tiles
}

// Draw simply pops the first item of the remaining tiles into hand
func (p *PlayerState) Draw(tiles []int) ([]int, error) {
	if len(tiles) < 1 {
		return nil, fmt.Errorf("[PlayerState.Draw]insufficient tiles")
	}
	drawn := tiles[0]
	tiles = tiles[1:]
	p.Hand = append(p.Hand, drawn)
	p.LastDrawnTile = drawn
	return tiles, nil
}

// Discard finds tile and deletes it from hand
func (p *PlayerState) Discard(id int) error {
	for i, t := range p.Hand {
		if t == id {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("[PlayerState.Discard]tile not found")
}

// can eat middle or eat side
func (p *PlayerState) Eat(t int, side Action) error {
	if !(p.CanEat || p.CanEatRight || p.CanEatLeft) {
		return fmt.Errorf("[PlayerState.Eat] eating is not allowed")
	}

	tile := TileList[t]
	var first, second int
	switch side {
	case Eat:
		first = tile.Value - 1
		second = tile.Value + 1
	case EatLeft:
		first = tile.Value + 1
		second = tile.Value + 2
	case EatRight:
		first = tile.Value - 1
		second = tile.Value - 2
	}

	var eatenFirst, eatenSecond bool
	triplet := []int{t}
	newHand := make([]int, 0, len(p.Hand)-2)
	for _, h := range p.Hand {
		hTile := TileList[h]
		if tile.Suit == hTile.Suit {
			if !eatenFirst && first == hTile.Value {
				triplet = append(triplet, h)
				eatenFirst = true
				continue
			} else if !eatenSecond && second == hTile.Value {
				triplet = append(triplet, h)
				eatenSecond = true
				continue
			}
		}
		newHand = append(newHand, h)
	}

	if len(triplet) != 3 {
		return fmt.Errorf("[PlayerState.Eat] triplet does not exist")
	}

	p.Hand = newHand
	p.Displayed = append(p.Displayed, triplet)
	return nil
}

func (p *PlayerState) Pong(t int) error {
	if !p.CanPong {
		return fmt.Errorf("[PlayerState.Pong] pong is not allowed")
	}

	tile := TileList[t]
	newHand := make([]int, 0, len(p.Hand)-2)
	triplet := []int{t}
	for _, h := range p.Hand {
		if tile.Suit == TileList[h].Suit && tile.Value == TileList[h].Value && len(triplet) < 3 {
			triplet = append(triplet, h)
		} else {
			newHand = append(newHand, h)
		}
	}

	if len(triplet) != 3 {
		return fmt.Errorf("[PlayerState.Pong] triplet does not exist")
	}

	p.Hand = newHand
	p.Displayed = append(p.Displayed, triplet)
	return nil
}

func (p *PlayerState) Gong(tile int) error {
	if !p.CanGong {
		return fmt.Errorf("[PlayerState.Gong] gong is not allowed")
	}

	newHand := make([]int, 0, len(p.Hand)-3)
	triplet := []int{tile}
	for _, h := range p.Hand {
		if TileList[tile].Suit == TileList[h].Suit && TileList[tile].Value == TileList[h].Value {
			triplet = append(triplet, h)
		} else {
			newHand = append(newHand, h)
		}
	}

	if len(triplet) != 4 {
		return fmt.Errorf("[PlayerState.Gong] triplet does not exist")
	}

	p.Hand = newHand
	p.Displayed = append(p.Displayed, triplet)

	return nil
}

func (p *PlayerState) InnerGong(t int, tiles []int) []int {
	tile := TileList[t]
	s := tile.Suit
	v := tile.Value
	k := int(s)*10 + v
	if mv, ok := p.InnerGongMap[k]; !ok || mv != 4 {
		return tiles
	}

	triplet := []int{}
	newHand := make([]int, 0)
	for _, h := range p.Hand {
		tile := TileList[h]
		if s == tile.Suit && v == tile.Value {
			triplet = append(triplet, h)
		} else {
			newHand = append(newHand, h)
		}
	}
	p.Hand = newHand
	p.Displayed = append(p.Displayed, triplet)

	tiles, _ = p.Draw(tiles)
	delete(p.InnerGongMap, k)
	return tiles
}

func (p *PlayerState) ResetStatus() {
	p.CanPong = false
	p.CanGong = false
	p.CanEat = false
	p.CanEatLeft = false
	p.CanEatRight = false
	p.LastDrawnTile = -1
}

// UpdateStatus happens after a discard to set limits on possible actions
func (p *PlayerState) UpdateStatus(t int) {
	tile := TileList[t]
	matching := 1
	var hasTwoBefore, hasBefore, hasAfter, hasTwoAfter bool
	for _, h := range p.Hand {
		hTile := TileList[h]
		if tile.Suit == hTile.Suit {
			// pong/gong
			if tile.Value == hTile.Value {
				matching += 1
				continue
			}

			// eat/eatleft/eatright
			switch tile.Suit {
			case Bamboo, Coin, Number:
				if tile.Value == hTile.Value+1 {
					hasBefore = true
				} else if tile.Value == hTile.Value-1 {
					hasAfter = true
				} else if tile.Value == hTile.Value-2 {
					hasTwoAfter = true
				} else if tile.Value == hTile.Value+2 {
					hasTwoBefore = true
				}
			}
		}
	}

	switch matching {
	case 3:
		p.CanPong = true
	case 4:
		p.CanPong = true
		p.CanGong = true
	}

	if hasBefore && hasAfter {
		p.CanEat = true
	}

	if hasTwoBefore && hasBefore {
		p.CanEatRight = true
	}

	if hasTwoAfter && hasAfter {
		p.CanEatLeft = true
	}
}

// for every flower/animal card in Hand, draw 1 card
// if the drawn card is flower/aniimal, draw another
// then move all flower/animal in Hand to Displayed stack
func (p *PlayerState) RepairHand(tiles []int) []int {
	hand := make([]int, 0, 14) // max 14 anyway
	for _, t := range p.Hand {
		tile := TileList[t]

		if tile.Suit != Flower && tile.Suit != Animal {
			hand = append(hand, t)
			continue
		}

		p.Displayed = append(p.Displayed, []int{t})

		// draw 1 new tile
		x := tiles[0]
		tiles = tiles[1:]
		xt := TileList[x]

		// keep replenishing until not flower or animal
		// lots of looping but realistically there are only 12 flower/animals
		for xt.Suit == Flower || xt.Suit == Animal {
			p.Displayed = append(p.Displayed, []int{x})

			// draw another
			x = tiles[0]
			tiles = tiles[1:]
			xt = TileList[x]
		}

		// newly drawn non-flower/animal tile to hand
		hand = append(hand, x)
		p.LastDrawnTile = x
	}

	p.Hand = hand

	return tiles
}

// only 9 possible values and 7 suits
// Suit * 10 + Value = tile identifier
func (p *PlayerState) updateInnerGMap() {
	p.InnerGongMap = make(map[int]int)
	// consolidate hand
	for _, t := range p.Hand {
		tile := TileList[t]
		k := int(tile.Suit)*10 + tile.Value
		count, ok := p.InnerGongMap[k]
		if !ok {
			p.InnerGongMap[k] = 1
		} else {
			p.InnerGongMap[k] = count + 1
		}
	}

	// displayed triplets can be upgraded to gong
	for _, s := range p.Displayed {
		if len(s) != 3 {
			continue
		}

		s0 := TileList[s[0]]
		s1 := TileList[s[1]]
		s2 := TileList[s[2]]

		// all 3 displayed tile must be the same to qualify as potential gang
		if s0.Suit == s1.Suit && s0.Suit == s2.Suit && s0.Value == s1.Value && s0.Value == s2.Value {
			k := int(s0.Suit)*10 + s0.Value
			count, ok := p.InnerGongMap[k]
			if !ok {
				p.InnerGongMap[k] = 3
			} else {
				p.InnerGongMap[k] = count + 3
			}
		}
	}
}
