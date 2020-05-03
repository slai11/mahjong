<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <Lobby v-model=gameID msg="TableSwim Mahjong"/>
    <Game :gameID=gameID msg="Table"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Lobby from './components/Lobby.vue';
import Game from './components/Game.vue';

export default Vue.extend({
  name: 'App',
  components: {
    Lobby,
    Game
  },
  data() {
    return {
      gameID: ""
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
