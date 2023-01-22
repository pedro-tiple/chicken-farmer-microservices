import { MutableRefObject } from "react";
import { V1Farm, V1Barn, V1Chicken } from "chicken-farmer-service/api";

export const SetupFarmSSE = (
  setFarm: (farm: V1Farm) => void,
  farmRef: MutableRefObject<V1Farm>
): EventSource => {
  // TODO endpoint from config
  // TODO all these maps iterate over values we know we don't need, improve performance on that.
  const eventSource = new EventSource("http://localhost:8083/event-feed", {
    withCredentials: false
  });

  eventSource.addEventListener("day", (event) => {
    setFarm({ ...farmRef.current, day: event.data });
  });

  eventSource.addEventListener("golden-eggs-change", (event) => {
    const data = JSON.parse(event.data);
    setFarm({
      ...farmRef.current,
      goldenEggs: farmRef.current.goldenEggs + data.count
    });
  });

  eventSource.addEventListener("feed-change", (event) => {
    const data = JSON.parse(event.data);
    setFarm({
      ...farmRef.current,
      barns: farmRef.current.barns.map<V1Barn>((barn: V1Barn): V1Barn => {
        if (barn.id == data.barnID) {
          barn.feed += data.count;
        }
        return barn;
      })
    });
  });

  eventSource.addEventListener("new-barn", (event) => {
    const data = JSON.parse(event.data);
    setFarm({
      ...farmRef.current,
      barns: [
        ...farmRef.current.barns,
        {
          id: data.barnID,
          feed: data.feed,
          hasAutoFeeder: data.hasAutoFeeder,
          chickens: new Array<V1Chicken>()
        }
      ]
    });
  });

  eventSource.addEventListener("new-chicken", (event) => {
    const data = JSON.parse(event.data);
    setFarm({
      ...farmRef.current,
      barns: farmRef.current.barns.map<V1Barn>((barn: V1Barn): V1Barn => {
        if (barn.id == data.barnID) {
          barn.chickens.push({
            id: data.chickenID,
            dateOfBirth: data.dateOfBirth,
            restingUntil: 0,
            normalEggsLaid: 0,
            goldEggsLaid: 0
          });
        }
        return barn;
      })
    });
  });

  eventSource.addEventListener("chicken-fed", (event) => {
    const data = JSON.parse(event.data);
    setFarm({
      ...farmRef.current,
      barns: farmRef.current.barns.map<V1Barn>((barn: V1Barn): V1Barn => {
        const chicken = barn.chickens.find((chicken: V1Chicken) => {
          return chicken.id == data.chickenID;
        });

        if (typeof chicken != "undefined") {
          chicken.restingUntil = data.restingUntil;
          chicken.normalEggsLaid += data.normalEggsLaid;
          chicken.goldEggsLaid += data.goldEggsLaid;
        }
        return barn;
      })
    });
  });

  return eventSource;
};
