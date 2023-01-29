import React, { useContext, useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { Configuration } from "chicken-farmer-service/configuration";
import { FarmerServiceApi } from "chicken-farmer-service";
import { UserAuthContext } from "../../context/UserContext";

const farmerServiceApi = new FarmerServiceApi(
  new Configuration({ basePath: "http://localhost:8082" })
);

export const Login = () => {
  const [farmerName, setFarmerName] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const { setJWT } = useContext(UserAuthContext);

  const login = useMutation({
    mutationKey: ["login", farmerName],
    mutationFn: ({
      farmerName,
      password
    }: {
      farmerName: string;
      password: string;
    }) => {
      return farmerServiceApi.farmerServiceLogin(farmerName, password);
    },
    onSuccess: (data) => setJWT(data.data.jwt ?? "")
  });

  return (
    <div className="flex-initial w-full max-w-xs">
      <form className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <div className="mb-4">
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="farmerName">
            Farmer
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="farmerName"
            type="text"
            placeholder="Farmer's Name"
            onChange={(e) => setFarmerName(e.target.value)}
          />
        </div>
        <div className="mb-6">
          <>
            <label
              className="block text-gray-700 text-sm font-bold mb-2"
              htmlFor="password">
              Password
            </label>
            <input
              className={`${
                login.error && "border-red-500"
              } shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline`}
              id="password"
              type="password"
              placeholder="******************"
              onChange={(e) => setPassword(e.target.value)}
            />
            {login.error && (
              <p className="text-red-500 text-xs italic">Invalid login.</p>
            )}
          </>
        </div>
        <div className="flex items-center justify-between">
          <button
            className="bg-blue-500 hover:bg-blue-700 disabled:bg-blue-100 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="button"
            disabled={!farmerName || !password}
            onClick={() => login.mutate({ farmerName, password })}>
            Sign In
          </button>
        </div>
      </form>
    </div>
  );
};
