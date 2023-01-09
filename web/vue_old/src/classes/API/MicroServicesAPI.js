import axios from "axios";

export class MicroServicesAPI {
  BASE_URL = "http://192.168.99.100:31479";
  BARN_MS_URL = `${this.BASE_URL}/barns`;
  CHICKEN_MS_URL = `${this.BASE_URL}/chickens`;
  USER_MS_URL = `${this.BASE_URL}/farmers`;

  constructor(_farmerId) {
    this.farmerId = _farmerId;
  }

  static randomFarmerId() {
    // got this from https://stackoverflow.com/questions/10726909/random-alpha-numeric-string-in-javascript
    let p = "ABCDEFabcdef0123456789";
    return [...Array(24)].reduce(
      (a) => a + p[~~(Math.random() * p.length)],
      ""
    );
  }

  async getJwtTokenConfig() {
    return {
      headers: { Authorization: `Bearer ${await this.getJwtToken()}` },
    };
  }

  async getJwtToken() {
    if (!this.jwtToken) {
      this.jwtToken = (
        await axios.get(`${this.USER_MS_URL}/login/${this.farmerId}`)
      ).data;
    }

    return this.jwtToken;
  }

  async getGoldEggs() {
    return await axios.get(
      `${this.USER_MS_URL}/getGoldEggs`,
      await this.getJwtTokenConfig()
    );
  }

  async registerNewBarn() {
    return await axios.get(
      `${this.BARN_MS_URL}/buy`,
      await this.getJwtTokenConfig()
    );
  }

  async getBarns() {
    return await axios.get(
      `${this.BARN_MS_URL}/`,
      await this.getJwtTokenConfig()
    );
  }

  async buyAutoFeeder(barndId) {
    return await axios.get(
      `${this.BARN_MS_URL}/${barndId}/buy/autoFeeder`,
      await this.getJwtTokenConfig()
    );
  }

  async buyFeed(barndId) {
    return await axios.get(
      `${this.BARN_MS_URL}/${barndId}/buy/feed`,
      await this.getJwtTokenConfig()
    );
  }

  async getChickens() {
    return await axios.get(
      `${this.CHICKEN_MS_URL}/`,
      await this.getJwtTokenConfig()
    );
  }

  async feedChicken(chickenId) {
    return await axios.get(
      `${this.CHICKEN_MS_URL}/${chickenId}/feed`,
      await this.getJwtTokenConfig()
    );
  }

  async bulkFeedChicken(chickenIds) {
    return await axios.post(
      `${this.CHICKEN_MS_URL}/bulkFeed`,
      { chickenIds },
      await this.getJwtTokenConfig()
    );
  }

  async buyChicken(barndId) {
    return await axios.get(
      `${this.CHICKEN_MS_URL}/buy/${barndId}`,
      await this.getJwtTokenConfig()
    );
  }

  async sellChicken(chickenId) {
    return await axios.get(
      `${this.CHICKEN_MS_URL}/${chickenId}/sell`,
      await this.getJwtTokenConfig()
    );
  }
}
