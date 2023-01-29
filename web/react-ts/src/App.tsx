import "./App.scss";
import React, { useState } from "react";
import { Farm } from "./features/Farm/Farm";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { UserAuthContext, UserAuth } from "./context/UserContext";
import { Auth } from "./features/Auth/Auth";

const queryClient = new QueryClient();

function App() {
  const [userAuth, setUserAuth] = useState<UserAuth>({
    farmName: "",
    jwt: "",
    name: "",
    setJWT(jwt: string) {
      return setUserAuth((prevValue) => {
        return { ...prevValue, jwt };
      });
    }
  });

  return (
    <div className="App">
      <UserAuthContext.Provider value={userAuth}>
        <QueryClientProvider client={queryClient}>
          {userAuth.jwt ? <Farm /> : <Auth />}
        </QueryClientProvider>
      </UserAuthContext.Provider>
    </div>
  );
}

export default App;
