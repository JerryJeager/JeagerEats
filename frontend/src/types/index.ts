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
