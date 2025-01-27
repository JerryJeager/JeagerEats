import useCartStore from "@/store/useCartStore";
import { RestaurantMenuCardType } from "@/types";
import Image from "next/image";
import { useState, useEffect } from "react";

const RestaurantMenuCard = (Props: RestaurantMenuCardType) => {
  const { addToMenu, removeFromMenu, menu } = useCartStore((state) => state);
  const [isAdded, setIsAdded] = useState(false);

  useEffect(() => {
    const isItemInCart = menu.some((item) => item.id === Props.id);
    setIsAdded(isItemInCart);
  }, [menu, Props.id]);

  return (
    <div className="border rounded-md hover:shadow-md flex justify-between gap-3 flex-col-reverse p-2 md:h-[150px] md:flex-row">
      <div className="h-full flex flex-col">
        <div className="md:p-3">
          <h3 className="font-semibold">{Props.name}</h3>
          <p className="text-sm text-black text-opacity-40">
            {Props.description}
          </p>
        </div>
        <div className="flex flex-col md:flex-row gap-2 md:justify-between md:items-center">
          <div className="md:mt-auto md:p-3">
            <p className="text-primary">
              {"₦"}
              {Props.price}
            </p>
          </div>
          {Props.stock > 0 ? (
            <>
              {isAdded ? (
                <button
                  onClick={() => {
                    removeFromMenu(Props.id);
                    setIsAdded(false);
                  }}
                  className="rounded-md text-center p-1 bg-white text-primary border-primary border"
                >
                  Remove
                </button>
              ) : (
                <button
                  onClick={() => {
                    const menuItem = { ...Props, quantity: 1 };
                    addToMenu(menuItem);
                    setIsAdded(true);
                  }}
                  className="rounded-md text-center p-1 bg-primary text-white"
                >
                  Add to cart
                </button>
              )}
            </>
          ) : (
            <button
              disabled={true}
              className="rounded-md text-center p-1 bg-white text-primary border-primary border"
            >
              Out of stock
            </button>
          )}
        </div>
      </div>
      <div className="w-full md:w-1/2 h-[120px] md:h-full">
        <Image
          src={Props.img_url}
          width={220}
          height={100}
          alt="restaurant menu item"
          className="object-cover w-full h-full rounded-md"
        />
      </div>
    </div>
  );
};

export default RestaurantMenuCard;
