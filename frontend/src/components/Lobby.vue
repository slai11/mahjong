<template>
  <div class="lobby">
    <h1>TableSwim Mahjong</h1>
    <h3>Step 1: Use the generated ID or paste your friend's table ID</h3>
    <h3>Step 2: Click "start game"</h3>
    <br>
    <input v-model="gameid" placeholder="Enter game ID" />
    <button v-on:click="enterGame">Start Game</button>

    <div v-if="gameFull">
      <h3>Game is full. Did you disconnect? Select player to rejoin as:</h3>
      <div v-for="(player, idx) in playerOptions" :key="idx">
        <input type="radio" id="idx" :value="idx" v-model="playerNumber" />
        <label for="idx">{{ player }}</label>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import { GameStateResponse } from "../models/game_state";

interface PlayerRegistrationResp {
  assigned_number: number;
}

export default Vue.extend({
  name: "Lobby",
  props: {
    gameID: String
  },
  data() {
    return {
      gameid: this.gameID,
      playerNumber: -1,
      gameFull: false,
      playerOptions: [1, 2, 3, 4]
    };
  },
  watch: {
    playerNumber: function(val) {
      this.$emit("registered", val);
    }
  },
  methods: {
    enterGame() {
      axios
        .get<GameStateResponse>(
          `${process.env.VUE_APP_BACKEND_URL}/game_state?game_id=${this.gameid}`
        )
        .then(response => {
          this.$emit("change", response.data.id);
          axios
            .post<PlayerRegistrationResp>(
              `${process.env.VUE_APP_BACKEND_URL}/player?game_id=${this.gameid}`
            )
            .then(
              response => {
                // updating player number will close lobby
                this.$emit("registered", response.data.assigned_number);
              }
            )
            .catch(error => {
              console.log(error);
              this.gameFull = true;
            });
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
</style>
