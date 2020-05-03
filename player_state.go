package main

type PlayerState struct {

	// displayed tiles on board
	Displayed [][]Tile `json:"displayed"`

	// starts off with 13/14
	// hand should only be visible to player
	Hand []Tile `json:"hand"`

	CanEat      bool `json:"can_eat"`
	CanEatLeft  bool `json:"can_eat_left"`
	CanEatRight bool `json:"can_eat_right"`
	CanPong     bool `json:"can_pong"`
	CanGong     bool `json:"can_gong"`

	// only accounts for Hand + Displayed triplets
	InnerGongMap map[int]int `json:"inner_gong_map"`
}

func NewPlayerState(size int, tiles []Tile) (*PlayerState, []Tile) {
	hand := tiles[0:size]
	tiles = tiles[size:]

	ps := PlayerState{Hand: hand, Displayed: [][]Tile{}, InnerGongMap: make(map[int]int)}
	tiles = ps.repairHand(tiles)

	return &ps, tiles
}

// Draw is always folloed by repair since flower/animal must be opened
func (p *PlayerState) Draw(tiles []Tile) []Tile {
	drawn := tiles[0]
	tiles = tiles[1:]
	p.Hand = append(p.Hand, drawn)

	tiles = p.repairHand(tiles)
	p.updateInnerGMap()
	return tiles
}

func (p *PlayerState) Discard(id int) {
	for i, t := range p.Hand {
		if t.ID == id {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
		}
	}
}

// can eat middle or eat side
func (p *PlayerState) Eat(tile Tile, side Action) {
	if !(p.CanEat || p.CanEatRight || p.CanEatLeft) {
		return
	}

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
	triplet := []Tile{tile}
	// bug!
	i := 0
	ate := false
	for _, h := range p.Hand {
		if tile.Suit == h.Suit {
			if !eatenFirst && first == h.Value {
				triplet = append(triplet, h)
				eatenFirst = true
				ate = true
			} else if !eatenSecond && second == h.Value {
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

func (p *PlayerState) Pong(tile Tile) {
	if !p.CanPong {
		return
	}

	count := 1
	triplet := []Tile{tile}
	for i, h := range p.Hand {
		if count == 3 {
			break
		}

		if tile.Suit == h.Suit && tile.Value == h.Value {
			triplet = append(triplet, h)
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			count += 1
		}
	}
	p.Displayed = append(p.Displayed, triplet)
}

func (p *PlayerState) Gong(tile Tile, tiles []Tile) []Tile {
	if !p.CanPong {
		return tiles
	}

	count := 1
	triplet := []Tile{tile}
	for i, h := range p.Hand {
		if count == 4 {
			break
		}

		if tile.Suit == h.Suit && tile.Value == h.Value {
			triplet = append(triplet, h)
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			count += 1
		}
	}
	p.Displayed = append(p.Displayed, triplet)

	tiles = p.Draw(tiles)
	return tiles
}

func (p *PlayerState) InnerGong(s Suit, v int, tiles []Tile) []Tile {
	k := int(s)*10 + v
	if mv, ok := p.InnerGongMap[k]; !ok || mv != v {
		return tiles
	}

	count := 0
	triplet := []Tile{}
	for i, h := range p.Hand {
		if count == 4 {
			break
		}

		if s == h.Suit && v == h.Value {
			triplet = append(triplet, h)
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			count += 1
		}
	}
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
func (p *PlayerState) UpdateStatus(tile Tile) {
	matching := 1
	var hasTwoBefore, hasBefore, hasAfter, hasTwoAfter bool
	for _, t := range p.Hand {
		if tile.Suit == t.Suit {
			if tile.Value == t.Value {
				matching += 1
				continue
			}

			switch tile.Suit {
			case Bamboo, Coin, Number:
				if tile.Value == t.Value+1 {
					hasBefore = true
				} else if tile.Value == t.Value-1 {
					hasAfter = true
				} else if tile.Value == t.Value-2 {
					hasTwoAfter = true
				} else if tile.Value == t.Value+2 {
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
func (p *PlayerState) repairHand(tiles []Tile) []Tile {
	// lots of looping but realistically there are only 12 flower/animals

	for _, t := range p.Hand {
		if t.Suit == Flower || t.Suit == Animal {
			p.Displayed = append(p.Displayed, []Tile{t})

			// draw 1 new tile
			x := tiles[0]
			tiles = tiles[1:]
			// keep replenishing until not flower or animal
			for x.Suit == Flower || x.Suit == Animal {
				p.Displayed = append(p.Displayed, []Tile{x})

				// draw another
				x = tiles[0]
				tiles = tiles[1:]
			}

			p.Hand = append(p.Hand, x)
		}
	}

	for _, t := range p.Displayed {
		rmId := -1
		for i := range p.Hand {
			if p.Hand[i].ID == t[0].ID { //TODO hack here since flower/animals would be in lenght-1 slices
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
		k := int(t.Suit)*10 + t.Value
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
			if s[0].Suit == s[1].Suit && s[0].Suit == s[2].Suit && s[0].Value == s[1].Value && s[0].Value == s[2].Value {
				k := int(s[0].Suit)*10 + s[0].Value
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
