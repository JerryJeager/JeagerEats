import Header from "@/components/Header";
import wave from "../../../../public/assets/wave.svg";
import Image from "next/image";
import serve from "../../../../public/assets/serve.jpg"
import SignupForm from "@/components/auth/SignupForm";

const Customer = () => {
  return (
    <>
      <Header />
      <section className="mt-10">
        <h2 className="text-2xl md:text-4xl font-bold  padx mb-3">
          Let's take care of chow agenda for you
        </h2>
        <div className="padx flex justify-between gap-8 relative  z-20">
          <SignupForm name="customer"  />
          <Image
            src={serve}
            placeholder="blur"
            className="rounded-2xl h-fit shadow-xl hidden md:block w-1/2"
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

export default Customer;
