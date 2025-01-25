import bike from "../../public/assets/Motorcycle.png";
import shop from "../../public/assets/shop.png";
import cutlery from "../../public/assets/cutlery.png";
import { StaticImageData } from "next/image";
import { CategoryCardType, RestaurantCardType, VendorDashboarNav } from "@/types";
import { IoHome } from "react-icons/io5";
import { FaUser } from "react-icons/fa";
import { MdOutlineRestaurantMenu } from "react-icons/md";

import africameals from "../../public/assets/africameals.webp";
import drinks from "../../public/assets/drinks.webp";
import fitfam from "../../public/assets/fitfam.webp";
import localMarket from "../../public/assets/local-market.svg"
import restaurantH from "../../public/assets/restaurantH.webp";
import supermarket from "../../public/assets/supermarket.webp";

import demofood from "../../public/assets/demofood.jpg"
import foodt1 from "../../public/assets/foodt1.jpg"
import foodt2 from "../../public/assets/foodt2.jpg"

export const BASE_URL = () => {
  const environment = process.env.NEXT_PUBLIC_ENVIRONMENT;
  const baseUrl =
    environment === "production"
      ? "https://jeagereats-production.up.railway.app/api/v1"
      : "http://localhost:8080/api/v1";
  return baseUrl;
};

export type Join = {
  title: string;
  content: string;
  cta: string;
  link: string;
  icon: StaticImageData;
  color: string;
};
export type Restaurant = {
  name: string;
  link: string;
};

export const JoinData: Join[] = [
  {
    title: "Earn with Flexibility",
    content:
      "Deliver with us and earn whenever you want. Flexibility meets opportunity",
    cta: "Sign up as a Rider",
    link: "auth/rider",
    icon: bike,
    color: "#FFE5CC",
  },
  {
    title: "Grow your Restaurant",
    content:
      "Reach more customers and boost your sales. Partner with JaegerEats",
    cta: "Become a Partner",
    link: "auth/vendor",
    icon: shop,
    color: "#F0EDEE",
  },
  {
    title: "Order in Seconds",
    content:
      "Get your favorite meals delivered fast. Simple, easy, and reliable.",
    cta: "Order Now",
    link: "shop",
    icon: cutlery,
    color: "#CDE7FF",
  },
];

export const restaurants: Restaurant[] = [
  {
    name: "Zaddy's Place",
    link: "",
  },
  {
    name: "Bistro",
    link: "",
  },
  {
    name: "Chitis",
    link: "",
  },
  {
    name: "Achla",
    link: "",
  },
  {
    name: "Basmati",
    link: "",
  },
];

export const vendorDashboardNav: VendorDashboarNav[] = [
  { name: "Home", link: "/dashboard/vendor", icon: IoHome },
  {
    name: "Menu Management",
    link: "/dashboard/vendor/menu",
    icon: MdOutlineRestaurantMenu,
  },
  { name: "Profile", link: "/dashboard/vendor/profile", icon: FaUser },
];

export const categoriesCards: CategoryCardType[] = [
  {
    color: "#ffeaf1",
    image: restaurantH,
    name: "Restaurants",
  },
  {
    color: "#fff3ed",
    image: supermarket,
    name: "Supermarkets",
  },
  {
    color: "#dcf7ed",
    image: localMarket,
    name: "Local Markets",
  },
  {
    color: "#fff8e4",
    image: africameals,
    name: "African Meals",
  },
  {
    color: "#daefe3",
    image: fitfam,
    name: "Fit Fam",
  },
  {
    color: "#fff8e4",
    image: drinks,
    name: "Drinks",
  },
];

export const allRestaurants: RestaurantCardType[] = [
  {
    id: "1",
    name: "Zaddy's Place",
    image: demofood,
  },
  {
    id: "1",
    name: "Zaddy's Place",
    image: foodt1,
  },
  {
    id: "1",
    name: "Zaddy's Place",
    image: demofood,
  },
  {
    id: "1",
    name: "Zaddy's Place",
    image: demofood,
  },
  {
    id: "1",
    name: "Zaddy's Place",
    image: foodt2,
  },
  {
    id: "1",
    name: "Zaddy's Place",
    image: demofood,
  },
  {
    id: "1",
    name: "Zaddy's Place",
    image: demofood,
  },
  {
    id: "1",
    name: "Zaddy's Place",
    image: foodt1,
  },
]