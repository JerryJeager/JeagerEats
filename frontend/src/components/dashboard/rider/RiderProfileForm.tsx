"use client";

import { getCookie } from "@/actions/handleCookies";
import { BASE_URL } from "@/data";
import useUserStore from "@/store/useUserStore";
import { User } from "@/types";
import axios from "axios";
import { ChangeEvent, FormEvent, useEffect, useState } from "react";

interface FormDataBasic {
  first_name: string;
  last_name: string;
  address: string;
  phone_number: string;
}

const RiderProfileForm = () => {
  const { user, setUser } = useUserStore((state) => state);
  const [isLoadingBasic, setIsLoadingBasic] = useState(false);
  const [formDataBasic, setFormDataBasic] = useState<FormDataBasic>({
    address: "",
    first_name: "",
    last_name: "",
    phone_number: "",
  });
  const handleChangeBasic = (
    e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    let { name, value } = e.target;
    setFormDataBasic((prev) => {
      return {
        ...prev,
        [name]: value,
      };
    });
  };

  useEffect(() => {
    const getUserData = async () => {
      try {
        const accessToken = await getCookie("jeagereats_token");
        const res = await axios.get(`${BASE_URL()}/users`, {
          headers: {
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
        const data: User = res.data;
        console.log(data);
        setUser(data);
        setFormDataBasic({
          address: data.address,
          first_name: data.first_name,
          last_name: data.last_name,
          phone_number: data.phone_number,
        });
      } catch (error) {}
    };

    if (!user) {
      getUserData();
    } else {
      setFormDataBasic({
        address: user.address,
        first_name: user.first_name,
        last_name: user.last_name,
        phone_number: user.phone_number,
      });
    }
  }, []);
  const handleSubmitBasic = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setIsLoadingBasic(true)
    try{
        

    }catch(error) {

    }finally{
        setIsLoadingBasic(false)
    }
  };
  return (
    <div>
      <form action="" onSubmit={handleSubmitBasic} className="w-full md:w-[60%]">
        <h2>Basic Information</h2>
        <div className="flex flex-col gap-6 w-full mt-3">
          <label>
            <p>First Name</p>
            <input
              name="first_name"
              type="text"
              placeholder="Chitis"
              className="outline w-full outline-none border border-primary p-2 mt-2 caret-primary rounded-lg"
              value={formDataBasic.first_name}
              onChange={handleChangeBasic}
            />
          </label>
          <label>
            <p>Last Name</p>
            <input
              name="last_name"
              type="text"
              className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
              value={formDataBasic.last_name}
              onChange={handleChangeBasic}
            />
          </label>
        </div>

        <div>
          <label>
            <p>Address</p>
            <input
              name="address"
              type="text"
              className="outline-none w-full border border-primary p-2 mt-2 caret-primary rounded-lg"
              value={formDataBasic.address}
              onChange={handleChangeBasic}
            />
          </label>
        </div>

        <div className="flex gap-8">
          <label className="w-full">
            <p>Phone Number</p>
            <input
              name="phone_number"
              type="text"
              className="outline-none border border-primary p-2 mt-2 caret-primary w-full rounded-lg"
              value={formDataBasic.phone_number}
              onChange={handleChangeBasic}
            />
          </label>
        </div>
        <button
          type="submit"
          className="bg-primary px-2 mt-4 text-white rounded-lg py-2"
          disabled={isLoadingBasic}
        >
          {isLoadingBasic ? "Saving Changes..." : "Save Changes"}
        </button>
      </form>
    </div>
  );
};

export default RiderProfileForm;
