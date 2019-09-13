<template>
  <div class="barn-registration-center">
    <div class="owner-info">
      <h1>My Farm</h1>
      <span><label>Session ID: <br><input v-model="sessionId" @change="loadSession"/></label></span>
      <span><label>currentDay:</label> {{ currentDay }}</span>
      <span><label>Golden Eggs:</label> {{ goldEggCount }}</span>
      <button @click="registerBarn">Buy Barn</button>


      <span class="error" v-if="error">{{ error }}</span>
    </div>
    <div class="barns">
      <Barn
        v-for="(barn, index) in barns"
        :key="index"
        :barn="barn"
        :api="api"
        :currentDay="currentDay"
        :onError="showError"
        @gold-egg-laid="handleGoldEggLaid"
        @gold-egg-spent="handleGoldEggSpent"
        @error="showError"
      />
    </div>
  </div>
</template>

<script>
import MicroServicesAPI from "@/classes/API/MicroServicesAPI";
import Barn from "@/components/Barn"

const BARN_COST = 10;

export default {
  name: "home",
  components: {
    Barn
  },
  data() {
    return {
      currentDay: 0,
      goldEggCount: 0,
      barns: [],
      error: undefined,
      api: undefined,
      sessionId: undefined,
    };
  },
  async created() {
    this.sessionId = MicroServicesAPI.randomUserId();
    this.loadSession();

    let ws = new WebSocket("ws://localhost:8083/users/ws");
    ws.onclose = () => {
      ws = null;
    };
    ws.onmessage = evt => {
      this.currentDay = parseInt(evt.data);
    };
    ws.onerror = console.log;
  },
  methods: {
    async loadSession() {
      // TODO validate sessionId as Hex
      this.api = new MicroServicesAPI(this.sessionId);

      await this.fetchGoldEggs();
      await this.fetchBarns();
      this.fetchChickens();
    },
    async registerBarn() {
      if (this.barns.length > 0 && this.goldEggCount < BARN_COST) {
        this.showError("Not enough gold eggs!");
        return;
      }

      let barn = (await this.api.registerNewBarn()).data;
      barn.chickens = [];
      this.barns.push(barn);

      if (this.barns.length > 1) {
        this.goldEggCount -= BARN_COST;
      }

      this.fetchChickens();
    },
    async fetchGoldEggs() {
      this.goldEggCount = (await this.api.getGoldEggs()).data.goldEggCount;
    },
    async fetchBarns() {
      let barns = (await this.api.getBarns()).data;
      this.barns = barns.map(barn => {
        barn["chickens"] = [];
        return barn;
      });
    },
    async fetchChickens() {
      const chickens = (await this.api.getChickens()).data;
      chickens.forEach(chicken => {
        let barn = this.barns.find(barn => barn.id === chicken.belongsToBarn);
        if (!barn.chickens.find(barnChicken => barnChicken.id === chicken.id)) {
          barn.chickens.push(chicken)
        }
      });
    },
    handleGoldEggLaid() {
      this.goldEggCount++;
    },
    handleGoldEggSpent(amount) {
      this.goldEggCount-= amount || 1;
    },
    showError(errorMessage) {
      this.error = errorMessage;

      setTimeout(() => {
        this.error = undefined;
      }, 5000);
    }
  }
};
</script>
<style scoped lang="scss">
  .barn-registration-center {
    display: flex;
    height: 100vh;

    .owner-info {
      flex-direction: column;
      align-items: center;
      min-width: 320px;
      padding: 20px;
      display: flex;
      width: 10vw;

      input {
        width: 100%;
      }

      button {
        margin-top: 20px;
        padding: 10px 30px;
      }

      .error {
        margin-top: 20px;
        color: red;
      }
    }
  }
</style>