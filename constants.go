package main

// possible actions to transition between states
type Action int

// 0 discard, 1 draw
// 2 eat, 3 eatright, 4 eatleft
// 5 pong
// 6 gong, 7 innergong
// 8 call
const (
	Discard Action = iota
	Draw
	Eat
	EatRight
	EatLeft
	Pong
	Gong
	InnerGong
	Call
)

// 4 possible players with a util function for next to loop them
type Player int

const (
	P0 Player = iota
	P1
	P2
	P3
)

func (p Player) next() Player {
	switch p {
	case P0, P1, P2:
		return p + 1
	case P3:
		return P0
	}
	panic("what??")
}

type PrevailingWind int

// TODO refactor to avoid using hanyupinyin for winds
const (
	Dong PrevailingWind = iota
	Nan
	Xi
	Bei
)

func (w PrevailingWind) next() PrevailingWind {
	switch w {
	case Dong, Nan, Xi:
		return w + 1
	case Bei:
		return Dong
	}
	panic("no such wind")
}

type Suit int

const (
	Bamboo Suit = iota
	Coin
	Number
	Wind
	Dragon
	Animal
	Flower

	East int = iota
	South
	West
	North

	Green int = iota
	Red
	White

	Centipede int = iota
	Chicken
	Cat
	Mouse
)
