<template>
  <div class="GameDashboard">
    <h4>Share this link with your friends: https://tableswim.com/{{gameID}}</h4>
    <h3>Table Number: {{gameID}} | You are player {{this.playerNumber + 1}} ({{ playerWind }}) | Prevailing wind: {{prevailingWind}} |  | Remaining Tile: {{remainingTileCount}}</h3>

    <div v-if="showWinningHand" class="text-center">
      <v-dialog v-if="this.info.game_state.last_winning_hand.stalemate === true" v-model="showWinningHand" width="800">
        <v-card>
          <h2>No one won the round!</h2>
        </v-card>
      </v-dialog>
      <v-dialog v-else v-model="showWinningHand" width="800">
        <v-card>
          <h2>Winner: Player {{this.info.game_state.last_winning_hand.player + 1}}</h2>
          <Hand :hand="this.info ? this.info.game_state.last_winning_hand.hand : null" :mark="winningTile"/>
        </v-card>
      </v-dialog>
    </div>

    <div v-if="info" class="container">
      <FriendInfo
        class="rightplayer"
        :info="info.game_state.player_map[friendOrder[0]]"
        :playerNumber="friendOrder[0]"
        :playerTurn="info.game_state.player_turn"
      />
      <FriendInfo
        class="oppositeplayer"
        :info="info.game_state.player_map[friendOrder[1]]"
        :playerNumber="friendOrder[1]"
        :playerTurn="info.game_state.player_turn"
      />
      <FriendInfo
        class="leftplayer"
        :info="info.game_state.player_map[friendOrder[2]]"
        :playerNumber="friendOrder[2]"
        :playerTurn="info.game_state.player_turn"
      />

      <div class="discard">
        <div v-for="(tile, id) in info.game_state.discarded_tiles" :key="id">
          <Tile :id="tile" />
        </div>
        <Tile
          v-if="info.game_state.last_discarded_tile !== null"
          :id="info.game_state.last_discarded_tile"
          :style="{outline: '2px double red'}"
        />
      </div>

      <Player
        class="player"
        :info="info.game_state.player_map[this.playerNumber]"
        :player_turn="info.game_state.player_turn"
        :player_number="playerNumber"
        :transiting="info.game_state.is_transitioning"
        @move="postMove($event)"
        @imove="postInterruptMove($event)"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import Player from "./Player.vue";
import FriendInfo from "./FriendInfo.vue";
import Tile from "./Tile.vue";
import Hand from "./Hand.vue";
import { GameStateResponse, IMove } from "../models/game_state";

export default Vue.extend({
  name: "Game",
  components: { Player, Tile, FriendInfo, Hand },
  props: {
    msg: String,
    gameID: String,
    playerNumber: Number,
  },
  data() {
    return {
      info: null, // GameStateResponse
      playerOptions: ["east", "south", "west", "north"],
      playerPos: ["right", "opposite", "leftplayer"],
      showWinningHand: false,
      lastShownWinningHand: -1
    };
  },
  computed: {
    prevailingWind(): string {
        return this.info ? this.playerOptions[this.info.game_state.prevailing_wind] :  "loading..."
    },
    remainingTileCount(): string {
      return this.info ? this.info.game_state.remaining_tiles.length :  "loading...";
    },
    playerWind(): string {
      if (!this.info) {
        return ".....loading"
      }
      const pos = this.info.game_state.starter - this.playerNumber;
      switch (pos) {
        case 0:
          return "east";
        case 2:
          return "west";
        case -2:
          return "west";
        case 1:
          return "north";
        case -3:
          return "north";
        case -1:
          return "south";
        case 3:
          return "south";
      }
      return "BUG. Pls screenshot and file an issue on github.com/slai11/mahjong"
    },
    winningTile(): number {
      // if the player wins on turn 1, there is no last discarded nor last drawn.
      // the player is a god of mahjong
      return this.info ? this.info.game_state.last_winning_hand.winning_tile : -1
    },
    turnNumber(): number {
      return this.info ? this.info.game_state.turn_number : null;
    },
    friendOrder(): number[] {
      switch (this.playerNumber) {
        case 0:
          return [1, 2, 3];
        case 1:
          return [2, 3, 0];
        case 2:
          return [3, 0, 1];
        case 3:
          return [0, 1, 2];
        default:
          return [];
      }
    }
  },
  mounted() {
    const fn = () => this.getGameState();
    fn();
    setInterval(function() {
      fn();
    }, 2000);
  },
  methods: {
    getGameState() {
      axios
        .get<GameStateResponse>(
          `${process.env.VUE_APP_BACKEND_URL}/game_state?game_id=${this.gameID}`
        )
        .then(response => {
          this.info = response.data
          // only update dialog status when its not true, else you will close the box while player is still viewing
          if (!this.showWinningHand) {
            this.showWinningHand = this.info.game_state.turn_number === this.info.game_state.last_winning_turn && this.info.game_state.turn_number !== this.lastShownWinningHand
            this.lastShownWinningHand = this.info.game_state.last_winning_turn
          }
        });
    },
    postMove(event: IMove) {
      event["turn_number"] = this.turnNumber;
      axios
        .post<GameStateResponse>(`${process.env.VUE_APP_BACKEND_URL}/move`, {
          "game_id": this.gameID,
          move: event
        })
        .then(response => {
          this.info = response.data;
        })
        .catch(error => {
          console.log(error);
          alert(`Move not allowed: ${error}`);
        });
    },
    postInterruptMove(event: IMove) {
      event["turn_number"] = this.turnNumber;
      event["tile"] = this.info.game_state.last_discarded_tile;
      console.log(event);
      axios
        .post<GameStateResponse>(`${process.env.VUE_APP_BACKEND_URL}/move`, {
          "game_id": this.gameID,
          move: event
        })
        .then(response => {
          this.info = response.data;
        })
        .catch(error => {
          console.log(error);
          alert(`Move not allowed: ${error}`);
        });
    }
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.discard {
  grid-area: main;
  outline: 2px solid green;
  width: 700px;
  height: 700px;
  display: flex; /* or inline-flex */
  flex-wrap: wrap;
  flex-direction: row;
  justify-content: flex-start;
  align-content: space-between;
}
.oppositeplayer {
  height: 100px;
  width: 700px;
  grid-area: oppositeplayer;
  place-self: center;
}
.leftplayer {
  grid-area: leftplayer;
  height: 100px;
  width: 700px;
  place-self: center;
  -webkit-transform: rotate(90deg);
  -moz-transform: rotate(90deg);
  -o-transform: rotate(90deg);
  -ms-transform: rotate(90deg);
  transform: rotate(90deg);
}
.rightplayer {
  grid-area: rightplayer;
  height: 100px;
  width: 700px;
   place-self: center;
  -webkit-transform: rotate(270deg);
  -moz-transform: rotate(270deg);
  -o-transform: rotate(270deg);
  -ms-transform: rotate(270deg);
  transform: rotate(270deg);
}

.player {
  grid-area: player;
}

.container {
  display: grid;
  grid-template-columns: 150px 700px 150px;
  grid-template-rows: 150px 700px 300px;
  grid-template-areas:
    ". oppositeplayer  ."
    "leftplayer main rightplayer"
    "player player player";
}

h3 {
  margin: 40px 0 0;
}
</style>
