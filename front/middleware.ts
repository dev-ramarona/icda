import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { mdlAllusrCookieObjson } from "./app/dashboard/allusr/model/params";

export async function middleware(req: NextRequest) {
  const tknnme = process.env.NEXT_PUBLIC_TKN_COOKIE || "x";
  const tokenx = req.cookies.get(tknnme)?.value || "";
  const pathnm = req.nextUrl.pathname.split("/")[2];

  // Jika belum login, arahkan ke "/"
  if (tokenx == "" || !tokenx) {
    return NextResponse.redirect(new URL("/loginp", req.url));
  }

  // Try hit API
  try {
    const rspnse = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/allusr/tokenx`, {
      method: "GET",
      headers: {
        Authorization: tokenx,
      },
      credentials: "include",
    });
    if (!rspnse.ok) throw new Error("Failed to register user");
    const fnlobj: mdlAllusrCookieObjson = await rspnse.json();
    if (pathnm && !fnlobj.access.includes(pathnm))
      return NextResponse.redirect(new URL("/dashboard/global", req.url));
    return NextResponse.next();
  } catch (error) {
    console.log(error);
    return NextResponse.redirect(new URL("/loginp", req.url));
  }
}

// Menerapkan middleware ke semua route di bawah '/global'
export const config = {
  // matcher:
  //   "/((?!$|_next|favicon.ico|.*\\.(?:png|jpg|jpeg|gif|svg|ico|webp|css|js)).*)",
  matcher: ["/dashboard/:path*"],
};
