<template>
  <div class="hand">
    <h2>Your Hand</h2>
    <div v-if="this.hand" class="container">
        <div v-for="tile of this.sortedHand" :key=tile.id>
            <Tile @click.native="discard(tile)" :value=tile.value :suit=tile.suit :id=tile.id />
        </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Tile  from "./Tile.vue";
import {ITile} from "../models/game_state"

export default Vue.extend({
  name: 'Hand',
  props: ['hand'],
  components: {
      Tile
  },
  computed: {
    sortedHand(): [] {
      const sortedHand = this.hand
      return sortedHand.sort((a: ITile, b: ITile) => {return a.id - b.id})
    }
  },
  methods: {
    discard(t: ITile) {
      this.$emit('move', {
        tile: t,
        action: 0
      })
    }
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.container {
  display: flex; /* or inline-flex */
  flex-direction: row;
  justify-content: space-evenly
}
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
