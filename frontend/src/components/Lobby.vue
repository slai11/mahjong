<template>
  <div class="lobby">
    <h1>{{ msg }}</h1>
    <button v-on:click="createGame">Create a Game.</button>
    <input v-model="input" v-on:keyup.enter="enterGame($event)" placeholder="Enter game ID" />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import { GameStateResponse } from "../models/game_state";
import uniqueIdGenerator from "../util/uniqueIdGenerator";

export default Vue.extend({
  name: "Lobby",
  model: {
    prop: 'gameID',
    event: 'change'
  },
  props: {
    msg: String,
    gameID: String,
  },
  data() {
    return {
      info: null,
      input: "",
      gameStarted: false,
    };
  },
  methods: {
    createGame() {
      const id = uniqueIdGenerator();
      axios
        .get<GameStateResponse>(`http://localhost:80/game_state?game_id=${id}`)
        .then(response => (this.$emit('change', response.data.id)));
    },
    enterGame(event: any) {
      axios
        .get<GameStateResponse>(
          `http://localhost:80/game_state?game_id=${this.input}`
        )
        .then(response => (this.$emit('change', response.data.id)));
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
