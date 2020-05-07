<template>
    <img class="tile" :src="`/tiles/${tileName}.jpg`">
</template>

<script lang="ts">
import Vue from "vue";
import {TileList} from "../models/tile"

export default Vue.extend({
  name: "Tile",
  props: {
    id: Number
  },
  computed: {
    tileName(): string {
      return TileList[this.id].id;
    }
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.tile {
  outline: 2px solid green;
  width: 50px;
  height: 80px;
}
</style>
