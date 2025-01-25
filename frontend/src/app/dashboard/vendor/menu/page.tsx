"use client";
import Image from "next/image";
import Link from "next/link";
import item from "../../../../../public/assets/jollof.png";
import { MdEdit } from "react-icons/md";
import empty from "../../../../../public/assets/empty.svg";
import { MenuItemCardType } from "@/types";
import { useEffect, useState } from "react";
import { getCookie } from "@/actions/handleCookies";
import axios from "axios";
import { BASE_URL } from "@/data";
import MenuItemCard from "@/components/dashboard/MenuItemCard";

const Menu = () => {
  const [menuItems, setMenuItems] = useState<MenuItemCardType[]>();
  useEffect(() => {
    const getMenuData = async () => {
      try {
        const accessToken = await getCookie("jeagereats_token");
        const res = await axios.get(`${BASE_URL()}/menus`, {
          headers: {
            Authorization: `Bearer ${accessToken?.value}`,
          },
        });
        const data = res.data;
        setMenuItems(data);
      } catch (error) {
        console.error("Error fetching restaurant data:", error);
      }
    };

    getMenuData();
  }, []);
  return (
    <section className="mt-8 min-h-screen">
      <h2 className="font-bold text-2xl md:text-3xl text-center">
        Menu Management
      </h2>
      <div className="rounded-lg w-full mt-8 p-4 md:p-8 shadow-lg bg-white  h-full md:text-lg ">
        <div className="flex flex-col items-center">
          <Image width={300} src={empty} alt="empty illustration" />
          <p className="text-center">Add more food items to your store</p>
          <Link
            href={"/dashboard/vendor/menu/add"}
            className="p-2 md:p-3 w-[220px] mt-6 mx-auto text-center bg-primary text-white rounded-lg"
          >
            Add Menu Item
          </Link>
        </div>
      </div>
      <div className="rounded-lg w-full mt-8 p-4 md:p-8 shadow-lg h-full md:text-lg ">
        <div className="">
          <p className="text-center font-semibold">Menu Items</p>
          <div className="flex flex-wrap gap-4">
            {" "}
            {menuItems &&
              menuItems.map((item) => <MenuItemCard key={item.id} {...item} />)}
          </div>
        </div>
      </div>
    </section>
  );
};

export default Menu;
