import Header from "@/components/shop/Header";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "JeagerEats - Shop",
  description: "Let's take care of chow agenda for you.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <main className="relative pb-8">
      <div className="shadow-md">
        <Header />
      </div>
      {children}
    </main>
  );
}
