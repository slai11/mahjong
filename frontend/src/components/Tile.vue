<template>
  <div class="tile">
    <h3>{{ tileName }}</h3>
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
h3 {
  margin: 40px 0 0;
}
.tile {
  outline: 2px solid green;
}
a {
  color: #42b983;
}
</style>
