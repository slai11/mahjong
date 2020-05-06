<template>
  <div v-if="this.displayed" class="container">
    <div v-for="(tile, idx) of [].concat(...this.displayed)" :key="idx">
      <Tile :id="tile" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Tile from "./Tile.vue";

export default Vue.extend({
  name: "Displayed",
  props: ["displayed"],
  components: {
    Tile
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.container {
  display: flex; /* or inline-flex */
  flex-direction: row;
  justify-content: center;
}
a {
  color: #42b983;
}
</style>
