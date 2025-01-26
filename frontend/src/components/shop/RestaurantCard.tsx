import Image from "next/image";
import { RestaurantCardType } from "@/types";
import Link from "next/link";

const RestaurantCard = ({id, profile_img, name}: RestaurantCardType) => {
  return (
    <Link href={`/shop/${id}`} className="p-4 hover:shadow-lg w-fit rounded-lg border">
      <div className="w-full h-[10rem]">
        <Image
          src={profile_img}
          width={230}
          height={100}
          alt="restaurant image"
          className="rounded-md h-full object-cover"
        />
      </div>
      <div>
        <p className="font-semibold text-xl">{name}</p>
        <p className="text-primary text-sm mt-4">Meals</p>
      </div>
    </Link>
  );
};

export default RestaurantCard;
