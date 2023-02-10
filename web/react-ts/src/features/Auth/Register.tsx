import React, { useContext, useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { V1RegisterRequest } from "chicken-farmer-service";
import { ServicesContext } from "../../context/ServicesContext";

export const Register = () => {
  const [farmerName, setFarmerName] = useState<string>("");
  const [farmName, setFarmName] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const { farmerServiceApi } = useContext(ServicesContext);

  const register = useMutation({
    mutationKey: ["register", farmerName],
    mutationFn: ({ farmerName, farmName, password }: V1RegisterRequest) => {
      return farmerServiceApi.farmerPublicServiceRegister({
        farmName,
        farmerName,
        password
      });
    }
  });

  return (
    <>
      {!register.isSuccess && (
        <div className="w-full max-w-xs">
          <form className="mb-4 rounded bg-white px-8 pt-6 pb-8 shadow-md">
            <div className="mb-4">
              <label
                className="mb-2 block text-sm font-bold text-gray-700"
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
            <div className="mb-4">
              <label
                className="mb-2 block text-sm font-bold text-gray-700"
                htmlFor="farmName">
                Farmer
              </label>
              <input
                className="focus:shadow-outline w-full appearance-none rounded border py-2 px-3 leading-tight text-gray-700 shadow focus:outline-none"
                id="farmName"
                type="text"
                placeholder="Farm's Name"
                onChange={(e) => setFarmName(e.target.value)}
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
                    register.error && "border-red-500"
                  } focus:shadow-outline mb-3 w-full appearance-none rounded border py-2 px-3 leading-tight text-gray-700 shadow focus:outline-none`}
                  id="password"
                  type="password"
                  placeholder="******************"
                  onChange={(e) => setPassword(e.target.value)}
                />
                {register.error && (
                  <p className="text-xs italic text-red-500">Invalid login.</p>
                )}
              </>
            </div>
            <div className="flex items-center justify-between">
              <button
                className="focus:shadow-outline rounded bg-blue-500 py-2 px-4 font-bold text-white hover:bg-blue-700 focus:outline-none disabled:bg-blue-100"
                type="button"
                disabled={!farmerName || !farmName || !password}
                onClick={() =>
                  register.mutate({ farmerName, farmName, password })
                }>
                Register
              </button>
            </div>
          </form>
        </div>
      )}
    </>
  );
};
