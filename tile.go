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

	// unique to every tile
	ID int `json:"id"`
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
	return mj
}

// TileList is a centralised point of reference for tile information
// its a length-148 array of Tiles.
var TileList = [148]Tile{
	// bamboo suit
	Tile{Suit: 0, Value: 0, ID: 0}, Tile{Suit: 0, Value: 0, ID: 1}, Tile{Suit: 0, Value: 0, ID: 2}, Tile{Suit: 0, Value: 0, ID: 3},
	Tile{Suit: 0, Value: 1, ID: 4}, Tile{Suit: 0, Value: 1, ID: 5}, Tile{Suit: 0, Value: 1, ID: 6}, Tile{Suit: 0, Value: 1, ID: 7},
	Tile{Suit: 0, Value: 2, ID: 8}, Tile{Suit: 0, Value: 2, ID: 9}, Tile{Suit: 0, Value: 2, ID: 10}, Tile{Suit: 0, Value: 2, ID: 11},
	Tile{Suit: 0, Value: 3, ID: 12}, Tile{Suit: 0, Value: 3, ID: 13}, Tile{Suit: 0, Value: 3, ID: 14}, Tile{Suit: 0, Value: 3, ID: 15},
	Tile{Suit: 0, Value: 4, ID: 16}, Tile{Suit: 0, Value: 4, ID: 17}, Tile{Suit: 0, Value: 4, ID: 18}, Tile{Suit: 0, Value: 4, ID: 19},
	Tile{Suit: 0, Value: 5, ID: 20}, Tile{Suit: 0, Value: 5, ID: 21}, Tile{Suit: 0, Value: 5, ID: 22}, Tile{Suit: 0, Value: 5, ID: 23},
	Tile{Suit: 0, Value: 6, ID: 24}, Tile{Suit: 0, Value: 6, ID: 25}, Tile{Suit: 0, Value: 6, ID: 26}, Tile{Suit: 0, Value: 6, ID: 27},
	Tile{Suit: 0, Value: 7, ID: 28}, Tile{Suit: 0, Value: 7, ID: 29}, Tile{Suit: 0, Value: 7, ID: 30}, Tile{Suit: 0, Value: 7, ID: 31},
	Tile{Suit: 0, Value: 8, ID: 32}, Tile{Suit: 0, Value: 8, ID: 33}, Tile{Suit: 0, Value: 8, ID: 34}, Tile{Suit: 0, Value: 8, ID: 35},
	// coin suit
	Tile{Suit: 1, Value: 0, ID: 36}, Tile{Suit: 1, Value: 0, ID: 37}, Tile{Suit: 1, Value: 0, ID: 38}, Tile{Suit: 1, Value: 0, ID: 39},
	Tile{Suit: 1, Value: 1, ID: 40}, Tile{Suit: 1, Value: 1, ID: 41}, Tile{Suit: 1, Value: 1, ID: 42}, Tile{Suit: 1, Value: 1, ID: 43},
	Tile{Suit: 1, Value: 2, ID: 44}, Tile{Suit: 1, Value: 2, ID: 45}, Tile{Suit: 1, Value: 2, ID: 46}, Tile{Suit: 1, Value: 2, ID: 47},
	Tile{Suit: 1, Value: 3, ID: 48}, Tile{Suit: 1, Value: 3, ID: 49}, Tile{Suit: 1, Value: 3, ID: 50}, Tile{Suit: 1, Value: 3, ID: 51},
	Tile{Suit: 1, Value: 4, ID: 52}, Tile{Suit: 1, Value: 4, ID: 53}, Tile{Suit: 1, Value: 4, ID: 54}, Tile{Suit: 1, Value: 4, ID: 55},
	Tile{Suit: 1, Value: 5, ID: 56}, Tile{Suit: 1, Value: 5, ID: 57}, Tile{Suit: 1, Value: 5, ID: 58}, Tile{Suit: 1, Value: 5, ID: 59},
	Tile{Suit: 1, Value: 6, ID: 60}, Tile{Suit: 1, Value: 6, ID: 61}, Tile{Suit: 1, Value: 6, ID: 62}, Tile{Suit: 1, Value: 6, ID: 63},
	Tile{Suit: 1, Value: 7, ID: 64}, Tile{Suit: 1, Value: 7, ID: 65}, Tile{Suit: 1, Value: 7, ID: 66}, Tile{Suit: 1, Value: 7, ID: 67},
	Tile{Suit: 1, Value: 8, ID: 68}, Tile{Suit: 1, Value: 8, ID: 69}, Tile{Suit: 1, Value: 8, ID: 70}, Tile{Suit: 1, Value: 8, ID: 71},
	// number suit
	Tile{Suit: 2, Value: 0, ID: 72}, Tile{Suit: 2, Value: 0, ID: 73}, Tile{Suit: 2, Value: 0, ID: 74}, Tile{Suit: 2, Value: 0, ID: 75},
	Tile{Suit: 2, Value: 1, ID: 76}, Tile{Suit: 2, Value: 1, ID: 77}, Tile{Suit: 2, Value: 1, ID: 78}, Tile{Suit: 2, Value: 1, ID: 79},
	Tile{Suit: 2, Value: 2, ID: 80}, Tile{Suit: 2, Value: 2, ID: 81}, Tile{Suit: 2, Value: 2, ID: 82}, Tile{Suit: 2, Value: 2, ID: 83},
	Tile{Suit: 2, Value: 3, ID: 84}, Tile{Suit: 2, Value: 3, ID: 85}, Tile{Suit: 2, Value: 3, ID: 86}, Tile{Suit: 2, Value: 3, ID: 87},
	Tile{Suit: 2, Value: 4, ID: 88}, Tile{Suit: 2, Value: 4, ID: 89}, Tile{Suit: 2, Value: 4, ID: 90}, Tile{Suit: 2, Value: 4, ID: 91},
	Tile{Suit: 2, Value: 5, ID: 92}, Tile{Suit: 2, Value: 5, ID: 93}, Tile{Suit: 2, Value: 5, ID: 94}, Tile{Suit: 2, Value: 5, ID: 95},
	Tile{Suit: 2, Value: 6, ID: 96}, Tile{Suit: 2, Value: 6, ID: 97}, Tile{Suit: 2, Value: 6, ID: 98}, Tile{Suit: 2, Value: 6, ID: 99},
	Tile{Suit: 2, Value: 7, ID: 100}, Tile{Suit: 2, Value: 7, ID: 101}, Tile{Suit: 2, Value: 7, ID: 102}, Tile{Suit: 2, Value: 7, ID: 103},
	Tile{Suit: 2, Value: 8, ID: 104}, Tile{Suit: 2, Value: 8, ID: 105}, Tile{Suit: 2, Value: 8, ID: 106}, Tile{Suit: 2, Value: 8, ID: 107},
	// wind suit
	Tile{Suit: 3, Value: 0, ID: 108}, Tile{Suit: 3, Value: 0, ID: 109}, Tile{Suit: 3, Value: 0, ID: 110}, Tile{Suit: 3, Value: 0, ID: 111},
	Tile{Suit: 3, Value: 1, ID: 112}, Tile{Suit: 3, Value: 1, ID: 113}, Tile{Suit: 3, Value: 1, ID: 114}, Tile{Suit: 3, Value: 1, ID: 115},
	Tile{Suit: 3, Value: 2, ID: 116}, Tile{Suit: 3, Value: 2, ID: 117}, Tile{Suit: 3, Value: 2, ID: 118}, Tile{Suit: 3, Value: 2, ID: 119},
	Tile{Suit: 3, Value: 3, ID: 120}, Tile{Suit: 3, Value: 3, ID: 121}, Tile{Suit: 3, Value: 3, ID: 122}, Tile{Suit: 3, Value: 3, ID: 123},
	// dragon suit
	Tile{Suit: 4, Value: 0, ID: 124}, Tile{Suit: 4, Value: 0, ID: 125}, Tile{Suit: 4, Value: 0, ID: 126}, Tile{Suit: 4, Value: 0, ID: 127},
	Tile{Suit: 4, Value: 1, ID: 128}, Tile{Suit: 4, Value: 1, ID: 129}, Tile{Suit: 4, Value: 1, ID: 130}, Tile{Suit: 4, Value: 1, ID: 131},
	Tile{Suit: 4, Value: 2, ID: 132}, Tile{Suit: 4, Value: 2, ID: 133}, Tile{Suit: 4, Value: 2, ID: 134}, Tile{Suit: 4, Value: 2, ID: 135},
	// animal suit
	Tile{Suit: 5, Value: 0, ID: 136}, Tile{Suit: 5, Value: 1, ID: 137}, Tile{Suit: 5, Value: 2, ID: 138}, Tile{Suit: 5, Value: 3, ID: 139},
	// flower suit
	Tile{Suit: 6, Value: 0, ID: 140}, Tile{Suit: 6, Value: 1, ID: 141}, Tile{Suit: 6, Value: 2, ID: 142}, Tile{Suit: 6, Value: 3, ID: 143},
	Tile{Suit: 6, Value: 4, ID: 144}, Tile{Suit: 6, Value: 5, ID: 145}, Tile{Suit: 6, Value: 6, ID: 146}, Tile{Suit: 6, Value: 7, ID: 147},
}
