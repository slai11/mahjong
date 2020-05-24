<template>
  <div class="lobby">
    <h1>TableSwim Mahjong</h1>
    <h3>Step 1: Use the generated ID or paste your friend's table ID</h3>
    <h3>Step 2: Click "start game"</h3>
    <br>

    <v-row justify="center" align="center">
      <v-spacer></v-spacer>
      <v-col cols="6" md="4">
        <v-text-field outlined v-model="gameid" label="Game ID" />
        <v-btn @click="enterGame">Start Game</v-btn>
      </v-col>
      <v-spacer></v-spacer>
    </v-row>

    <div v-if="gameFull">
      <h3>Game is full. Did you disconnect? Select player to rejoin as:</h3>
      <div v-for="(player, idx) in playerOptions" :key="idx">
        <input type="radio" id="idx" :value="idx" v-model="playerNumber" />
        <label for="idx">{{ player }}</label>
      </div>
    </div>
    <br>
    <div>
      <h2>New Changes</h2>
      <ul>
         <li>(Added 2020-05-24): 3 second disable on "Draw" button to allow other players to pong/gong/hu.</li>
         <li>(Added 2020-05-24): Added tooltip for "eat", "eat left" and "eat right" which pop-up an example for each case.</li>
      </ul>
    </div>
    <br>

    <div>
      <h3>Points to note:</h3>
      <ul>
         <li>1. "Move not allowed" occurs when more than 1 player selects an action in the same turn (draw and pong). First come first served.</li>
         <li>2. Report bugs and feature requests here: https://forms.gle/NzR2oSbg6Kj548jB9</li>
         <li>3. More technical users can submit an issue here: https://github.com/slai11/mahjong</li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import { GameStateResponse } from "../models/game_state";
import uniqueIdGenerator from "../util/uniqueIdGenerator";

interface PlayerRegistrationResp {
  assigned_number: number;
}

export default Vue.extend({
  name: "Lobby",
  props: {
    routed: String,
  },
  data() {
    return {
      gameid: uniqueIdGenerator(),
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
  mounted() {
    if (this.routed) {
      this.gameid = this.routed;
      this.enterGame();
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
