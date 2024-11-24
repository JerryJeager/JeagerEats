import Header from "@/components/Header";
import wave from "../../../../public/assets/wave.svg";
import Image from "next/image";
import LoginForm from "@/components/auth/LoginForm";
import welcome from "../../../../public/assets/welcome.jpg"

const Login = () => {
  return (
    <>
      <Header />
      <section className="mt-10">
        <h2 className="text-2xl md:text-4xl font-bold  padx mb-3">
          Welcome Back!
        </h2>
        <div className="padx flex flex-row-reverse justify-between gap-8 relative  z-20">
          <Image
            src={welcome}
            placeholder="blur"
            className="rounded-2xl w-1/2 h-[600px] shadow-xl hidden md:block"
            alt="chef image"
          />
          <LoginForm />
        </div>
      </section>
      <div className="absolute bottom-0 w-full z-10">
        <Image src={wave} className="w-full" alt="waves" />
        <div className="bg-primary h-[100px]"></div>
      </div>
    </>
  );
};

export default Login;
