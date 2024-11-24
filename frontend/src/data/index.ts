import bike from "../../public/assets/Motorcycle.png";
import shop from "../../public/assets/shop.png";
import cutlery from "../../public/assets/cutlery.png";
import { StaticImageData } from "next/image";

export const BASE_URL = () => {
  return process.env.NEXT_PUBLIC_ENVIRONMENT == "production"
    ? "https://jeagereats-production.up.railway.app/api/v1"
    : "http://localhost:8080/api/v1";
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
    link: "",
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
