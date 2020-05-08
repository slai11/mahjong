<template>
  <div class="hand">
    <h2>{{msg}}</h2>
    <div v-if="this.hand" class="container">
        <div v-for="tile of this.sortedHand" :key=tile>
            <Tile @click.native="discard(tile)" :id=tile />
        </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Tile  from "./Tile.vue";

export default Vue.extend({
  name: 'Hand',
  props: ['hand', 'msg'],
  components: {
      Tile
  },
  computed: {
    sortedHand(): [] {
      const sortedHand = this.hand
      return sortedHand.sort((a: number, b: number) => {return a - b})
    }
  },
  methods: {
    discard(t: number) {
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
  justify-content: center;
}
</style>
