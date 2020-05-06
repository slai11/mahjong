<template>
  <div class="friend">
    <Displayed v-bind:displayed="this.info ? this.info.displayed : null" />
    {{ "Player " +  this.player_number + " has " + this.info.hand.length + " tiles left"}}
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Displayed from "./Displayed.vue";

export default Vue.extend({
  name: "FriendInfo",
  props: ["info", "player_number"],
  components: { Displayed }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.friend {
    outline: 1px solid green;
}
</style>
