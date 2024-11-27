import Image from "next/image";
import logo from "../../../public/assets/logo.png";
import { vendorDashboardNav } from "@/data";
import Link from "next/link";
const Sidebar = () => {
  return (
    <div className="h-full border  bg-white">
      <div className="shadow-md">
        <Image alt="logo" className="p-8" src={logo} />
      </div>
      <div className="p-8 mt-10">
        <ul>
          {vendorDashboardNav.map((v, index) => (
            <li key={index} className="mb-8">
              <Link href={v.link} className="flex items-center gap-3 md:text-xl">
                <v.icon />
                <p>{v.name}</p>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Sidebar;
