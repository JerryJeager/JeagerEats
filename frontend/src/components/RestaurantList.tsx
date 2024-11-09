import { Restaurant } from "@/data";
import bag from "../../public/assets/bag.png";
import Image from "next/image";

const RestaurantList = ({ link, name }: Restaurant) => {
  return (
    <div className="flex justify-between items-center border border-black border-opacity-40 py-2 px-4  rounded-md">
      <p>{name}</p>
      <button className="rounded-full bg-primary bg-opacity-20">
        <Image src={bag} className="w-[30px] lg:w-[40px] h-[30px] lg:h-[40px]" alt="bag" />
      </button>
    </div>
  );
};

export default RestaurantList;
