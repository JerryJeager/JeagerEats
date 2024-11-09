import Header from "@/components/Header";
import Contact from "@/components/home/Contact";
import Hero from "@/components/home/Hero";
import Join from "@/components/home/Join";
import Restaurants from "@/components/home/Restaurants";
import Footer from "../components/home/Footer";

export default function Home() {
  return (
    <div>
      <Header />
      <Hero />
      <Join />
      <Restaurants />
      <Contact />
      <Footer />
    </div>
  );
}
