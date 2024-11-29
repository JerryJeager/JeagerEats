import Sidebar from "@/components/dashboard/Sidebar";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "JeagerEats - Vendor Dashboard",
  description: "Let's take care of chow agenda for you.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <main className="bg-[#FFF1EA] relative pb-8 md:h-screen flex justify-between  md:pl-0 w-full md:overflow-hidden">
      <div className="md:w-[25%]">
        <Sidebar />
      </div>
      <div className="w-full md:w-[85%] md:mx-[5%] md:overflow-y-scroll scroll-cs">{children}</div>
    </main>
  );
}
