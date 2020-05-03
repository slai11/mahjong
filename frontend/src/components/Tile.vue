<template>
  <div class="tile">
    <h3>{{ tileName }}</h3>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "Tile",
  props: {
    suit: Number,
    value: Number,
    id: Number
  },
  computed: {
    // do a lookup instead?
    tileName(): string {
      let sname = "";
      switch (this.suit) {
        case 0:
          sname = "Bamboo:" + (this.value + 1);
          break;
        case 1:
          sname = "Coin:" + (this.value + 1);
          break;
        case 2:
          sname = "Number:" + (this.value + 1);
          break;
        case 3:
          if (this.value == 0) {
            sname = "East-Wind"
          } else if (this.value == 1) {
            sname = "South-Wind"
          } else if (this.value == 2) {
            sname = "West-Wind"
          } else if (this.value == 3) {
            sname = "North-Wind"
          }
          break;
        case 4:
          if (this.value == 0) {
            sname = "Green-Dragon"
          } else if (this.value == 1) {
            sname = "Red-Dragon"
          } else if (this.value == 2) {
            sname = "White-Dragon"
          }
          break;
        case 5:
          if (this.value == 0) {
            sname = "Centipede"
          } else if (this.value == 1) {
            sname = "Chicken"
          } else if (this.value == 2) {
            sname = "Cat"
          } else if (this.value == 3) {
            sname = "Mouse"
          }
          break;
        case 6:
          sname = "Flower";
          break;
      }
      return sname;
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
