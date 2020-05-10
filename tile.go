package main

import (
	"math/rand"
	"time"
)

// 148 tiles in total
type Tile struct {

	// bamboo, coin, number, wind, dragon, flower, animal
	Suit `json:"suit"`

	// 0-8
	Value int `json:"value"`
}

// run once at start of server
func initSet() []int {
	n := 148
	mj := make([]int, n)
	for i := 0; i < n; i++ {
		mj[i] = i
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(mj), func(i, j int) { mj[i], mj[j] = mj[j], mj[i] })
	rand.Shuffle(len(mj), func(i, j int) { mj[i], mj[j] = mj[j], mj[i] })
	rand.Shuffle(len(mj), func(i, j int) { mj[i], mj[j] = mj[j], mj[i] })
	rand.Shuffle(len(mj), func(i, j int) { mj[i], mj[j] = mj[j], mj[i] })

	if len(mj) != 148 {
		panic("invalid starting set")
	}
	return mj
}

// TileList is a centralised point of reference for tile information
// its a length-148 array of Tiles.
var TileList = [148]Tile{
	// bamboo suit
	Tile{Suit: 0, Value: 0}, Tile{Suit: 0, Value: 0}, Tile{Suit: 0, Value: 0}, Tile{Suit: 0, Value: 0}, // 0
	Tile{Suit: 0, Value: 1}, Tile{Suit: 0, Value: 1}, Tile{Suit: 0, Value: 1}, Tile{Suit: 0, Value: 1}, // 4
	Tile{Suit: 0, Value: 2}, Tile{Suit: 0, Value: 2}, Tile{Suit: 0, Value: 2}, Tile{Suit: 0, Value: 2}, // 8
	Tile{Suit: 0, Value: 3}, Tile{Suit: 0, Value: 3}, Tile{Suit: 0, Value: 3}, Tile{Suit: 0, Value: 3}, //12
	Tile{Suit: 0, Value: 4}, Tile{Suit: 0, Value: 4}, Tile{Suit: 0, Value: 4}, Tile{Suit: 0, Value: 4}, // 16
	Tile{Suit: 0, Value: 5}, Tile{Suit: 0, Value: 5}, Tile{Suit: 0, Value: 5}, Tile{Suit: 0, Value: 5}, // 20
	Tile{Suit: 0, Value: 6}, Tile{Suit: 0, Value: 6}, Tile{Suit: 0, Value: 6}, Tile{Suit: 0, Value: 6}, // 24
	Tile{Suit: 0, Value: 7}, Tile{Suit: 0, Value: 7}, Tile{Suit: 0, Value: 7}, Tile{Suit: 0, Value: 7}, // 28
	Tile{Suit: 0, Value: 8}, Tile{Suit: 0, Value: 8}, Tile{Suit: 0, Value: 8}, Tile{Suit: 0, Value: 8}, // 32
	// coin suit
	Tile{Suit: 1, Value: 0}, Tile{Suit: 1, Value: 0}, Tile{Suit: 1, Value: 0}, Tile{Suit: 1, Value: 0}, // 36
	Tile{Suit: 1, Value: 1}, Tile{Suit: 1, Value: 1}, Tile{Suit: 1, Value: 1}, Tile{Suit: 1, Value: 1}, // 40
	Tile{Suit: 1, Value: 2}, Tile{Suit: 1, Value: 2}, Tile{Suit: 1, Value: 2}, Tile{Suit: 1, Value: 2}, // 44
	Tile{Suit: 1, Value: 3}, Tile{Suit: 1, Value: 3}, Tile{Suit: 1, Value: 3}, Tile{Suit: 1, Value: 3}, // 48
	Tile{Suit: 1, Value: 4}, Tile{Suit: 1, Value: 4}, Tile{Suit: 1, Value: 4}, Tile{Suit: 1, Value: 4}, // 52
	Tile{Suit: 1, Value: 5}, Tile{Suit: 1, Value: 5}, Tile{Suit: 1, Value: 5}, Tile{Suit: 1, Value: 5}, // 56
	Tile{Suit: 1, Value: 6}, Tile{Suit: 1, Value: 6}, Tile{Suit: 1, Value: 6}, Tile{Suit: 1, Value: 6}, // 60
	Tile{Suit: 1, Value: 7}, Tile{Suit: 1, Value: 7}, Tile{Suit: 1, Value: 7}, Tile{Suit: 1, Value: 7}, // 64
	Tile{Suit: 1, Value: 8}, Tile{Suit: 1, Value: 8}, Tile{Suit: 1, Value: 8}, Tile{Suit: 1, Value: 8}, // 68
	// number suit
	Tile{Suit: 2, Value: 0}, Tile{Suit: 2, Value: 0}, Tile{Suit: 2, Value: 0}, Tile{Suit: 2, Value: 0}, // 72
	Tile{Suit: 2, Value: 1}, Tile{Suit: 2, Value: 1}, Tile{Suit: 2, Value: 1}, Tile{Suit: 2, Value: 1}, // 76
	Tile{Suit: 2, Value: 2}, Tile{Suit: 2, Value: 2}, Tile{Suit: 2, Value: 2}, Tile{Suit: 2, Value: 2}, // 80
	Tile{Suit: 2, Value: 3}, Tile{Suit: 2, Value: 3}, Tile{Suit: 2, Value: 3}, Tile{Suit: 2, Value: 3}, // 84
	Tile{Suit: 2, Value: 4}, Tile{Suit: 2, Value: 4}, Tile{Suit: 2, Value: 4}, Tile{Suit: 2, Value: 4}, // 88
	Tile{Suit: 2, Value: 5}, Tile{Suit: 2, Value: 5}, Tile{Suit: 2, Value: 5}, Tile{Suit: 2, Value: 5}, // 92
	Tile{Suit: 2, Value: 6}, Tile{Suit: 2, Value: 6}, Tile{Suit: 2, Value: 6}, Tile{Suit: 2, Value: 6}, // 96
	Tile{Suit: 2, Value: 7}, Tile{Suit: 2, Value: 7}, Tile{Suit: 2, Value: 7}, Tile{Suit: 2, Value: 7}, // 100
	Tile{Suit: 2, Value: 8}, Tile{Suit: 2, Value: 8}, Tile{Suit: 2, Value: 8}, Tile{Suit: 2, Value: 8}, // 104
	// wind suit
	Tile{Suit: 3, Value: 0}, Tile{Suit: 3, Value: 0}, Tile{Suit: 3, Value: 0}, Tile{Suit: 3, Value: 0}, // 108
	Tile{Suit: 3, Value: 1}, Tile{Suit: 3, Value: 1}, Tile{Suit: 3, Value: 1}, Tile{Suit: 3, Value: 1}, // 112
	Tile{Suit: 3, Value: 2}, Tile{Suit: 3, Value: 2}, Tile{Suit: 3, Value: 2}, Tile{Suit: 3, Value: 2}, // 116
	Tile{Suit: 3, Value: 3}, Tile{Suit: 3, Value: 3}, Tile{Suit: 3, Value: 3}, Tile{Suit: 3, Value: 3}, // 120
	// dragon suit
	Tile{Suit: 4, Value: 0}, Tile{Suit: 4, Value: 0}, Tile{Suit: 4, Value: 0}, Tile{Suit: 4, Value: 0}, // 124
	Tile{Suit: 4, Value: 1}, Tile{Suit: 4, Value: 1}, Tile{Suit: 4, Value: 1}, Tile{Suit: 4, Value: 1}, // 128
	Tile{Suit: 4, Value: 2}, Tile{Suit: 4, Value: 2}, Tile{Suit: 4, Value: 2}, Tile{Suit: 4, Value: 2}, // 132
	// animal suit
	Tile{Suit: 5, Value: 0}, Tile{Suit: 5, Value: 1}, Tile{Suit: 5, Value: 2}, Tile{Suit: 5, Value: 3}, // 136
	// flower suit
	Tile{Suit: 6, Value: 0}, Tile{Suit: 6, Value: 1}, Tile{Suit: 6, Value: 2}, Tile{Suit: 6, Value: 3}, // 140
	Tile{Suit: 6, Value: 4}, Tile{Suit: 6, Value: 5}, Tile{Suit: 6, Value: 6}, Tile{Suit: 6, Value: 7}, // 144
}
