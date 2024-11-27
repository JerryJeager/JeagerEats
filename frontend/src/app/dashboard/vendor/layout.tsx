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
    <main className="bg-[#FFF1EA] relative pb-8 h-screen flex gap-8 padx md:pl-0 ">
      <div className="w-[25%]">
        <Sidebar />
      </div>
      <div className="w-[60$]">{children}</div>
    </main>
  );
}
