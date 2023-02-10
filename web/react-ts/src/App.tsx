import "./App.scss";
import React, { useEffect, useState } from "react";
import { Farm } from "./features/Farm/Farm";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Auth } from "./features/Auth/Auth";
import { ServicesContext } from "./context/ServicesContext";
import { FarmServiceApi } from "chicken-farmer-service/api";
import { FarmerPublicServiceApi } from "chicken-farmer-service";
import { Configuration } from "chicken-farmer-service/configuration";

const queryClient = new QueryClient();

function App() {
  const [authToken, setAuthToken] = useState<string>("");
  const [farmServiceApi, setFarmServiceApi] = useState<FarmServiceApi>(
    new FarmServiceApi()
  );
  const [farmerServiceApi, setFarmerServiceApi] =
    useState<FarmerPublicServiceApi>(new FarmerPublicServiceApi());

  useEffect(() => {
    console.log("setting services", authToken);
    setFarmServiceApi(
      new FarmServiceApi(
        new Configuration({
          basePath: "http://localhost:8081",
          apiKey: `Bearer ${authToken}`
        })
      )
    );
    setFarmerServiceApi(
      new FarmerPublicServiceApi(
        new Configuration({
          basePath: "http://localhost:8082",
          apiKey: `Bearer ${authToken}`
        })
      )
    );
  }, [authToken]);

  return (
    <div className="App">
      <ServicesContext.Provider
        value={{ authToken, setAuthToken, farmServiceApi, farmerServiceApi }}>
        <QueryClientProvider client={queryClient}>
          {authToken ? <Farm /> : <Auth />}
        </QueryClientProvider>
      </ServicesContext.Provider>
    </div>
  );
}

export default App;
