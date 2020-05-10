<template>
  <div id="app">
    <Lobby v-if="playerNumber === -1"
      :routed="gameid"
      @change="gameID = $event"
      @registered="playerNumber = $event"/>
    <Game v-else :gameID=gameID :playerNumber=playerNumber msg="Table"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Lobby from './components/Lobby.vue';
import Game from './components/Game.vue';

export default Vue.extend({
  name: 'Home',
  components: {
    Lobby,
    Game,
  },
  props: ['gameid'],
  data() {
    return {
      playerNumber: -1,
      gameID: this.gameid,
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
}
</style>
