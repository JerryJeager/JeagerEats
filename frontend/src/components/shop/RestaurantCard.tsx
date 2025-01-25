import Image from "next/image";
import { RestaurantCardType } from "@/types";

const RestaurantCard = ({id, image, name}: RestaurantCardType) => {
  return (
    <div className="p-4 hover:shadow-lg rounded-lg border">
      <div className="w-full h-[10rem]">
        <Image
          src={image}
          alt="restaurant image"
          className="rounded-md w-full h-full object-cover"
        />
      </div>
      <div>
        <p className="font-semibold text-xl">{name}</p>
        <p className="text-primary text-sm mt-4">Meals</p>
      </div>
    </div>
  );
};

export default RestaurantCard;
