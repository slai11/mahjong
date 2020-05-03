<template>
  <div class="GameDashboard">
    <div v-if="playerNumber === -1">
      <div v-for="(player, idx) in playerOptions" :key="idx">
        <input type="radio" id="idx" :value="idx" v-model="playerNumber" />
        <label for="idx">{{ player }}</label>
      </div>
      <span>Picked: {{ playerOptions[playerNumber] }}</span>
    </div>
    <h3>Player: {{ playerNumber }}</h3>

    <div v-if="info" class="Board">
      <h3>Prevailing wind: {{info.game_state.prevailing_wind}}</h3>
      <h3>Dealer this round: {{info.game_state.starter}}</h3>
      <div v-for="(tile, id) in info.game_state.discarded_tiles" :key="id">
        <Tile :value="tile.value" :suit="tile.suit" :id="tile.id" />
      </div>

      <Tile
        v-if="info.game_state.last_discarded_tile"
        :value="info.game_state.last_discarded_tile.value"
        :suit="info.game_state.last_discarded_tile.suit"
        :id="info.game_state.last_discarded_tile.id"
      />

      <Player
        :info="info.game_state.player_map[this.playerNumber]"
        :player_turn="info.game_state.player_turn"
        :player_number="playerNumber"
        :transiting="info.game_state.is_transitioning"
        @move="postMove($event)"
        @imove="postInterruptMove($event)"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import Player from "./Player.vue";
import Tile from "./Tile.vue";
import { GameStateResponse, IMove } from "../models/game_state";

export default Vue.extend({
  name: "Game",
  components: { Player, Tile },
  props: {
    msg: String,
    gameID: String
  },
  data() {
    return {
      info: null, // GameStateResponse
      playerNumber: -1,
      playerOptions: ["east", "south", "west", "north"]
    };
  },
  watch: {
    playerNumber: function(val) {
      axios
        .post<GameStateResponse>(`http://localhost:80/player_select`, {
          "game_id": this.gameID,
          selection: val
        })
        .then(response => (this.info = response.data))
        .catch(error => {
          console.log(error);
          this.playerNumber = -1;
          alert(`${val} is selected, please choose another.`);
        });
    }
  },
  computed: {
    turnNumber(): number {
      return this.info ? this.info.game_state.turn_number : null;
    }
  },
  mounted() {
    const fn = () => this.getGameState();
    setInterval(function() {
      fn();
    }, 5000);
  },
  methods: {
    getGameState() {
      axios
        .get<GameStateResponse>(
          `http://localhost:80/game_state?game_id=${this.gameID}`
        )
        .then(response => (this.info = response.data));
    },
    postMove(event: IMove) {
      event["turn_number"] = this.turnNumber;
      axios
        .post<GameStateResponse>(`http://localhost:80/move`, {
          "game_id": this.gameID,
          move: event
        })
        .then(response => {
          this.info = response.data;
          console.log(this.info);
        })
        .catch(error => {
          console.log(error);
          alert(`Move not allowed`);
        });
    },
    postInterruptMove(event: IMove) {
      event["turn_number"] = this.turnNumber;
      event["tile"] = this.info.game_state.last_discarded_tile;
      console.log(event);
      axios
        .post<GameStateResponse>(`http://localhost:80/move`, {
          "game_id": this.gameID,
          move: event
        })
        .then(response => {
          this.info = response.data;
          console.log(this.info);
        })
        .catch(error => {
          console.log(error);
          alert(`Move not allowed`);
        });
    }
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.Board {
  display: flex; /* or inline-flex */
  flex-wrap: wrap;
  flex-direction: row;
  justify-content: flex-start;
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
