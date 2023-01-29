import * as React from "react";
import { Login } from "./Login";
import { Register } from "./Register";

export const Auth = () => {
  return (
    <div className={"h-screen flex flex-row justify-center items-center"}>
      <Login />
      <div className={"mx-8"} />
      <Register />
    </div>
  );
};
