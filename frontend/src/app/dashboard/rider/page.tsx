"use client";

import { getCookie } from "@/actions/handleCookies";
import { BASE_URL } from "@/data";
import useUserStore from "@/store/useUserStore";
import { User } from "@/types";
import axios from "axios";
import { useEffect, useState } from "react";

const RiderDashboard = () => {
  const [userData, setUserData] = useState<User>();
  const [riderData, setRiderData] = useState()
  const { user, setUser } = useUserStore((state) => state);
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
        setUser(data);
        setUserData(data);
      } catch (error) {}
    };

    if (!user) {
      getUserData();
    } else {
      setUserData(user);
    }
  }, []);
  return (
    <main className="mt-8">
      <h2 className="font-bold text-2xl md:text-3xl text-center">
        Welcome Back, {userData && userData?.first_name}
      </h2>
    </main>
  );
};

export default RiderDashboard;
