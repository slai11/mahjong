<template>
  <div class="lobby">
    <h1>{{ msg }}</h1>
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
    msg: String,
    gameID: String
  },
  data() {
    return {
      gameid: this.gameID,
      playerNumber: -1,
      gameFull: false,
      playerOptions: ["east", "south", "west", "north"]
    };
  },
  watch: {
    playerNumber: function(val) {
      this.$emit("registered", val);
    }
  },
  methods: {
    enterGame(event: any) {
      axios
        .get<GameStateResponse>(
          `http://localhost:80/game_state?game_id=${this.gameid}`
        )
        .then(response => {
          this.$emit("change", response.data.id);
          axios
            .post<PlayerRegistrationResp>(
              `http://localhost:80/player?game_id=${this.gameid}`
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
