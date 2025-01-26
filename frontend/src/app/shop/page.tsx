"use client";
import CategoryCard from "@/components/shop/CategoryCard";
import { BASE_URL, categoriesCards } from "@/data";
import foodt1 from "../../../public/assets/foodt1.jpg";
import foodt2 from "../../../public/assets/foodt2.jpg";
import Image from "next/image";
import RestaurantCard from "@/components/shop/RestaurantCard";
import { useEffect, useState } from "react";
import axios from "axios";
import { RestaurantCardType } from "@/types";

const Shop = () => {
  const [allRestaurants, setAllRestaurants] = useState<RestaurantCardType[]>();
  const [isLoading, setIsLoading] = useState(false);
  useEffect(() => {
    const fetchRestaurants = async () => {
      setIsLoading(true);
      try {
        const res = await axios.get(`${BASE_URL()}/restaurants`);
        setAllRestaurants(res.data as RestaurantCardType[]);
        console.log(res.data)
      } catch (error) {
        console.log(error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchRestaurants();
  }, []);
  return (
    <>
      <section className="px-[5%] lg:px-[8%] py-8">
        <h2 className="font-bold text-2xl">Place your orders now</h2>
        <div className="flex flex-wrap gap-4 mt-8">
          {categoriesCards.map((c, index) => (
            <CategoryCard key={index} {...c} />
          ))}
        </div>

        <div className="w-full flex gap-4 justify-between mt-5">
          <div className="w-full h-[10rem]">
            <Image
              src={foodt1}
              className="object-cover h-full w-full rounded-2xl"
              alt="food t image"
            />
          </div>
          <div className="w-full h-[10rem]">
            <Image
              alt="food image"
              className="object-cover h-full w-full rounded-2xl"
              src={foodt2}
            />
          </div>
        </div>
      </section>

      <section className="padx">
        <h2 className="font-bold text-2xl">All Restaurants</h2>
        <div
          className="grid gap-4 p-4"
          style={{
            gridTemplateColumns: "repeat(auto-fit, minmax(250px, 1fr))",
          }}
        >
          {allRestaurants &&
            allRestaurants.length > 0 &&
            allRestaurants.map((r, index) => (
              <RestaurantCard key={index} {...r} />
            ))}
        </div>
      </section>
    </>
  );
};

export default Shop;
