<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <Lobby msg="TableSwim Mahjong"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Lobby from './components/Lobby.vue';

export default Vue.extend({
  name: 'App',
  components: {
    Lobby
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
