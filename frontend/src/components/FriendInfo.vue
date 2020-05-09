<template>
  <div :class="(this.playerNumber === this.playerTurn)? 'active': 'friend'">
    <Displayed v-bind:displayed="this.info ? this.info.displayed : null" />
    {{ "Player " +  (this.playerNumber+1) + " has " + this.info.hand.length + " tiles left"}}
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Displayed from "./Displayed.vue";

export default Vue.extend({
  name: "FriendInfo",
  props: ["info", "playerNumber", "playerTurn"],
  components: { Displayed }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.friend {
    outline: 2px solid green;
}
.active {
    outline: 3px solid red;
}
</style>
