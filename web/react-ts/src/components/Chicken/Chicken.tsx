import "./Chicken.scss";
import { FarmServiceApi, V1Chicken } from "chicken-farmer-service/api";
import { Configuration } from "chicken-farmer-service/configuration";
import cakeImg from "./assets/cake.gif";
import clockImg from "./assets/clock.png";
import eggImg from "./assets/egg.png";
import goldEggImg from "./assets/gold_egg.png";
import { useEffect, useState } from "react";
import { useMutation } from "@tanstack/react-query";

enum Action {
  STANDING_LEFT = "standing-left",
  STANDING_RIGHT = "standing-right",
  SITTING_LEFT = "sitting-left",
  SITTING_RIGHT = "sitting-right",
  FEEDING = "feeding",
}

// TODO have this on a shared file instead of one per chicken?
const actions = Object.values(Action);

enum Medal {
  BRONZE = "bronze",
  SILVER = "silver",
  GOLD = "gold",
  GOLDPLUS = "gold-plus",
}

interface State {
  action: string;
  block: number;
  medal: Medal;
}
const farmServiceApi = new FarmServiceApi(
  new Configuration({ basePath: "http://localhost:8081" })
);

export const Chicken = (props: { chicken: V1Chicken; day: number }) => {
  const [action, setAction] = useState(Action.STANDING_LEFT);

  const feedChicken = useMutation(["feedChicken"], async (chickenId: string) =>
    farmServiceApi.farmServiceFeedChicken(chickenId, {})
  );

  useEffect(() => {
    const interval = setInterval(() => {
      if (props.chicken.restingUntil ?? 0 >= props.day) {
        setAction(Action.FEEDING);
        return;
      }
      setAction(actions[Math.floor(Math.random() * (actions.length - 1))]);
    }, Math.max(2000, Math.random() * 5000)); // random wait between [2s, 5s]);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="chicken">
      {/*<div className={`medal-img ${this.state.medal}`} />*/}
      <div className={`chicken-img ${action}`} />
      <div className="chicken-stats">
        <span>
          <img src={cakeImg} alt="birthday" width="30" />{" "}
          {props.chicken.dateOfBirth}
        </span>
        <span>
          <img src={eggImg} alt="eggs laid" width="20" />{" "}
          {props.chicken.normalEggsLaid}
        </span>
        <span>
          <img src={goldEggImg} alt="goldeggs laid" width="20" />{" "}
          {props.chicken.goldEggsLaid}
        </span>
        <span>
          <img src={clockImg} alt="resting until" width="20" />{" "}
          {Math.max(props.chicken.restingUntil ?? 0 - props.day, 0)}
        </span>
      </div>
      <div className="actions">
        <button
          onClick={() => feedChicken.mutate(props.chicken.id ?? "")}
          disabled={(props.chicken.restingUntil ?? 0) >= props.day}
        >
          Feed
        </button>
        {/*<button onClick={this.sellChicken}>Sell</button>*/}
      </div>
    </div>
  );
};
