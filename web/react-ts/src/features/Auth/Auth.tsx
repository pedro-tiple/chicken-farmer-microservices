import * as React from "react";
import { Login } from "./Login";
import { Register } from "./Register";

export const Auth = () => {
  return (
    <div className={"flex h-screen flex-row items-center justify-center"}>
      <Login />
      <div className={"mx-8"} />
      <Register />
    </div>
  );
};
