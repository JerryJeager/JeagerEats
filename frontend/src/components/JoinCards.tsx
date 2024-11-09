import { Join } from "@/data";
import Image from "next/image";
import Link from "next/link";

const JoinCards = ({ title, content, cta, icon, link, color }: Join) => {
  return (
    <div
      style={{ backgroundColor: color }}
      className={`relative rounded-xl border-black border-2 lg:border-4 p-4 max-w-[350px]`}
    >
      <h3 className="font-semibold text-xl">{title}</h3>
      <p className="mt-4">{content}</p>
      <button className="py-2 px-4 mt-6 bg-black text-white rounded-lg">
        <Link href={link}>{cta}</Link>
      </button>
      <div
        style={{ backgroundColor: color }}
        className={`rounded-full border-black border-2 lg:border-4 p-2 absolute w-[70px] h-[70px] top-0 right-0 -translate-x-[-50%] translate-y-[-50%] hidden lg:block`}
      >
        <Image src={icon} alt="join icon" className="w-full h-full" />
      </div>
    </div>
  );
};

export default JoinCards;
