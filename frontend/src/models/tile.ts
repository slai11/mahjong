import {ITile} from "./game_state";

export interface UniqueTile {
	name: string;
	id: number;
}

export const TileNameMap: { [key: number]: UniqueTile } = {
	0: {"name": "bamboo-1", "id": 0},
	1: {"name": "bamboo-2", "id": 4},
	2: {"name": "bamboo-3", "id": 8},
	3: {"name": "bamboo-4", "id": 12},
	4: {"name": "bamboo-5", "id": 16},
	5: {"name": "bamboo-6", "id": 20},
	6: {"name": "bamboo-7", "id": 24},
	7: {"name": "bamboo-8", "id": 28},
	8: {"name": "bamboo-9", "id": 32},
	10: {"name": "coin-1", "id": 36},
	11: {"name": "coin-2", "id": 40},
	12: {"name": "coin-3", "id": 44},
	13: {"name": "coin-4", "id": 48},
	14: {"name": "coin-5", "id": 52},
	15: {"name": "coin-6", "id": 56},
	16: {"name": "coin-7", "id": 60},
	17: {"name": "coin-8", "id": 64},
	18: {"name": "coin-9", "id": 68},
	20: {"name": "wan-1", "id": 72},
	21: {"name": "wan-2", "id": 76},
	22: {"name": "wan-3", "id": 80},
	23: {"name": "wan-4", "id": 84},
	24: {"name": "wan-5", "id": 88},
	25: {"name": "wan-6", "id": 92},
	26: {"name": "wan-7", "id": 96},
	27: {"name": "wan-8", "id": 100},
	28: {"name": "wan-9", "id": 104},
	30: {"name": "east", "id": 108},
	31: {"name": "south", "id": 112},
	32: {"name": "west", "id": 116},
	33: {"name": "north", "id": 120},
	40: {"name": "green-dragon", "id": 124},
	41: {"name": "red-dragon", "id": 128},
	42: {"name": "white-dragon", "id": 132},
}

export const TileList: ITile[] = [
	// bamboo
    {suit: 0, value: 0, id: "bamboo-1"}, {suit: 0, value: 0, id: "bamboo-1"}, {suit: 0, value: 0, id: "bamboo-1"}, {suit: 0, value: 0, id: "bamboo-1"},
	{suit: 0, value: 1, id: "bamboo-2"}, {suit: 0, value: 1, id: "bamboo-2"}, {suit: 0, value: 1, id: "bamboo-2"}, {suit: 0, value: 1, id: "bamboo-2"},
	{suit: 0, value: 2, id: "bamboo-3"}, {suit: 0, value: 2, id: "bamboo-3"}, {suit: 0, value: 2, id: "bamboo-3"}, {suit: 0, value: 2, id: "bamboo-3"},
	{suit: 0, value: 3, id: "bamboo-4"}, {suit: 0, value: 3, id: "bamboo-4"}, {suit: 0, value: 3, id: "bamboo-4"}, {suit: 0, value: 3, id: "bamboo-4"},
	{suit: 0, value: 4, id: "bamboo-5"}, {suit: 0, value: 4, id: "bamboo-5"}, {suit: 0, value: 4, id: "bamboo-5"}, {suit: 0, value: 4, id: "bamboo-5"},
	{suit: 0, value: 5, id: "bamboo-6"}, {suit: 0, value: 5, id: "bamboo-6"}, {suit: 0, value: 5, id: "bamboo-6"}, {suit: 0, value: 5, id: "bamboo-6"},
	{suit: 0, value: 6, id: "bamboo-7"}, {suit: 0, value: 6, id: "bamboo-7"}, {suit: 0, value: 6, id: "bamboo-7"}, {suit: 0, value: 6, id: "bamboo-7"},
	{suit: 0, value: 7, id: "bamboo-8"}, {suit: 0, value: 7, id: "bamboo-8"}, {suit: 0, value: 7, id: "bamboo-8"}, {suit: 0, value: 7, id: "bamboo-8"},
	{suit: 0, value: 8, id: "bamboo-9"}, {suit: 0, value: 8, id: "bamboo-9"}, {suit: 0, value: 8, id: "bamboo-9"}, {suit: 0, value: 8, id: "bamboo-9"},
	// coin suit
	{suit: 1, value: 0, id: "coin-1"}, {suit: 1, value: 0, id: "coin-1"}, {suit: 1, value: 0, id: "coin-1"}, {suit: 1, value: 0, id: "coin-1"},
	{suit: 1, value: 1, id: "coin-2"}, {suit: 1, value: 1, id: "coin-2"}, {suit: 1, value: 1, id: "coin-2"}, {suit: 1, value: 1, id: "coin-2"},
	{suit: 1, value: 2, id: "coin-3"}, {suit: 1, value: 2, id: "coin-3"}, {suit: 1, value: 2, id: "coin-3"}, {suit: 1, value: 2, id: "coin-3"},
	{suit: 1, value: 3, id: "coin-4"}, {suit: 1, value: 3, id: "coin-4"}, {suit: 1, value: 3, id: "coin-4"}, {suit: 1, value: 3, id: "coin-4"},
	{suit: 1, value: 4, id: "coin-5"}, {suit: 1, value: 4, id: "coin-5"}, {suit: 1, value: 4, id: "coin-5"}, {suit: 1, value: 4, id: "coin-5"},
	{suit: 1, value: 5, id: "coin-6"}, {suit: 1, value: 5, id: "coin-6"}, {suit: 1, value: 5, id: "coin-6"}, {suit: 1, value: 5, id: "coin-6"},
	{suit: 1, value: 6, id: "coin-7"}, {suit: 1, value: 6, id: "coin-7"}, {suit: 1, value: 6, id: "coin-7"}, {suit: 1, value: 6, id: "coin-7"},
	{suit: 1, value: 7, id: "coin-8"}, {suit: 1, value: 7, id: "coin-8"}, {suit: 1, value: 7, id: "coin-8"}, {suit: 1, value: 7, id: "coin-8"},
	{suit: 1, value: 8, id: "coin-9"}, {suit: 1, value: 8, id: "coin-9"}, {suit: 1, value: 8, id: "coin-9"}, {suit: 1, value: 8, id: "coin-9"},
	// number suit
	{suit: 2, value: 0, id: "wan-1"}, {suit: 2, value: 0, id: "wan-1"}, {suit: 2, value: 0, id: "wan-1"}, {suit: 2, value: 0, id: "wan-1"},
	{suit: 2, value: 1, id: "wan-2"}, {suit: 2, value: 1, id: "wan-2"}, {suit: 2, value: 1, id: "wan-2"}, {suit: 2, value: 1, id: "wan-2"},
	{suit: 2, value: 2, id: "wan-3"}, {suit: 2, value: 2, id: "wan-3"}, {suit: 2, value: 2, id: "wan-3"}, {suit: 2, value: 2, id: "wan-3"},
	{suit: 2, value: 3, id: "wan-4"}, {suit: 2, value: 3, id: "wan-4"}, {suit: 2, value: 3, id: "wan-4"}, {suit: 2, value: 3, id: "wan-4"},
	{suit: 2, value: 4, id: "wan-5"}, {suit: 2, value: 4, id: "wan-5"}, {suit: 2, value: 4, id: "wan-5"}, {suit: 2, value: 4, id: "wan-5"},
	{suit: 2, value: 5, id: "wan-6"}, {suit: 2, value: 5, id: "wan-6"}, {suit: 2, value: 5, id: "wan-6"}, {suit: 2, value: 5, id: "wan-6"},
	{suit: 2, value: 6, id: "wan-7"}, {suit: 2, value: 6, id: "wan-7"}, {suit: 2, value: 6, id: "wan-7"}, {suit: 2, value: 6, id: "wan-7"},
	{suit: 2, value: 7, id: "wan-8"}, {suit: 2, value: 7, id: "wan-8"}, {suit: 2, value: 7, id: "wan-8"}, {suit: 2, value: 7, id: "wan-8"},
	{suit: 2, value: 8, id: "wan-9"}, {suit: 2, value: 8, id: "wan-9"}, {suit: 2, value: 8, id: "wan-9"}, {suit: 2, value: 8, id: "wan-9"},
	// wind suit
	{suit: 3, value: 0, id: "east"}, {suit: 3, value: 0, id: "east"}, {suit: 3, value: 0, id: "east"}, {suit: 3, value: 0, id: "east"},
	{suit: 3, value: 1, id: "south"}, {suit: 3, value: 1, id: "south"}, {suit: 3, value: 1, id: "south"}, {suit: 3, value: 1, id: "south"},
	{suit: 3, value: 2, id: "west"}, {suit: 3, value: 2, id: "west"}, {suit: 3, value: 2, id: "west"}, {suit: 3, value: 2, id: "west"},
	{suit: 3, value: 3, id: "north"}, {suit: 3, value: 3, id: "north"}, {suit: 3, value: 3, id: "north"}, {suit: 3, value: 3, id: "north"},
	// dragon suit
	{suit: 4, value: 0, id: "green-dragon"}, {suit: 4, value: 0, id: "green-dragon"}, {suit: 4, value: 0, id: "green-dragon"}, {suit: 4, value: 0, id: "green-dragon"},
	{suit: 4, value: 1, id: "red-dragon"}, {suit: 4, value: 1, id: "red-dragon"}, {suit: 4, value: 1, id: "red-dragon"}, {suit: 4, value: 1, id: "red-dragon"},
	{suit: 4, value: 2, id: "white-dragon"}, {suit: 4, value: 2, id: "white-dragon"}, {suit: 4, value: 2, id: "white-dragon"}, {suit: 4, value: 2, id: "white-dragon"},
	// animal suit
	{suit: 5, value: 0, id: "centipede"}, {suit: 5, value: 1, id: "chicken"}, {suit: 5, value: 2, id: "cat"}, {suit: 5, value: 3, id: "mouse"},
	// flower suit
	{suit: 6, value: 0, id: "flower-1"}, {suit: 6, value: 1, id: "flower-2"}, {suit: 6, value: 2, id: "flower-3"}, {suit: 6, value: 3, id: "flower-4"},
	{suit: 6, value: 4, id: "season-1"}, {suit: 6, value: 5, id: "season-2"}, {suit: 6, value: 6, id: "season-3"}, {suit: 6, value: 7, id: "season-4"},
]