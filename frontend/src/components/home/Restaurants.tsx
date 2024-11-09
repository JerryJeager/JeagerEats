import { comforta } from "./Hero";
import map from "../../../public/assets/map.png";
import Image from "next/image";
import { restaurants } from "@/data";
import RestaurantList from "../RestaurantList";

const Restaurants = () => {
  return (
    <section className="bg-[#FFF1EA] mt-8 pt-8 lg:pt-10 pb-10  lg:pb-14 px-[5%] lg:px-[8%] ">
      <h2
        className={`text-2xl lg:text-4xl mb-8 lg:mb-10 text-primary ${comforta.className} text-center lg:max-w-[600px] mx-auto`}
      >
        Your Favourite Restaurants, All in One Place
      </h2>
      <div className="flex flex-col lg:flex-row">
        <div className="w-full">
          <Image
            src={map}
            alt="map"
            className="border-black border-2 lg:border-4 rounded-t-2xl lg:rounded-l-2xl"
          />
        </div>
        <div className="border-black border-2 lg:border-4 mt-2 lg:mt-0 rounded-b-2xl lg:rounded-r-2xl w-full px-8 pb-4 lg:pb-0 lg:px-14 pt-6 ">
          <p
            className={`text-primary text-xl font-semibold ${comforta.className}`}
          >
            Restaurants
          </p>
          <div className="flex flex-col gap-4 mt-4">
            {restaurants.map((r, index) => (
              <RestaurantList key={index} {...r} />
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};

export default Restaurants;
