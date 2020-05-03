<template>
  <div class="Board">
    <h1>{{ msg }}: {{this.gameID}}</h1>

    <div v-if="playerNumber === -1">
      <div v-for="(player, idx) in playerOptions" :key="idx">
        <input type="radio" id="idx" :value="idx" v-model="playerNumber" />
        <label for="idx">{{ player }}</label>
      </div>
      <span>Picked: {{ playerOptions[playerNumber] }}</span>
    </div>
    <h3>Player: {{ playerNumber }}</h3>

    <Player :info="info ? info.game_state.player_map[this.playerNumber] : null" :player_turn="playerTurn" :player_number="playerNumber" @move="postMove($event)"/>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import { IMove } from "../models/game_state";
import Player from "./Player.vue";
import { GameStateResponse } from "../models/game_state";

export default Vue.extend({
  name: "Game",
  components: { Player },
  props: {
    msg: String,
    gameID: String
  },
  data() {
    return {
      info: null,
      playerNumber: -1,
      playerOptions: ["east", "south", "west", "north"]
    };
  },
  watch: {
    playerNumber: function(val) {
      axios
        .post<GameStateResponse>(`http://localhost:80/player_select`, {
          "game_id": this.gameID,
          selection: val
        })
        .then(response => (this.info = response.data))
        .catch(error => {
          console.log(error);
          this.playerNumber = -1;
          alert(`${val} is selected, please choose another.`);
        });
    }
  },
  computed: {
    playerInfo(): object {
      return this.info
        ? this.info.game_state.player_map[this.playerNumber]
        : null;
    },
    playerTurn(): number {
      return this.info ? this.info.game_state.player_turn : null;
    },
    turnNumber(): number {
      return this.info ? this.info.game_state.turn_number : null;
    }
  },
  methods: {
    getGameState() {
      axios
        .get<GameStateResponse>(
          `http://localhost:80/game_state?game_id=${this.gameID}`
        )
        .then(response => (this.info = response.data));
    },
    postMove(event: IMove) {
      event["turn_number"] = this.turnNumber
      axios
        .post<GameStateResponse>(`http://localhost:80/move`, {
          "game_id": this.gameID,
          move: event
        })
        .then(response => {
          this.info = response.data
          console.log(this.info)
          })
        .catch(error => {
          console.log(error);
          alert(`Move not allowed`);
        });
    }
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
