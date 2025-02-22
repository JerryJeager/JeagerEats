import { StaticImageData } from "next/image";
import { IconType } from "react-icons";

export type Roles = "customer" | "vendor" | "rider";

export type Role = {
  name: Roles;
};

export type VendorDashboarNav = {
  name: string;
  icon: IconType;
  link: string;
};

export type User = {
  id: string;
  email: string;
  first_name: string;
  last_name: string;
  role: Roles;
  phone_number: string;
  address: string;
  created_at: string;
};

export type Restaurant = {
  id: string;
  name: string;
  description: string;
  address: string;
  profile_img: string;
  rating: string;
  is_active: boolean;
  cuisine_type: string;
  opening_time: string;
  closing_time: string;
  phone_number: string;
  created_at: string;
};

export type RestaurantSelf = {
  id: string;
  name: string;
  user_id: string;
  description: string;
  address: string;
  profile_img: string;
  rating: string;
  is_active: boolean;
  cuisine_type: string;
  opening_time: string;
  closing_time: string;
  created_at: string;
};

export type MenuItem = {
  name: string;
  description: string;
  price: number;
  stock: number;
};

export type MenuItemCardType = {
  id: string;
  restaurant_id: string;
  name: string;
  description: string;
  price: number;
  is_available: boolean;
  img_url: string;
  stock: number;
  category: string;
};

export type CategoryCardType = {
  color: string;
  image: StaticImageData;
  name: string;
};

export type RestaurantCardType = {
  id: string;
  name: string;
  description: string;
  address: string;
  profile_img: string;
  rating: number;
  is_active: boolean;
  cuisine_type: string;
  opening_time: string;
  closing_time: string;
  created_at: string;
  updated_at: string;
};

export type RestaurantMenuCardType = {
  id: string;
  restaurant_id: string;
  name: string;
  description: string;
  price: number;
  is_available: boolean;
  img_url: string;
  stock: number;
  category: string;
  quantity: number
  created_at: string;
  updated_at: string;
  deleted_at: string;
};

export type PlaceOrderType = {
  restaurant_id: string 
  total_price: number
  delivery_fee: number 
  delivery_address: string
  items: {
    menu_id: string 
    price_per_item: number 
    quantity: number
  }[]
}