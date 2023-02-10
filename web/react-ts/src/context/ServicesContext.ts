import { createContext } from "react";
import { FarmServiceApi } from "chicken-farmer-service/api";
import { FarmerPublicServiceApi } from "chicken-farmer-service";

export type Services = {
  farmServiceApi: FarmServiceApi;
  farmerServiceApi: FarmerPublicServiceApi;
  authToken: string;
  setAuthToken: (newAuthToken: string) => void;
};

export const ServicesContext = createContext<Services>({
  farmServiceApi: new FarmServiceApi(),
  farmerServiceApi: new FarmerPublicServiceApi(),
  authToken: "",
  setAuthToken: () => {
    return;
  }
});
