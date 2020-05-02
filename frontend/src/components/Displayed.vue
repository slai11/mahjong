<template>
  <div class="hand">
    <h1>Your Displayed</h1>
    <h3>{{ prettyDisplayed }}</h3>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Tile } from "../models/game_state";

export default Vue.extend({
  name: 'Displayed',
  props: ['displayed'], // list of list
  computed: {
      prettyDisplayed(): string[][] {
          if (this.displayed) {
              const displayed = this.displayed;
              console.log(displayed);
              return displayed.map((ts: Tile[]) => {
                  return ts.map((t: Tile) => {
                      switch (t.suit) {
                      case 0:
                          return "Bamboo:" + (t.value + 1);
                      case 1:
                          return "Coin:" + (t.value + 1);
                      case 2:
                          return "Number:" + (t.value + 1);
                      case 3:
                          return "Wind:" + (t.value + 1);
                      case 4:
                          return "Dragon:" + (t.value + 1);
                      case 5:
                          return "Animal:" + (t.value + 1);
                      case 6:
                          return "Flower:" + (t.value + 1);
                  }
                })
              });
          }
          return [[]];
      }
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
