<template>
  <div class="chicken">
    <div class="medal-img" :class="medal"/>
    <div class="chicken-img" :class="action"/>
    <div class="chicken-stats">
      <span><img src="@/assets/cake.gif" alt="birthday" width="30"/> {{ chicken.purchaseDay }}</span>
      <span><img src="@/assets/egg.png" alt="eggs laid" width="20"/> {{ chicken.eggsLaid }}</span>
      <span><img src="@/assets/gold_egg.png" alt="goldeggs laid" width="20"/> {{ chicken.goldEggsLaid }}</span>
      <span><img src="@/assets/clock.png" alt="resting until" width="20"/> {{ Math.max(chicken.restingUntil - currentDay, 0) }}</span>
    </div>
    <div class="actions">
      <button @click="feedChicken" :disabled="chicken.restingUntil - currentDay > 0">Feed</button>
      <button @click="sellChicken">Sell</button>
    </div>
  </div>
</template>

<script>
const Action = {
  STANDING_LEFT:  "standing-left",
  STANDING_RIGHT: "standing-right",
  SITTING_LEFT:   "sitting-left",
  SITTING_RIGHT:  "sitting-right",
  FEEDING:        "feeding"
};

const Medal = {
  BRONZE:   "bronze",
  SILVER:   "silver",
  GOLD:     "gold",
  GOLDPLUS: "gold-plus",
};

export default {
  name: "chicken",
  components: {},
  props: {
    barn: { type: Object, default: undefined },
    api: { type: Object, default: undefined },
    chicken: { type: Object, default: undefined },
    currentDay: 0,
  },
  data() {
    return {
      medal: undefined,
      action: Action.STANDING_LEFT
    };
  },
  async created() {
    const goldEggChance = this.chicken.goldEggChance;
    this.medal = Medal.BRONZE;
    if (goldEggChance > 33 && goldEggChance <= 66) {
      this.medal = Medal.SILVER
    } else if (goldEggChance > 66 && goldEggChance <= 90) {
      this.medal = Medal.GOLD
    } else if (goldEggChance > 90) {
      this.medal = Medal.GOLDPLUS
    }

    await this.checkFeedingState();
    this.randomizeAction();

    // check if still feeding every second
    setInterval(this.checkFeedingState, 1000);
  },
  methods: {
    randomizeAction(timeout) {
      setTimeout(() => {
        // if feeding then stop randomizing, except when timeout is set
        if (this.action === Action.FEEDING && timeout === undefined) {
          return
        }

        const actions = Object.values(Action);
        // pick random action except for the last one which should be the feeding one
        this.action = actions[Math.floor(Math.random() * (actions.length - 1))];
        this.randomizeAction();
      }, timeout || Math.max(2000, Math.random() * 5000)) // random wait between [2s, 5s]
    },
    async checkFeedingState() {
      if (this.chicken.restingUntil >= this.currentDay) {
        this.action= Action.FEEDING;
      } else {
        // only restart randomizing state if still feeding
        this.action === Action.FEEDING && this.randomizeAction(1);
      }
    },
    async feedChicken() {
      try {
        const response = (await this.api.feedChicken(this.chicken.id)).data;
        this.action = Action.FEEDING;
        this.chicken.restingUntil = response.restingUntil;
        this.chicken.eggsLaid++;
        this.$emit('feed-spent', 1);

        if (response.laidGoldEgg) {
          this.chicken.goldEggsLaid++;
          this.$emit('gold-egg-laid');
        }
      } catch(error) {
        this.$emit('error', "Couldn't feed proto, make sure there is feed available!");
      }
    },

    async sellChicken() {
      await this.api.sellChicken(this.chicken.id);

      this.$emit('proto-sold', this.chicken.id);
    }
  }
};
</script>
<style scoped lang="scss">
  .chicken {
    position: relative;
    width: 100%;

    .medal-img {
      background-image: url("../assets/medals.jpg");
      background-size: 200px;
      position: absolute;
      margin: 0 auto;
      height: 26px;
      width: 17px;
      left: 0;
      top: 10px;

      &.bronze {
        background-position: -43px -224px;
      }

      &.silver {
        background-position: -66px -224px;
      }

      &.gold {
        background-position: -91px -224px;
      }

      &.gold-plus {
        background-position: -141px -224px;
      }
    }

    .chicken-img {
      background-image: url("../assets/chicken.gif");
      margin: 0 auto;
      height: 50px;
      width: 50px;

      &.standing-left {
        background-position: -32px -148px;
      }

      &.standing-right {
        background-position: -108px -148px;
      }

      &.sitting-left {
        background-position: -182px -148px;
      }

      &.sitting-right {
        background-position: -32px -148px;
      }

      &.feeding {
        background-position: -320px -207px;
      }
    }

    .chicken-stats {
      justify-content: space-between;
      align-items: baseline;
      display: flex;
      width: 100%;

      span {
        vertical-align: center;
        height: 20px;
      }
    }

    .actions {
      justify-content: space-between;
      margin-top: 10px;
      display: flex;
    }
  }
</style>