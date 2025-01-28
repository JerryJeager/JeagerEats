"use client";
import Image from "next/image";
import { riderDashboardNav } from "@/data";
import logo from "../../../../public/assets/logo.png"
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useState } from "react";
const Sidebar = () => {
  const pathname = usePathname();
  const [isHamburgerClicked, setIsHamburgerClicked] = useState(false);
  const handleClickedHamburger = () => {
    setIsHamburgerClicked((preValue) => !preValue);
  };
  return (
    <>
      <div
        className={`md:block md:bg-white md:h-screen ${
          isHamburgerClicked
            ? "z-10  bg-white top-0 overflow-y-auto h-full  px-4 md:px-0  pt-[100px] md:pt-0 fixed md:static left-[0%] "
            : "hidden left-[-100%]"
        }`}
      >
        <div className="md:shadow-md ">
          <Image alt="logo" className="md:py-4 mx-3" width={140} src={logo} />
        </div>
        <div className="p-8 mt-10 ">
          <ul>
            {riderDashboardNav.map((v, index) => (
              <li
                key={index}
                className={`mb-8 ${
                  pathname == v.link ? "text-primary" : "text-black"
                }`}
              >
                <Link
                  href={v.link}
                  onClick={() => setIsHamburgerClicked(false)}
                  className="flex items-center gap-3 md:text-xl"
                >
                  <v.icon />
                  <p>{v.name}</p>
                </Link>
              </li>
            ))}
          </ul>
        </div>
      </div>
      {/* hamburger menu icon */}
      <div
        onClick={handleClickedHamburger}
        className={`mt-8 ml-4 cursor-pointer top-0 md:hidden ${
          isHamburgerClicked ? "z-20 fixed" : "absolute"
        }`}
      >
        <div
          className={`h-1 w-[27px] rounded-sm bg-black duration-500 ${
            isHamburgerClicked
              ? "translate-x-[-4.5px] translate-y-[6px] rotate-[-405deg] bg-black "
              : "bg-black"
          } `}
        ></div>
        <div
          className={`h-1 w-[27px] rounded-sm bg-black mt-2 duration-500 ${
            isHamburgerClicked
              ? "translate-x-[-4.5px] translate-y-[-6px] rotate-[405deg] bg-black"
              : "bg-black"
          }`}
        ></div>
      </div>
    </>
  );
};

export default Sidebar;
