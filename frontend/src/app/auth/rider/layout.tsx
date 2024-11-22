import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "JeagerEats - Rider Signup",
  description: "Let's take care of chow agenda for you.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return <main className="bg-[#FFF1EA] relative pb-8">{children}</main>;
}
