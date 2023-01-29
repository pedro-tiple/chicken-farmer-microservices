import { createContext } from "react";

export type UserAuth = {
  name: string;
  farmName: string;
  jwt: string;
  setJWT(jwt: string): void;
};

export const UserAuthContext = createContext<UserAuth>({
  name: "",
  farmName: "",
  jwt: "",
  setJWT: (jwt: string) => {
    return null;
  }
});
