<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <h1>TableSwim Mahjong</h1>
    <Lobby v-if="playerNumber === -1"
      :gameID="gameID" 
      @change="gameID = $event"
      @registered="playerNumber = $event"/>
    <Game v-else :gameID=gameID :playerNumber=playerNumber msg="Table"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Lobby from './components/Lobby.vue';
import Game from './components/Game.vue';
import uniqueIdGenerator from "./util/uniqueIdGenerator";

export default Vue.extend({
  name: 'App',
  components: {
    Lobby,
    Game,
  },
  data() {
    return {
      playerNumber: -1,
      gameID: uniqueIdGenerator(),
    }
  }
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
