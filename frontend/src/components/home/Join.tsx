import { JoinData } from "@/data";
import { comforta } from "./Hero";
import JoinCards from "../JoinCards";

const Join = () => {
  return (
    <section className="px-[5%] lg:px-[8%] py-4 lg:py-6 mt-4">
      <h2
        className={`text-2xl lg:text-4xl mb-8 lg:mb-10 text-primary ${comforta.className} text-center`}
      >
        Join the JeagerEats Community
      </h2>

      <div className="flex flex-col gap-4 lg:gap-0 lg:flex-row justify-between  lg:mt-14 ">
        {JoinData.map((j, key) => (
          <JoinCards key={key} {...j} />
        ))}
      </div>
    </section>
  );
};

export default Join;
