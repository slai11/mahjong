<template>
  <div class="hand">
    <h1>Your Displayed Sets</h1>
    <Displayed v-bind:displayed="this.info ? this.info.displayed : null" />
    <h1>Your Hand</h1>
    <Hand v-bind:hand="this.info ? this.info.hand : null" />

    <button v-if="this.info ? this.player_turn === this.player_number : false">Call</button>
    <button v-if="this.info ? this.info.can_eat : false">Eat</button>
    <button v-if="this.info ? this.info.can_eat_right : false">Eat Right</button>
    <button v-if="this.info ? this.info.can_eat_left : false">Eat Left</button>
    <button v-if="this.info ? this.info.can_pong : false">Pong</button>
    <button v-if="this.info ? this.info.can_gong : false">Gong</button>
    <!-- need to display all the possible inner gong options -->
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Hand from "./Hand.vue";
import Displayed from "./Displayed.vue";

export default Vue.extend({
  name: "Player",
  props: ["info", "player_turn", "player_number"],
  components: {
    Hand,
    Displayed
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
