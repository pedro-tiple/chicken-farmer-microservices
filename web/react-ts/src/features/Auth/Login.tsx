import React, { FormEvent, useContext, useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { V1LoginRequest } from "chicken-farmer-service";
import { ServicesContext } from "../../context/ServicesContext";

export const Login = () => {
  const [farmerName, setFarmerName] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const { setAuthToken, farmerServiceApi } = useContext(ServicesContext);

  const login = useMutation({
    mutationKey: ["login", farmerName],
    mutationFn: ({ farmerName, password }: V1LoginRequest) => {
      return farmerServiceApi.farmerPublicServiceLogin({
        farmerName,
        password
      });
    },
    onSuccess: (data) => setAuthToken(data.data.authToken ?? "")
  });

  return (
    <div className="w-full max-w-xs flex-initial">
      <form
        className="mb-4 rounded bg-white px-8 pt-6 pb-8 shadow-md"
        onSubmit={(e: FormEvent<HTMLFormElement>) => {
          e.preventDefault();
          login.mutate({ farmerName, password });
        }}>
        <div className="mb-4">
          <label
            className="mb-2 mb-2 block text-sm font-bold text-gray-700"
            htmlFor="farmerName">
            Farmer
          </label>
          <input
            className="focus:shadow-outline w-full appearance-none rounded border py-2 px-3 leading-tight text-gray-700 shadow focus:outline-none"
            id="farmerName"
            type="text"
            placeholder="Farmer's Name"
            onChange={(e) => setFarmerName(e.target.value)}
          />
        </div>
        <div className="mb-6">
          <>
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="password">
              Password
            </label>
            <input
              className={`${
                login.error && "border-red-500"
              } focus:shadow-outline mb-3 w-full appearance-none rounded border py-2 px-3 leading-tight text-gray-700 shadow focus:outline-none`}
              id="password"
              type="password"
              placeholder="******************"
              onChange={(e) => setPassword(e.target.value)}
            />
            {login.error && (
              <p className="text-xs italic text-red-500">Invalid login.</p>
            )}
          </>
        </div>
        <div className="flex items-center justify-between">
          <button
            className="focus:shadow-outline rounded bg-blue-500 py-2 px-4 font-bold text-white hover:bg-blue-700 focus:outline-none disabled:bg-blue-100"
            type="submit"
            disabled={!farmerName || !password}>
            Sign In
          </button>
        </div>
      </form>
    </div>
  );
};
