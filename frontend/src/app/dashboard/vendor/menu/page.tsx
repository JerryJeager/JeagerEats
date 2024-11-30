import Image from "next/image";
import Link from "next/link";
import item from "../../../../../public/assets/jollof.png";
import { MdEdit } from "react-icons/md";
import empty from "../../../../../public/assets/empty.svg"

const Menu = () => {
  return (
    <section className="mt-8 min-h-screen">
      <h2 className="font-bold text-2xl md:text-3xl text-center">
        Menu Management
      </h2>
      <div className="rounded-lg w-full mt-8 p-4 md:p-8 shadow-lg bg-white  h-full md:text-lg ">
        <div className="flex flex-col items-center">
            <Image width={300} src={empty} alt="empty illustration" />
          <p className="text-center">You've not added any menu item yet</p>
          <Link
            href={"/dashboard/vendor/menu/add"}
            className="p-2 md:p-3 w-[220px] mt-6 mx-auto text-center bg-primary text-white rounded-lg"
          >
            Add Menu Item
          </Link>
        </div>
      </div>
      <div className="rounded-lg w-full mt-8 p-4 md:p-8 shadow-lg bg-white  h-full md:text-lg ">
        <div className="">
          <p className="text-center font-semibold">Menu Items</p>
          <div className="mt-2">
            <div className="flex gap-3 border border-slate-300 p-2 border-opacity-90 shadow-md rounded-lg w-fit">
              <Image width={120} src={item} alt="item-image" />
              <div className="flex flex-col w-[170px]">
                <p className="font-medium">A Plate of Rice with Chicken</p>
                <p className="line-clamp-1 text-black text-opacity-70">
                  1 plate of jollof rice, fried rice, plantain and chicken
                </p>
                <div className="flex gap-2 mt-auto">
                  <p className="mb-0">$4000</p>
                  <div className="flex items-center text-primary">
                    <MdEdit />
                    <p>Edit</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Menu;
