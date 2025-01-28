import { MenuItemCardType } from "@/types";
import Image from "next/image";
import Link from "next/link";
import { MdEdit } from "react-icons/md";

const MenuItemCard = ({
  id,
  restaurant_id,
  name,
  description,
  price,
  is_available,
  img_url,
  stock,
  category,
}: MenuItemCardType) => {
  return (
    <div className="rounded-lg w-full md:text-lg ">
      <div className="">
        <div className="mt-2">
          <div className="flex gap-3 border border-slate-300 p-2 border-opacity-90 rounded-lg w-fit bg-white">
            <Image width={120} height={120} src={img_url} alt="item-image" />
            <div className="flex flex-col w-[170px]">
              <p className="font-medium">{name}</p>
              <p className="line-clamp-1 text-black text-opacity-70">
                {description}
              </p>
              <div className="flex gap-2 mt-auto">
                <p className="mb-0">
                  {"$"}
                  {price}
                </p>
                <Link
                  href={`/dashboard/vendor/menu/${id}`}
                  className="flex items-center text-primary"
                >
                  <MdEdit />
                  <p>Edit</p>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MenuItemCard;
