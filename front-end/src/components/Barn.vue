<template>
  <div class="barn">
    <img src="@/assets/barn.png" alt="barn" width="200"/>
    <button @click="buyAutoFeeder" v-if="!barn.autoFeeder">Buy AutoFeeder(100)</button>
    <div class="stats">
      <span><label>Feed:</label> {{ barn.feed }}</span>
      <span><label>Chickens:</label> {{ barn.chickens.length }}</span>
    </div>
    <div class="actions">
      <button @click="buyFeed">Buy Feed</button>
      <button @click="buyChicken">Buy Chicken</button>
    </div>
    <div class="chickens">
      <Chicken
        v-for="chicken in barn.chickens"
        :key="chicken.id"
        :chicken="chicken"
        :api="api"
        :currentDay="currentDay"
        @gold-egg-laid="handleGoldEggLaid"
        @feed-spent="handleFeedSpent"
        @chicken-sold="handleChickenSold"
        @error="handleChickenError"
      />
    </div>
  </div>
</template>

<script>

import Chicken from "@/components/Chicken"

const FEED_PER_PURCHASE = 10;
const FEED_COST = 1;
const CHICKEN_COST = 1;
const AUTOFEEDER_COST = 100;

export default {
  name: "barn",
  components: {
    Chicken
  },
  props: {
    barn: { type: Object, default: undefined},
    api: { type: Object, default: undefined },
    currentDay: 0
  },
  data() {
    return {
    };
  },
  async created() {
    this.startAutoFeeder();
  },
  methods: {
    async buyChicken() {
      try {
        const chicken = (await this.api.buyChicken(this.barn.id)).data;
        this.barn.chickens.push(chicken);
        this.$emit('gold-egg-spent', CHICKEN_COST);
      } catch (error) {
        this.$emit('error', "Couldn't buy a chicken, make sure you have enough gold eggs!");
      }
    },
    async buyFeed() {
      try {
        await this.api.buyFeed(this.barn.id);
        this.barn.feed += FEED_PER_PURCHASE;
        this.$emit('gold-egg-spent', FEED_COST);
      } catch (error) {
        this.$emit('error', "Couldn't buy feed, make sure you have enough gold eggs!");
      }
    },
    async buyAutoFeeder() {
      try {
        await this.api.buyAutoFeeder(this.barn.id);
        this.barn.autoFeeder = true;
        this.startAutoFeeder();
        this.$emit('gold-egg-spent', AUTOFEEDER_COST);
      } catch (error) {
        this.$emit('error', "Couldn't buy AutoFeeder, make sure you have enough gold eggs!");
      }
    },
    async startAutoFeeder() {
      if (this.barn.autoFeeder) {
        setInterval(async () => {
          let chickensToFeed = [];
          this.barn.chickens.forEach(
            async chicken => {
              if (this.barn.feed - chickensToFeed.length > 0 && this.currentDay > chicken.restingUntil) {
                chickensToFeed.push(chicken.id)
              }
            }
          );

          if (chickensToFeed.length > 0) {
            const results = (await this.api.bulkFeedChicken(chickensToFeed)).data;
            results.forEach(result => {
              let chicken = this.barn.chickens.find(chicken => chicken.id === result.id);
              this.barn.feed--;
              chicken.eggsLaid++;
              chicken.restingUntil = result.restingUntil;
              if (result.laidGoldEgg) {
                this.$emit('gold-egg-laid');
                chicken.goldEggsLaid++;
              }
            });
          }
        }, 5000)
      }
    },
    handleGoldEggLaid() {
      this.$emit('gold-egg-laid');
    },
    handleFeedSpent(amount) {
      this.barn.feed -= amount || 1
    },
    handleChickenSold(chickenId) {
      this.barn.chickens.find((chicken, index) => {
        if (chicken.id === chickenId) {
          this.barn.chickens.splice(index, 1);
          return true
        }
        return false;
      });
    },
    handleChickenError(error) {
      this.$emit('error', error);
    }
  }
};
</script>
<style scoped lang="scss">
  .barn {
    border: 1px dashed lightgray;
    flex-direction: column;
    align-items: center;
    min-height: 100vh;
    padding: 20px;
    display: inline-flex;

    .stats {
      margin-top: 20px;

      justify-content: space-between;
      display: flex;
      width: 100%;
    }

    .actions {
      justify-content: space-between;
      display: flex;
      width: 100%;
    }

    .chickens {
      margin-top: 50px;
      width: 100%;
    }
  }

</style>