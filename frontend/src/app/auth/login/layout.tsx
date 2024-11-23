import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "JeagerEats - Login",
  description: "Let's take care of chow agenda for you.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return <main className="bg-[#FFF1EA] relative pb-8 h-screen">{children}</main>;
}
