<template>
  <div class="hand">
    <Displayed v-bind:displayed="this.info ? this.info.displayed : null" />
    <Hand :hand="this.info ? this.info.hand : null" @move="emitMove" :mark="lastDrawnTile"/>


    <div v-if="this.transiting">
      <v-btn v-if="this.player_turn === this.player_number" @click="emitInterruptMove(1)" :disabled="drawDisabled">Draw</v-btn>
      <v-btn v-if="this.info.can_pong" @click="emitInterruptMove(5)" >Pong</v-btn>
      <v-btn v-if="this.info.can_gong" @click="emitInterruptMove(6)" >Gong</v-btn>

      <div v-if="this.player_turn === this.player_number &&  this.info.can_eat" >
        <v-tooltip bottom>
          <template v-slot:activator="{ on }">
            <v-btn v-on="on" @click="emitInterruptMove(2)" >Eat</v-btn>
          </template>
          <span>Example: 5, _, 7, eating a 6 would be "eat"</span>
        </v-tooltip>
      </div>

      <div v-if="this.player_turn === this.player_number &&  this.info.can_eat_right" >
        <v-tooltip bottom>
          <template v-slot:activator="{ on }">
            <v-btn v-on="on" @click="emitInterruptMove(3)" >Eat Right</v-btn>
          </template>
          <span>Example: 5, 6, _, eating a 7 would be "eat right"</span>
        </v-tooltip>
      </div>

      <div v-if="this.player_turn === this.player_number &&  this.info.can_eat_left" >
        <v-tooltip bottom>
          <template v-slot:activator="{ on }">
            <v-btn v-on="on" @click="emitInterruptMove(4)" >Eat Left</v-btn>
          </template>
          <span>Example: _, 5, 6, eating a 4 would be "eat left"</span>
        </v-tooltip>
      </div>
    </div>
    <v-btn v-if="this.transiting || this.player_turn === this.player_number"  @click="dialog = true">Hu</v-btn>

    <v-row justify="center">
      <v-dialog v-model="dialog" persistent max-width="400">
        <v-card>
          <v-card-title class="headline">HU!</v-card-title>
          <v-card-text>Are you sure you want to hu?</v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="green darken-1" text @click="dialog = false">Oh shoot, nearly zha hu.</v-btn>
            <v-btn color="green darken-1" text @click="emitInterruptMove(8)">Yes.</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>

    <div v-if="this.player_turn === this.player_number">
      <div v-for="(value, name) in this.info.inner_gong_map" :key="name">
        <li v-if="value === 4">
          <v-btn @click="emitInnerGong(7, innerGongTile(name).id)">Gong {{innerGongTile(name).name}}</v-btn>
        </li>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Hand from "./Hand.vue";
import Displayed from "./Displayed.vue";
import { TileNameMap, UniqueTile } from "../models/tile";

export default Vue.extend({
  name: "Player",
  props: ["info", "player_turn", "player_number", "transiting", "turnNumber"],
  components: {
    Hand,
    Displayed
  },
  data() {
    return {
      dialog: false,
      drawDisabled: false,
      lastPlayerTurn: -1,
    };
  },
  updated() {
    // 3rd boolean condition ensures only 1 block per game turn
    if (this.transiting && this.player_turn === this.player_number && this.lastPlayerTurn !== this.turnNumber) {
      console.log(`trigger block on draw for turn ${this.turnNumber}`)
      this.lastPlayerTurn = this.turnNumber; // trigger only on first instance of update each turn
      this.drawDisabled = true;
      // set draw button to clickable
      const fn = () => {
        this.drawDisabled = false;
      }

      setTimeout(function() {
        fn();
      }, 3000);
    }
  },
  computed: {
    lastDrawnTile(): number {
      return this.info.last_drawn_tile ? this.info.last_drawn_tile : -1
    },
    possibleGone(): number[][] {
      return this.info.inner_gong_map.filter((l: number[]) => {
        l.length === 4;
      });
    }
  },
  methods: {
    emitMove(event: object) {
      if (this.player_number === this.player_turn && !this.transiting) {
        this.$emit(
          "move",
          Object.assign(event, { player: this.player_number })
        );
      }
    },
    emitInterruptMove(action: number) {
      this.$emit("imove", { action: action, player: this.player_number });
      this.dialog = false;
    },
    innerGongTile(name: number): UniqueTile {
      return TileNameMap[name];
    },
    emitInnerGong(action: number, id: string) {
      this.$emit("move", {
        player: this.player_number,
        action: action,
        tile: id
      });
    }
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
a {
  color: #42b983;
}
</style>
