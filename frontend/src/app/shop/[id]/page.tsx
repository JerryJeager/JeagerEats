"use client";
import Link from "next/link";
import React, { useEffect, useState } from "react";
import { IoArrowBackOutline } from "react-icons/io5";
import foodt1 from "../../../../public/assets/foodt1.jpg";
import Image from "next/image";
import { Separator } from "@/components/ui/separator";
import { BASE_URL } from "@/data";
import RestaurantMenuCard from "@/components/shop/RestaurantMenuCard";
import axios from "axios";
import { Restaurant, RestaurantMenuCardType } from "@/types";
import { usePathname } from "next/navigation";

const RestaurantPage = () => {
  const [restaurantDetails, setRestauntDetails] = useState<Restaurant>();
  const [menus, setMenus] = useState<RestaurantMenuCardType[]>();
  const pathname = usePathname();
  let id = pathname.split("/").pop();
  const formatTime = (time: string): string => {
    const [hour, minute] = time.split(":");
    const intHour = parseInt(hour, 10);
    const isPM = intHour >= 12;
    const formattedHour = intHour > 12 ? intHour - 12 : intHour || 12;
    return `${formattedHour}:${minute}${isPM ? "pm" : "am"}`;
  };
  useEffect(() => {
    const getRestaurantDetails = async () => {
      try {
        const res = await axios.get(`${BASE_URL()}/restaurants/${id}`);
        setRestauntDetails(res.data as Restaurant);
      } catch (error) {}
    };
    const getMenus = async () => {
      try {
        const res = await axios.get(`${BASE_URL()}/menus/restaurants/${id}`);
        setMenus(res.data as RestaurantMenuCardType[]);
      } catch (error) {}
    };

    getRestaurantDetails();
    getMenus();
  }, []);
  return (
    <main className="padx pt-8">
      <Link href="/shop" className="flex gap-2 items-center">
        <IoArrowBackOutline /> <span>Restaurants</span>
      </Link>

      <section className="w-full md:w-[65vw] mt-6">
        <div className="h-[220px]">
          <Image
            src={restaurantDetails?.profile_img ?? foodt1}
            width={100}
            height={100}
            alt="restaurant profile image"
            className="object-cover h-full w-full rounded-md"
          />
        </div>
        <h2 className="font-bold text-2xl mt-3">
          {restaurantDetails && restaurantDetails.name}
        </h2>
        <p>{restaurantDetails && restaurantDetails.description}</p>
        <div className="mt-2">
          <p>Opening Time</p>
          <p>
            {restaurantDetails && formatTime(restaurantDetails?.opening_time)}-
            {restaurantDetails && formatTime(restaurantDetails?.closing_time)}
          </p>
        </div>
        <p className="mt-4 text-sm text-primary">Meals</p>
        <Separator />
        <div className="grid grid-cols-2 gap-3 mt-6">
          {menus &&
            menus.length > 0 &&
            menus.map((item, index) => (
              <RestaurantMenuCard key={index} {...item} />
            ))}
        </div>
      </section>
    </main>
  );
};

export default RestaurantPage;
