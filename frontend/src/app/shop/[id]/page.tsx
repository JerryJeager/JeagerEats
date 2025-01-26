import Link from "next/link";
import React from "react";
import { IoArrowBackOutline } from "react-icons/io5";
import foodt1 from "../../../../public/assets/foodt1.jpg";
import Image from "next/image";
import { Separator } from "@/components/ui/separator";
import { menuItemsDummyData } from "@/data";
import RestaurantMenuCard from "@/components/shop/RestaurantMenuCard";

const RestaurantPage = () => {
  return (
    <main className="padx pt-8">
      <Link href="/shop" className="flex gap-2 items-center">
        <IoArrowBackOutline /> <span>Restaurants</span>
      </Link>

      <section className="w-full md:w-[65vw] mt-6">
        <div className="h-[220px]">
          <Image
            src={foodt1}
            alt="restaurant profile image"
            className="object-cover h-full w-full rounded-md"
          />
        </div>
        <h2 className="font-bold text-2xl mt-3">Zaddy's Place</h2>
        <p className="mt-4 text-sm text-primary">Meals</p>
        <Separator />
        <div className="grid grid-cols-2 gap-3 mt-6">
          {menuItemsDummyData.map((item, index) => (
            <RestaurantMenuCard key={index} {...item} />
          ))}
        </div>
      </section>
    </main>
  );
};

export default RestaurantPage;
