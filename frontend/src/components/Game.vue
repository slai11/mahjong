<template>
  <div class="Board">
    <h1>{{ msg }}: {{this.gameID}}</h1>

    <div v-if="playerNumber === -1">
      <div v-for="(player, idx) in playerOptions" :key="idx">
        <input type="radio" id="idx" :value="idx" v-model="picked" />
        <label for="idx">{{ player }}</label>
      </div>
      <span>Picked: {{ playerOptions[picked] }}</span>
    </div>
    {{ playerNumber }}

    <Player v-bind:info="playerInfo" :player_turn="playerTurn" :player_number="playerNumber" />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import { Tile } from "../models/game_state";
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
      picked: -1,
      playerOptions: ['east', 'south', 'west', 'north']
    };
  },
  watch: {
    picked: function (val) {
      this.playerNumber = val
      axios
        .get<GameStateResponse>(
          `http://localhost:80/game_state?game_id=${this.gameID}`
        )
        .then(response => (this.info = response.data))
        .catch((error) => {
            console.log(error);
            this.playerNumber = -1;
            alert(`${val} is selected, please choose another.`)
        });
    }
  },
  computed: {
    playerInfo(): Tile[] {
      return this.info ? this.info.game_state.player_map[this.playerNumber] : null;
    },
    playerTurn(): Tile[] {
      return this.info ? this.info.game_state.player_turn : null;
    }
  },
  methods: {
    getGameState() {
      axios
        .get<GameStateResponse>(
          `http://localhost:80/game_state?game_id=${this.gameID}`
        )
        .then(response => (this.info = response.data));
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
