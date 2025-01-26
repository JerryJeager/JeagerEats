import { RestaurantMenuCardType } from "@/types";
import Image from "next/image";

const RestaurantMenuCard = ({
  id,
  image,
  name,
  description,
  price,
  stock,
}: RestaurantMenuCardType) => {
  return (
    <div className="border rounded-md hover:shadow-md flex justify-between gap-3 flex-col-reverse md:flex-row  h-[120px]">
      <div className="h-full flex flex-col">
        <div className="p-3">
          <h3 className="font-semibold">{name}</h3>
          <p className="text-sm text-slate-200">{description}</p>
        </div>
        <div className="mt-auto p-3">
            <p className="text-primary">{"₦"}{price}</p>
        </div>
      </div>
      <div className="w-1/2 h-full">
        <Image src={image} alt="restaurant menu item" className="object-cover w-full h-full rounded-md"/>
      </div>
    </div>
  );
};

export default RestaurantMenuCard;
