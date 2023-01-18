import { useMutation, useQuery } from "@tanstack/react-query";
import * as React from "react";
import { FarmServiceApi, V1Farm } from "chicken-farmer-service/api";
import { Configuration } from "chicken-farmer-service/configuration";
import { Barn } from "../../components/Barn/Barn";
import { useEffect, useState } from "react";

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

const getFarm = (): { result: V1Farm; error: any } => {
  const { data, error } = useQuery(["getFarm"], async () =>
    farmServiceApi.farmServiceGetFarm()
  );
  return { result: data?.data.farm as V1Farm, error };
};

export const Farm = () => {
  const [farm, setFarm] = useState<V1Farm>({});
  const { result, error } = getFarm();

  const buyBarn = useMutation({
    mutationFn: () => farmServiceApi.farmServiceBuyBarn({})
  });

  useEffect(() => {
    setFarm(result);
  }, [result]);

  useEffect(() => {
    const eventSource = new EventSource("http://localhost:8083/event-feed", {
      withCredentials: false
    });

    eventSource.onmessage = (event) => {
      const data = JSON.parse(event.data);

      farm?.barns?.every((barn) => {
        const chicken = barn.chickens?.find((chicken) => {
          return chicken.id == data.chickenID;
        });
        if (typeof chicken === "undefined") {
          return true;
        }

        chicken.restingUntil = data.restingUntil;

        switch (data.eggType) {
          case 0:
            chicken.normalEggsLaid++;
            return false;
          case 1:
            chicken.goldEggsLaid++;
            return false;
          default:
            return true;
        }
      });

      setFarm((farm) => ({ ...farm, day: data.day ?? farm.day }));
    };

    return () => eventSource.close();
  }, []);

  return farm ? (
    <div className="flex h-screen">
      <div className="flex flex-col items-center basis-2/12 p-4">
        <h1 className="text-center text-3xl font-extrabold mb-3">
          {farm?.name}
        </h1>
        <span>
          <label>Current Day:</label> <span key={farm.day}>{farm?.day}</span>
        </span>

        <span>
          <label>Golden Eggs:</label>{" "}
          <span key={farm.goldenEggs}>{farm?.goldenEggs}</span>
        </span>

        <button className="btn-primary mt-4" onClick={() => buyBarn.mutate()}>
          Buy Barn
        </button>

        {error && <span className="mt-4 text-red-700">{error}</span>}
      </div>
      <div className="barns">
        {farm?.barns?.map((barn, index: number) => {
          return <Barn key={index} barn={barn} day={farm.day ?? 0} />;
        })}
      </div>
    </div>
  ) : (
    <div>Loading your farm</div>
  );
};
