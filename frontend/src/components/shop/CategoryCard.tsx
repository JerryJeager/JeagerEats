import { CategoryCardType } from "@/types";
import Image from "next/image";

const CategoryCard = ({ color, image, name }: CategoryCardType) => {
  return (
    <div
      style={{ backgroundColor: color }}
      className="flex flex-col gap-3 items-center justify-center rounded-lg w-[120px]  p-4"
    >
      <div>
        <Image
          src={image}
          width={30}
          height={30}
          className="mx-auto"
          alt={"category-image"}
        />
      </div>
      <p className="text-sm">{name}</p>
    </div>
  );
};

export default CategoryCard;
