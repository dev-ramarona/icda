import type { Metadata } from "next";
import { Geist, Geist_Mono, Poppins } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

const poppins = Poppins({
  weight: ["100", "200", "300", "400", "500", "600", "700", "800", "900"],
  variable: "--font-ubuntu",
  subsets: ["latin"],
});

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
        <div className="w-screen max-w-full h-full max-h-full text-xs">
          <div className="afull fixed bg-linear-to-br from-zinc-100 via-white to-cyan-100 -z-50"></div>
          {children}
        </div>
      </body>
    </html>
  );
}
