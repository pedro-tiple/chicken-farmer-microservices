import { MutableRefObject } from "react";
import { V1Farm, V1Barn, V1Chicken } from "chicken-farmer-service/api";

enum Events {
  Day = "day",
  GoldenEggsChange = "golden-eggs-change",
  FeedChange = "feed-change",
  NewBarn = "new-barn",
  NewChicken = "new-chicken",
  SoldChicken = "sold-chicken",
  ChickenFed = "chicken-fed"
}

export const SetupFarmSSE = (
  setFarm: (farm: V1Farm) => void,
  farmRef: MutableRefObject<V1Farm>
): EventSource => {
  // TODO endpoint from config
  // TODO all these maps iterate over values we know we don't need, improve performance on that.
  const eventSource = new EventSource("http://localhost:8083/event-feed", {
    withCredentials: false
  });

  eventSource.addEventListener(Events.Day, (event) => {
    setFarm({ ...farmRef.current, day: event.data });
  });

  eventSource.addEventListener(Events.GoldenEggsChange, (event) => {
    const data = JSON.parse(event.data);
    setFarm({
      ...farmRef.current,
      goldenEggs: farmRef.current.goldenEggs + data.count
    });
  });

  eventSource.addEventListener(Events.FeedChange, (event) => {
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

  eventSource.addEventListener(Events.NewBarn, (event) => {
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

  eventSource.addEventListener(Events.NewChicken, (event) => {
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

  eventSource.addEventListener(Events.SoldChicken, (event) => {
    const data = JSON.parse(event.data);
    setFarm({
      ...farmRef.current,
      barns: farmRef.current.barns.map<V1Barn>((barn: V1Barn): V1Barn => {
        const index = barn.chickens.findIndex(
          (chicken) => chicken.id == data.chickenID
        );
        if (index != -1) {
          barn.chickens.splice(index, 1);
        }
        return barn;
      })
    });
  });

  eventSource.addEventListener(Events.ChickenFed, (event) => {
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
