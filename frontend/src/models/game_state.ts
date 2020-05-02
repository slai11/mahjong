

export interface GameState {
    starter: number;
    player_turn: number;
    turn_number: number;
    is_transitioning: boolean;
    player_map: object;
    remaining_tiles: Tile[];
    discarded_tiles: Tile[];
    last_discarded_tile: Tile;
}

export interface Tile {
    suit: number;
    value: number;
    id: number;
}
