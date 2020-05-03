<template>
  <div class="hand">
    <h2>Your Displayed</h2>
    <div v-if="this.displayed" class="container">
        <div v-for="(set, idx) of this.displayed" :key=idx>
          <div v-for="tile of set" :key=tile.id>
            <Tile :value=tile.value :suit=tile.suit :id=tile.id />
          </div>
        </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Tile  from "./Tile.vue";

export default Vue.extend({
  name: 'Displayed',
  props: ['displayed'], // list of list
  components: {
    Tile
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
