import { Pacifico, Comfortaa } from "next/font/google";
import jollof from "../../../public/assets/jollof.png";
import spag from "../../../public/assets/spag.png";
import bikeman from "../../../public/assets/bikeman.png";
import Image from "next/image";

const pacifico = Pacifico({
  subsets: ["latin"],
  display: "swap",
  weight: ["400"],
});

export const comforta = Comfortaa({
  subsets: ["latin"],
  display: "swap",
  weight: ["400", "500", "600", "700"],
});
const Hero = () => {
  return (
    <section className="px-[5%] lg:px-[8%] py-14 lg:py-20 bg-hero flex flex-col lg:flex-row justify-between">
      <div className="lg:w-2/3">
        <p className={`${pacifico.className} text-sm`}>Fast & Reliable Food Delivery</p>
        <h1 className={`${comforta.className} text-white text-3xl lg:text-6xl font-black mt-2`}>
          Get your favorite dishes delivered at lightning speed, so you can focus on what truly matters
        </h1>
        <button className="bg-white rounded-lg border-black border-2 lg:border-4 text-primary text-lg py-2 px-8 mt-6 font-semibold">Order Now</button>
      </div>
      <div className="mt-4 lg:mt-0">
        <div className="flex self-end">
          <div className="relative mt-[30%] translate-x-1/3">
            <Image src={spag} alt="spag image" className="rounded-xl" />
            <Image
              src={bikeman}
            //   width={50}
              alt="jollof rice image"
              className="rounded-full absolute top-0 right-0 -translate-x-[-50%] w-[100px] h-[100px]"
            />
          </div>
          <div>
            <Image src={jollof} alt="bikeman image" className="rounded-xl" />
          </div>
        </div>
      </div>
    </section>
  );
};

export default Hero;
