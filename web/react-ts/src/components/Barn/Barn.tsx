import "./Barn.scss";
import barnImg from "./barn.png";
import { Chicken } from "../Chicken/Chicken";
import { useMutation, useQuery } from "@tanstack/react-query";
import { FarmServiceApi, V1Barn } from "chicken-farmer-service/api";
import { Configuration } from "chicken-farmer-service/configuration";

const farmServiceApi = new FarmServiceApi(
  new Configuration({ basePath: "http://localhost:8081" })
);

const buyChicken = (barnId: string): { error: any } => {
  const { error } = useQuery(["buyFeed"], async () =>
    farmServiceApi.farmServiceBuyChicken({ barnId })
  );

  return { error };
};

export const Barn = (props: { barn: V1Barn; day: number }) => {
  const buyFeed = useMutation(["buyFeed"], async (barnId: string) =>
    farmServiceApi.farmServiceBuyFeedBag(barnId, { amount: 1 })
  );

  const buyChicken = useMutation(["buyChicken"], async (barnId: string) =>
    farmServiceApi.farmServiceBuyChicken({ barnId })
  );

  return (
    <div className="barn">
      <img src={barnImg} alt="barn" width="200" />
      {!props.barn.hasAutoFeeder && <button>Buy AutoFeeder</button>}
      <div className="stats">
        <span>
          <label>Feed:</label> {props.barn.feed}
        </span>
        <span>
          <label>Chickens:</label> {props.barn.chickens?.length}
        </span>
      </div>
      <div className="actions">
        <button onClick={() => buyFeed.mutate(props.barn.id ?? "")}>
          Buy Feed
        </button>
        <button onClick={() => buyChicken.mutate(props.barn.id ?? "")}>
          Buy Chicken
        </button>
      </div>
      <div className="chickens">
        {props.barn?.chickens?.map((chicken, index: number) => {
          return <Chicken key={index} chicken={chicken} day={props.day} />;
        })}
      </div>
    </div>
  );
};
