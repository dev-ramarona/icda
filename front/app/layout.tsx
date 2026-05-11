import type { Metadata } from "next";
import "./globals.css";
import { geistMono, geistSans, poppins } from "./fonts";

export const metadata: Metadata = {
  title: "Data Analyst Web",
  description: "Created by Data Analyst Lion Tower Internal Control",
  icons: {
    icon: "/lionairblack.png",
  },
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${poppins.className} ${geistSans.className} ${geistMono.className} antialiased`}
      >
        <div className="h-full max-h-full w-screen max-w-full text-xs">
          <div className="afull fixed -z-50 bg-linear-to-br from-zinc-100 via-white to-cyan-100"></div>
          {children}
        </div>
      </body>
    </html>
  );
}
