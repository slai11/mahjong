<template>
  <div class="hand">
    <div v-if="this.hand" class="container">
        <div v-for="tile of this.sortedHand" :key=tile>
            <Tile v-if="mark === tile" @click.native="discard(tile)" :id="tile" :style="{outline: '2px double red'}"/>
            <Tile v-else @click.native="discard(tile)" :id="tile" />
        </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Tile  from "./Tile.vue";

export default Vue.extend({
  name: 'Hand',
  props: ['hand', 'mark'],
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
