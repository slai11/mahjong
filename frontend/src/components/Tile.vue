<template>
    <!-- <img class="tile" :src="`/tiles/${tileName}.jpg`"> -->
    <div class="tile">
      {{tileName}}
    </div>
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
  outline: 1px solid green;
  width: 70px;
  height: 100px;
}
</style>
