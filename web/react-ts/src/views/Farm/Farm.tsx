import "./Farm.scss";
import { useQuery } from "@tanstack/react-query";
import * as React from "react";
import { FarmServiceApi, V1Farm } from "chicken-farmer-service/api";
import { Configuration } from "chicken-farmer-service/configuration";
import { Barn } from "../../components/Barn/Barn";
import { useEffect, useState } from "react";
import { Buffer } from "buffer";

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

  useEffect(() => {
    setFarm(result);
    const es = new EventSource("http://localhost:8083/event-feed", {
      withCredentials: false
    });
    es.addEventListener(
      "data",
      (event) => {
        // TODO deal with this stupid quotes thing server side later
        if (!farm || event.data.trim() == '"accepted"') {
          return;
        }

        const data = JSON.parse(Buffer.from(event.data, "base64").toString());

        setFarm((farm) => ({ ...farm, day: data.day ?? farm.day }));

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
      },
      false
    );

    es.addEventListener("close", () => {
      console.log("close");
      es.close();
    });

    return () => es.close();
  }, [result]);
  console.log("farm update");

  return farm ? (
    <div className="farm" key={farm.day}>
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
  ) : (
    <div>Loading your farm</div>
  );
};
