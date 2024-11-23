import Header from "@/components/Header";
import wave from "../../../../public/assets/wave.svg";
import Image from "next/image";
import rider from "../../../../public/assets/rider.png";
import SignupForm from "@/components/auth/SignupForm";

const Rider = () => {
  return (
    <>
      <Header />
      <section className="mt-10">
        <h2 className="text-2xl md:text-4xl font-bold  padx mb-3">
          Earn while delivering happiness
        </h2>
        <div className="padx flex justify-between gap-8 relative  z-20">
          <SignupForm name="rider" />
          <Image
            src={rider}
            placeholder="blur"
            className="rounded-2xl h-fit shadow-xl hidden md:block"
            alt="chef image"
          />
        </div>
      </section>
      <div className="absolute bottom-0 w-full z-10">
        <Image src={wave} className="w-full" alt="waves" />
        <div className="bg-primary h-[100px]"></div>
      </div>
    </>
  );
};

export default Rider;
