package main

type PlayerState struct {

	// displayed tiles on board
	Displayed [][]int `json:"displayed"`

	// starts off with 13/14
	// hand should only be visible to player
	Hand []int `json:"hand"`

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

	ps := PlayerState{Hand: hand, Displayed: [][]int{}, InnerGongMap: make(map[int]int)}
	tiles = ps.repairHand(tiles)

	return &ps, tiles
}

// Draw is always folloed by repair since flower/animal must be opened
func (p *PlayerState) Draw(tiles []int) []int {
	drawn := tiles[0]
	tiles = tiles[1:]
	p.Hand = append(p.Hand, drawn)

	tiles = p.repairHand(tiles)
	p.updateInnerGMap()
	return tiles
}

func (p *PlayerState) Discard(id int) {
	for i, t := range p.Hand {
		if t == id {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
		}
	}
}

// can eat middle or eat side
func (p *PlayerState) Eat(t int, side Action) {
	if !(p.CanEat || p.CanEatRight || p.CanEatLeft) {
		return
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
	// bug!
	i := 0
	ate := false
	for _, h := range p.Hand {
		hTile := TileList[h]
		if tile.Suit == hTile.Suit {
			if !eatenFirst && first == hTile.Value {
				triplet = append(triplet, h)
				eatenFirst = true
				ate = true
			} else if !eatenSecond && second == hTile.Value {
				triplet = append(triplet, h)
				eatenSecond = true
				ate = true
			}
		}
		if !ate {
			p.Hand[i] = h
			i++
		}
		ate = false
	}
	p.Hand = p.Hand[:i]

	if len(triplet) != 3 {
		panic("NO TRIPLET")
	}

	p.Displayed = append(p.Displayed, triplet)
}

func (p *PlayerState) Pong(tile int) {
	if !p.CanPong {
		return
	}

	newHand := make([]int, 0, len(p.Hand)-2)
	triplet := []int{tile}
	for _, h := range p.Hand {
		if TileList[tile].Suit == TileList[h].Suit && TileList[tile].Value == TileList[h].Value {
			triplet = append(triplet, h)
		} else {
			newHand = append(newHand, h)
		}
	}

	if len(triplet) == 3 {
		p.Hand = newHand
		p.Displayed = append(p.Displayed, triplet)
	}
}

func (p *PlayerState) Gong(tile int, tiles []int) []int {
	if !p.CanGong {
		return tiles
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

	p.Hand = newHand
	p.Displayed = append(p.Displayed, triplet)

	tiles = p.Draw(tiles) // always draw tile after a gong
	return tiles
}

func (p *PlayerState) InnerGong(t int, tiles []int) []int {
	tile := TileList[t]
	s := tile.Suit
	v := tile.Value
	k := int(s)*10 + v
	if mv, ok := p.InnerGongMap[k]; !ok || mv != v {
		return tiles
	}

	triplet := []int{}
	newHand := make([]int, 0, len(p.Hand)-4)
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

	tiles = p.Draw(tiles)
	delete(p.InnerGongMap, k)
	return tiles
}

func (p *PlayerState) ResetStatus() {
	p.CanPong = false
	p.CanGong = false
	p.CanEat = false
	p.CanEatLeft = false
	p.CanEatRight = false
}

// this happens after drawing
func (p *PlayerState) UpdateStatus(t int) {
	tile := TileList[t]
	matching := 1
	var hasTwoBefore, hasBefore, hasAfter, hasTwoAfter bool
	for _, h := range p.Hand {
		hTile := TileList[h]
		if tile.Suit == hTile.Suit {
			if tile.Value == hTile.Value {
				matching += 1
				continue
			}

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

	if matching == 3 {
		p.CanPong = true
	}

	if matching == 4 {
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
func (p *PlayerState) repairHand(tiles []int) []int {
	// lots of looping but realistically there are only 12 flower/animals

	for _, t := range p.Hand {
		tile := TileList[t]
		if tile.Suit == Flower || tile.Suit == Animal {
			p.Displayed = append(p.Displayed, []int{t})

			// draw 1 new tile
			x := tiles[0]
			tiles = tiles[1:]
			xt := TileList[x]
			// keep replenishing until not flower or animal
			for xt.Suit == Flower || xt.Suit == Animal {
				p.Displayed = append(p.Displayed, []int{x})

				// draw another
				x = tiles[0]
				tiles = tiles[1:]
				xt = TileList[x]
			}

			p.Hand = append(p.Hand, x)
		}
	}

	for _, t := range p.Displayed {
		rmId := -1
		for i := range p.Hand {
			if p.Hand[i] == t[0] { //TODO hack here since flower/animals would be in lenght-1 slices
				rmId = i
				break
			}
		}

		if rmId != -1 {
			// important to preserve order
			p.Hand = append(p.Hand[:rmId], p.Hand[rmId+1:]...)
		}
	}
	return tiles
}

// only 9 possible values and 7 suits
// Suit * 10 + Value = tile identifier
func (p *PlayerState) updateInnerGMap() {
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

	// if a Displayed triple exist
	for _, s := range p.Displayed {
		if len(s) == 3 {
			s0 := TileList[s[0]]
			s1 := TileList[s[1]]
			s2 := TileList[s[2]]
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
}
