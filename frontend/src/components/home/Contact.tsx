import Link from "next/link";
import { comforta } from "./Hero";
// import arrow from "../../../public/assets/arrow.png"
// import Image from "next/image";

const Contact = () => {
  return (
    <section className="mt-8 pt-8 lg:pt-10 pb-10  lg:pb-14 px-[5%] lg:px-[8%]">
      <h2
        className={`text-2xl lg:text-4xl text-primary ${comforta.className} text-center lg:max-w-[600px] mx-auto`}
      >
        Get in Touch
      </h2>
      <p className="lg:text-lg text-center lg:max-w-[500px] mx-auto">Have questions, feedback, or need assistance? Reach out to us, and we’ll get back to you as soon as possible.</p>
      <div className="bg-primary rounded-full w-full h-[100px] lg:h-[150px] flex items-center justify-center text-white mt-4">
        <Link href={"/"} className="flex gap-4">
            <p className="text-2xl lg:text-4xl">Contact Us</p>
            {/* <Image src={arrow} alt="arrow" width={30} /> */}
        </Link>
      </div>
    </section>
  );
};

export default Contact;
