import "./Farm.scss";
import { useQuery } from "@tanstack/react-query";
import * as React from "react";
import { FarmServiceApi, V1Farm } from "chicken-farmer-service/api";
import { Configuration } from "chicken-farmer-service/configuration";
import { Barn } from "../../components/Barn/Barn";

// Used to update TanStack client.
// const queryClient = useQueryClient();

// const farmValidator = z.object({
//   name: z.string(),
//   day: z.number(),
//   golden_eggs: z.number(),
// });

const farmServiceApi = new FarmServiceApi(
    new Configuration({ basePath: "http://localhost:8081" })
);

const getFarm = (): { farm: V1Farm; error: any } => {
  const { data, error } = useQuery(["getFarm"], async () => farmServiceApi.farmServiceGetFarm());
  return { farm: data?.data.farm as V1Farm, error };
};

export const Farm = () => {
  const { farm, error } = getFarm();

  return (
    <div className="farm">
      <div className="farm-info">
        <h1>{farm?.name}</h1>
        <span>
          <label>Current Day:</label> <span>{farm?.day}</span>
        </span>

        <span>
          <label>Golden Eggs:</label> <span>{farm?.goldenEggs}</span>
        </span>

        <button>Buy Barn</button>

        {error && <span className="error">{error}</span>}
      </div>
      <div className="barns">
        {farm?.barns?.map((barn, index: number) => {
          return <Barn key={index} barn={barn} day={farm.day ?? 0} />;
        })}
      </div>
    </div>
  );
};
