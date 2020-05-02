<template>
  <div class="lobby">
    <h1>{{ msg }}</h1>
    <button v-on:click="getGameState">Load Game.</button>
    <Player v-bind:info=playerInfo />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import axios from "axios";
import {GameState, Tile} from "../models/game_state";
import Player from "./Player.vue";

interface GameStateResponse {
    game_state: GameState;
    id: string;
    created_at: string;
    updated_at: string;
}

export default Vue.extend({
  name: 'Lobby',
  components: { Player },
  props: {
    msg: String,
  },
  data() {
      return {
          info: null
      }
  },
  computed: {
      playerInfo(): Tile[] {
          return this.info ? this.info.game_state.player_map[0] : null
      },
  },
  methods: {
      getGameState() {
          axios.get<GameStateResponse>('http://localhost:80/game_state?game_id=12').then(response => (this.info = response.data))
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
