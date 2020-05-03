export interface GameStateResponse {
    game_state: GameState;
    id: string;
    created_at: string;
    updated_at: string;
}

export interface GameState {
    starter: number;
    player_turn: number;
    turn_number: number;
    is_transitioning: boolean;
    player_map: object;
    remaining_tiles: ITile[];
    discarded_tiles: ITile[];
    last_discarded_tile: ITile;
}

export interface ITile {
    suit: number;
    value: number;
    id: string;
}

export interface IMove {
    tile: ITile;
    action: number;
    player: number;
    turn_number: number;
}