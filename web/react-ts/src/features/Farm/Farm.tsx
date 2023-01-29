import * as React from "react";
import { useMutation, useQuery } from "@tanstack/react-query";
import { FarmServiceApi, V1Farm, V1Barn } from "chicken-farmer-service/api";
import { Configuration } from "chicken-farmer-service/configuration";
import { Barn } from "./Barn/Barn";
import { useEffect, useRef, useState } from "react";
import { SetupFarmSSE } from "./SSE";

// const farmValidator = z.object({
//   name: z.string(),
//   day: z.number(),
//   golden_eggs: z.number(),
// });

const farmServiceApi = new FarmServiceApi(
  new Configuration({ basePath: "http://localhost:8081" })
);

export const Farm = () => {
  const [farm, _setFarm] = useState<V1Farm>({
    barns: new Array<V1Barn>(),
    day: 0,
    goldenEggs: 0,
    name: ""
  });

  // Need a ref to access the current farm inside the listeners created for SSE.
  const farmRef = useRef(farm);

  // Override _setFarm to keep the ref always updated.
  function setFarm(farm: V1Farm) {
    farmRef.current = farm;
    _setFarm(farm);
  }

  const { data, error, isLoading, isError, isFetched } = useQuery(
    ["getFarm"],
    async () => farmServiceApi.farmServiceFarmDetails()
  );

  const buyBarn = useMutation({
    mutationFn: () => farmServiceApi.farmServiceBuyBarn({})
  });

  useEffect(() => {
    if (!data) {
      return;
    }

    setFarm(data.data.farm);
  }, [isFetched]);

  useEffect(() => {
    const eventSource = SetupFarmSSE(setFarm, farmRef);
    return () => eventSource.close();
  }, []);

  return isLoading ? (
    <div>Loading your farm</div>
  ) : isError ? (
    <div>Errored: {error.message}</div>
  ) : (
    <div className="flex h-screen">
      <div className="flex flex-col items-center basis-2/12 p-4">
        <h1 className="text-center text-3xl font-extrabold mb-3">
          {farm.name}
        </h1>
        <span>
          <label>Current Day:</label> <span key={farm.day}>{farm.day}</span>
        </span>

        <span>
          <label>Golden Eggs:</label>{" "}
          <span key={farm.goldenEggs}>{farm.goldenEggs}</span>
        </span>

        <button className="btn-primary mt-4" onClick={() => buyBarn.mutate()}>
          Buy Barn
        </button>

        {error && <span className="mt-4 text-red-700">{error.message}</span>}
      </div>
      <div className="barns">
        {farm.barns.map((barn, index: number) => {
          return <Barn key={index} barn={barn} day={farm.day} />;
        })}
      </div>
    </div>
  );
};
