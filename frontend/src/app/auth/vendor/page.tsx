import Header from "@/components/Header";
import wave from "../../../../public/assets/wave.svg";
import Image from "next/image";
import chef from "../../../../public/assets/chef.png";
import VendorForm from "@/components/auth/VendorForm";

const Vendor = () => {
  return (
    <>
      <Header />
      <section className="mt-10">
        <h2 className="text-2xl md:text-4xl font-bold  padx mb-3">
          We'll let you Cook while we take care of the rest!
        </h2>
        <div className="padx flex justify-between gap-8 relative  z-20">
          <Image
            src={chef}
            className="rounded-2xl shadow-xl hidden md:block"
            alt="chef image"
          />
          <VendorForm />
        </div>
      </section>
      <div className="absolute bottom-0 w-full z-10">
        <Image src={wave} className="w-full" alt="waves" />
        <div className="bg-primary h-[100px]"></div>
      </div>
    </>
  );
};

export default Vendor;
