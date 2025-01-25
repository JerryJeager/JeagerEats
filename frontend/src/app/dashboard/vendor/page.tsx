"use client";
import Image from "next/image";
import profile from "../../../../public/assets/chef.png";
import Link from "next/link";
import { MdEdit } from "react-icons/md";
import item from "../../../../public/assets/jollof.png";
import { useEffect, useState } from "react";
import axios from "axios";
import { BASE_URL } from "@/data";
import { getCookie } from "@/actions/handleCookies";
import MenuItemCard from "@/components/dashboard/MenuItemCard";
import { MenuItemCardType } from "@/types";

const VendorDashboard = () => {
  const [restaurantData, setRestaurantData] = useState({
    name: "Restaurant Name",
    profile_img: null,
    opening_time: "8:00am",
    closing_time: "8:00pm",
  });

  const formatTime = (time: string): string => {
    const [hour, minute] = time.split(":");
    const intHour = parseInt(hour, 10);
    const isPM = intHour >= 12;
    const formattedHour = intHour > 12 ? intHour - 12 : intHour || 12;
    return `${formattedHour}:${minute}${isPM ? "pm" : "am"}`;
  };

  const [menuItems, setMenuItems] = useState<MenuItemCardType[]>();
  useEffect(() => {
    const getRestaurantData = async () => {
      try {
        const accessToken = await getCookie("jeagereats_token");
        const res = await axios.get(`${BASE_URL()}/restaurants/self`, {
          headers: {
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
        const data = res.data;

        setRestaurantData({
          name: data.name || "Restaurant Name",
          profile_img: data.profile_img || null,
          opening_time: data.opening_time
            ? formatTime(data.opening_time.slice(11, 16)) // Extract and format
            : "8:00am",
          closing_time: data.closing_time
            ? formatTime(data.closing_time.slice(11, 16)) // Extract and format
            : "8:00pm",
        });
      } catch (error) {
        console.error("Error fetching restaurant data:", error);
      }
    };
    const getMenuData = async () => {
      try {
        const accessToken = await getCookie("jeagereats_token");
        const res = await axios.get(`${BASE_URL()}/menus`, {
          headers: {
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
        const data = res.data;
        setMenuItems(data);
      } catch (error) {
        console.error("Error fetching restaurant data:", error);
      }
    };

    getRestaurantData();
    getMenuData()
  }, []);

  return (
    <section className="mt-8">
      <h2 className="font-bold text-2xl md:text-3xl text-center">Welcome</h2>
      <div className="rounded-lg w-full mt-8 p-4 md:p-8 shadow-lg bg-white h-full md:text-lg">
        <div className="flex flex-wrap gap-8">
          <Image
            src={restaurantData.profile_img || profile}
            alt="profile image"
            width={200}
            height={200}
            className="rounded-full h-[200px] w-[200px]"
          />
          <div className="flex flex-col gap-3">
            <h3 className="text-2xl font-semibold">{restaurantData.name}</h3>
            <div className="flex gap-4 mt-4">
              <div className="flex flex-col gap-3">
                <p>Opening Time</p>
                <p>{restaurantData.opening_time}</p>
              </div>
              <div className="flex flex-col gap-3">
                <p>Closing Time</p>
                <p>{restaurantData.closing_time}</p>
              </div>
            </div>
            <div className="flex gap-8 mt-auto">
              <div>
                <p>0</p>
                <p>Items Added</p>
              </div>
              <div>
                <p>0</p>
                <p>Disabled Items</p>
              </div>
              <Link
                href={"/dashboard/vendor/profile"}
                className="mt-auto p-3 rounded-lg bg-primary text-white w-[130px] text-center"
              >
                Edit Profile
              </Link>
            </div>
          </div>
        </div>
      </div>
      <div className="rounded-lg w-full mt-8 p-4 md:p-8 shadow-lg h-full md:text-lg ">
        <div className="">
          <p className="text-center font-semibold">Menu Items</p>
          <div className="flex flex-wrap gap-4">
            {" "}
            {menuItems &&
              menuItems.map((item) => <MenuItemCard key={item.id} {...item} />)}
          </div>
        </div>
      </div>
    </section>
  );
};

export default VendorDashboard;
