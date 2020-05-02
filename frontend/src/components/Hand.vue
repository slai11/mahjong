<template>
  <div class="hand">
    <div v-if="this.hand" class="container">
        <div v-for="tile of this.hand" :key=tile.id>
            <Tile :value=tile.value :suit=tile.suit :id=tile.id />
        </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Tile  from "./Tile.vue";

export default Vue.extend({
  name: 'Hand',
  props: ['hand'],
  components: {
      Tile
  },
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
