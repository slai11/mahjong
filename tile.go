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
func initSet() []Tile {
	i := 0
	mj := make([]Tile, 148)

	// init 3 suits
	for _, s := range []Suit{Bamboo, Coin, Number} {
		for v := 0; v < 9; v++ {
			for j := 0; j < 4; j++ {
				mj[i] = Tile{Suit: s, Value: v, ID: i}
				i += 1
			}
		}
	}
	for _, v := range []int{East, South, West, North} {
		for j := 0; j < 4; j++ {
			mj[i] = Tile{Suit: Wind, Value: v, ID: i}
			i += 1
		}
	}

	for _, v := range []int{Green, Red, White} {
		for j := 0; j < 4; j++ {
			mj[i] = Tile{Suit: Dragon, Value: v, ID: i}
			i += 1
		}
	}

	for v := 0; v < 8; v++ {
		mj[i] = Tile{Suit: Flower, Value: v, ID: i}
		i += 1
	}

	for _, v := range []int{Centipede, Chicken, Cat, Mouse} {
		mj[i] = Tile{Suit: Animal, Value: v, ID: i}
		i += 1
	}

	if i != 148 {
		panic("need 148 tiles")
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(mj), func(i, j int) { mj[i], mj[j] = mj[j], mj[i] })
	return mj
}
