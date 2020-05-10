package main

import (
	"testing"
)

func TestDraw(t *testing.T) {
	type expected struct {
		lastDrawn     int
		handSize      int
		remainingSize int
	}
	testcases := []struct {
		name      string
		hand      []int
		remaining []int
		expected  expected
	}{
		{
			name:      "standard draw",
			hand:      []int{1, 2, 3, 4},
			remaining: []int{5, 6, 7, 8},
			expected: expected{
				handSize:      5,
				remainingSize: 3,
				lastDrawn:     5,
			},
		},
		{
			name:      "empty remaining",
			hand:      []int{1, 2, 3, 4},
			remaining: []int{},
			expected: expected{
				handSize:      4,
				remainingSize: 0,
				lastDrawn:     0,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ps := PlayerState{
				Hand: tc.hand,
			}

			remaining, _ := ps.Draw(tc.remaining)
			if len(ps.Hand) != tc.expected.handSize {
				t.Errorf("hand: expected %v, got %v", tc.expected.handSize, len(ps.Hand))
				return
			}

			if len(remaining) != tc.expected.remainingSize {
				t.Errorf("remaining: expected %v, got %v", tc.expected.remainingSize, len(remaining))
				return
			}

			if ps.LastDrawnTile != tc.expected.lastDrawn {
				t.Errorf("lastDrawn: expected %v, got %v", tc.expected.lastDrawn, ps.LastDrawnTile)
				return
			}
			return
		})
	}
}

func TestDiscard(t *testing.T) {
	type expected struct {
		handSize int
		hasErr   bool
	}
	testcases := []struct {
		name     string
		hand     []int
		discard  int
		expected expected
	}{
		{
			name:    "discard found",
			hand:    []int{1, 2, 3, 4},
			discard: 1,
			expected: expected{
				handSize: 3,
				hasErr:   false,
			},
		},
		{
			name:    "discard only 1 tile",
			hand:    []int{1, 2, 3, 4, 1},
			discard: 1,
			expected: expected{
				handSize: 4,
				hasErr:   false,
			},
		},
		{
			name:    "discard not found",
			hand:    []int{1, 2, 3, 4},
			discard: 5,
			expected: expected{
				handSize: 4,
				hasErr:   true,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ps := PlayerState{
				Hand: tc.hand,
			}

			err := ps.Discard(tc.discard)
			if err != nil && !tc.expected.hasErr {
				t.Errorf("expected %v, got %v", tc.expected.hasErr, err)
				return
			}
			if len(ps.Hand) != tc.expected.handSize {
				t.Errorf("hand: expected %v, got %v", tc.expected.handSize, len(ps.Hand))
				return
			}
			return
		})
	}
}

func TestEat(t *testing.T) {
	type expected struct {
		displayed   []int
		length      int
		returnError bool
	}
	testcases := []struct {
		name        string
		hand        []int
		canEat      bool
		canEatLeft  bool
		canEatRight bool
		action      Action
		tile        int
		expected
	}{
		{
			name:   "exits if boolean flag not set",
			hand:   []int{1, 2, 3, 4},
			action: Eat,
			expected: expected{
				displayed:   []int{},
				length:      4,
				returnError: true,
			},
		},
		{
			// eats 2 bamboo
			name:   "eats correctly",
			hand:   []int{0, 8},
			tile:   5,
			action: Eat,
			canEat: true,
			expected: expected{
				displayed: []int{5, 0, 8},
				length:    0,
			},
		},
		{
			// eats 2 bamboo on left of 3 + 4 bamboo
			name:       "eats left correctly",
			hand:       []int{0, 8, 12},
			tile:       5,
			action:     EatLeft,
			canEatLeft: true,
			expected: expected{
				displayed: []int{5, 8, 12},
				length:    1,
			},
		},
		{
			// eats 5 bamboo on left of 3 + 4 bamboo
			name:        "eats right correctly",
			hand:        []int{0, 8, 12},
			tile:        17,
			action:      EatRight,
			canEatRight: true,
			expected: expected{
				displayed: []int{17, 8, 12},
				length:    1,
			},
		},
		{
			// failing to eat should leave hand untouched
			name:        "eats right correctly",
			hand:        []int{0, 8, 20}, // 1, 3, 5 bamboo
			tile:        17,              // 5 bamboo
			action:      EatRight,
			canEatRight: true,
			expected: expected{
				length:      3,
				returnError: true,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ps := PlayerState{
				Hand:        tc.hand,
				CanEat:      tc.canEat,
				CanEatLeft:  tc.canEatLeft,
				CanEatRight: tc.canEatRight,
				Displayed:   [][]int{},
			}

			err := ps.Eat(tc.tile, tc.action)
			if tc.returnError {
				if len(ps.Hand) != tc.expected.length || err == nil {
					t.Errorf("handsize should not have changed")
					return
				}
				return
			}

			if tc.expected.length != len(ps.Hand) {
				t.Errorf("expected length %v but got %v", tc.expected.length, len(ps.Hand))
				return
			}

			found := 0
			for _, tile := range ps.Displayed[0] {
				for _, dtile := range tc.expected.displayed {
					if tile == dtile {
						found++
					}
				}
			}

			if found != 3 {
				t.Errorf("expected length 3 but got %v", found)
				return
			}

		})
	}
}

func TestPong(t *testing.T) {
	type expected struct {
		returnError bool
	}

	testcases := []struct {
		name    string
		tile    int
		hand    []int
		canPong bool
		expected
	}{

		{
			name:    "cannot pong",
			tile:    0,
			canPong: false,
			expected: expected{
				returnError: true,
			},
		},
		{
			name:    "pong",
			tile:    0,
			hand:    []int{1, 2},
			canPong: true,
			expected: expected{
				returnError: false,
			},
		},
		{
			name:    "can pong but triplet missing",
			tile:    0,
			hand:    []int{1, 10},
			canPong: true,
			expected: expected{
				returnError: true,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			ps := PlayerState{
				Hand:      tc.hand,
				Displayed: [][]int{},
				CanPong:   tc.canPong,
			}

			err := ps.Pong(tc.tile)
			if tc.expected.returnError {
				if err == nil {
					t.Errorf("expected no error but got %v", err)
					return
				}
				return
			}

			for _, dtile := range ps.Displayed[0] {
				if TileList[dtile].Suit != TileList[tc.tile].Suit || TileList[dtile].Value != TileList[tc.tile].Value {
					t.Errorf("ponged %v but got %v", tc.tile, dtile)
					return
				}
			}
		})
	}
}

func TestGong(t *testing.T) {
	type expected struct {
		returnError bool
	}

	testcases := []struct {
		name    string
		tile    int
		hand    []int
		canGong bool
		expected
	}{

		{
			name:    "cannot gong",
			tile:    0,
			canGong: false,
			expected: expected{
				returnError: true,
			},
		},
		{
			name:    "gong",
			tile:    0,
			hand:    []int{1, 2, 3},
			canGong: true,
			expected: expected{
				returnError: false,
			},
		},
		{
			name:    "can gong but triplet missing",
			tile:    0,
			hand:    []int{1, 10, 11},
			canGong: true,
			expected: expected{
				returnError: true,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			ps := PlayerState{
				Hand:      tc.hand,
				Displayed: [][]int{},
				CanGong:   tc.canGong,
			}

			err := ps.Gong(tc.tile)
			if tc.expected.returnError {
				if err == nil {
					t.Errorf("expected no error but got %v", err)
					return
				}
				return
			}

			for _, dtile := range ps.Displayed[0] {
				if TileList[dtile].Suit != TileList[tc.tile].Suit || TileList[dtile].Value != TileList[tc.tile].Value {
					t.Errorf("ponged %v but got %v", tc.tile, dtile)
					return
				}
			}
		})
	}
}

func TestRepair(t *testing.T) {
	type expected struct {
		tileCount      int
		displayedCount int
	}
	testcases := []struct {
		name  string
		tiles []int
		hand  []int
		expected
	}{
		{
			name:  "nothing to repair",
			hand:  []int{1, 2, 3, 4, 5},
			tiles: []int{6, 7, 8, 9, 10},
			expected: expected{
				tileCount:      5,
				displayedCount: 0,
			},
		},
		{
			name:  "1 animal",
			hand:  []int{1, 2, 3, 4, 136},
			tiles: []int{6, 141, 8, 9, 10},
			expected: expected{
				tileCount:      4,
				displayedCount: 1,
			},
		},
		{
			name:  "1 flower",
			hand:  []int{1, 2, 3, 4, 140},
			tiles: []int{6, 141, 8, 9, 10},
			expected: expected{
				tileCount:      4,
				displayedCount: 1,
			},
		},
		{
			name:  "1 flower in hand, 1 flower repaired",
			hand:  []int{1, 2, 3, 4, 140},
			tiles: []int{141, 7, 8, 9, 10},
			expected: expected{
				tileCount:      3,
				displayedCount: 2,
			},
		},
		{
			name:  "2 animal/flower in hand, 1 flower repaired",
			hand:  []int{1, 2, 3, 136, 140},
			tiles: []int{141, 7, 8, 9, 10},
			expected: expected{
				tileCount:      2,
				displayedCount: 3,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ps := PlayerState{
				Hand:      tc.hand,
				Displayed: [][]int{},
			}

			remaining := ps.RepairHand(tc.tiles)

			if len(remaining) != tc.expected.tileCount {
				t.Errorf("remaining: expected %v but got %v", tc.expected.tileCount, len(remaining))
				return

			}

			if len(tc.hand) != len(ps.Hand) {
				t.Errorf("hand: expected %v but got %v", len(tc.hand), len(ps.Hand))
				return

			}

			if len(ps.Displayed) != tc.expected.displayedCount {
				t.Errorf("displayed: expected %v but got %v", tc.expected.displayedCount, len(ps.Displayed))
				return
			}

		})
	}
}

func TestUpdateStatus(t *testing.T) {
	type expected struct {
		canPong     bool
		canEat      bool
		canEatLeft  bool
		canEatRight bool
		canGong     bool
	}
	testcases := []struct {
		name string
		hand []int
		tile int
		expected
	}{
		{
			// 1 bamboo cant be eaten
			name:     "nothing to update",
			hand:     []int{1, 10},
			tile:     1,
			expected: expected{},
		},
		{
			name: "can eat middle",
			hand: []int{1, 10},
			tile: 5,
			expected: expected{
				canEat: true,
			},
		},
		{
			name: "can eat left",
			hand: []int{10, 15},
			tile: 5,
			expected: expected{
				canEatLeft: true,
			},
		},
		{
			name: "can eat right",
			hand: []int{10, 15},
			tile: 19,
			expected: expected{
				canEatRight: true,
			},
		},
		{
			name: "can pong",
			hand: []int{10, 11},
			tile: 9,
			expected: expected{
				canPong: true,
			},
		},
		{
			name: "can gong",
			hand: []int{8, 10, 11},
			tile: 9,
			expected: expected{
				canGong: true,
				canPong: true,
			},
		},
		{
			name: "can eat left and middle",
			hand: []int{1, 8, 12},
			tile: 4,
			expected: expected{
				canEat:     true,
				canEatLeft: true,
			},
		},
		{
			name: "can eat right and middle",
			hand: []int{4, 8, 16},
			tile: 12,
			expected: expected{
				canEat:      true,
				canEatRight: true,
			},
		},
		{
			name: "can eat left, right and middle, pong and gong",
			hand: []int{4, 8, 16, 13, 14, 15, 20},
			tile: 12,
			expected: expected{
				canEat:      true,
				canEatRight: true,
				canEatLeft:  true,
				canPong:     true,
				canGong:     true,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ps := PlayerState{
				Hand: tc.hand,
			}
			ps.UpdateStatus(tc.tile)

			if tc.expected.canEat != ps.CanEat {
				t.Errorf("expected canEat to be %v but got %v", tc.expected.canEat, ps.CanEat)
				return
			}
			if tc.expected.canEatLeft != ps.CanEatLeft {
				t.Errorf("expected canEatLeft to be %v but got %v", tc.expected.canEatLeft, ps.CanEatLeft)
				return
			}
			if tc.expected.canEatRight != ps.CanEatRight {
				t.Errorf("expected canEatRight to be %v but got %v", tc.expected.canEatRight, ps.CanEatRight)
				return
			}
			if tc.expected.canPong != ps.CanPong {
				t.Errorf("expected canPong to be %v but got %v", tc.expected.canPong, ps.CanPong)
				return
			}
			if tc.expected.canGong != ps.CanGong {
				t.Errorf("expected canGong to be %v but got %v", tc.expected.canGong, ps.CanGong)
				return
			}
		})
	}
}
