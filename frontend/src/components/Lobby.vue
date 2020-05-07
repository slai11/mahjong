<template>
  <div class="lobby">
    <h1>{{ msg }}</h1>
    <input v-model="input" placeholder="Enter game ID" />
    <button v-on:click="enterGame"> Start Game </button>
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
      input: uniqueIdGenerator(),
      gameStarted: false,
    };
  },
  methods: {
    enterGame(event: any) {
      axios
        .get<GameStateResponse>(
          `https://tableswim.com/game_state?game_id=${this.input}`
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
