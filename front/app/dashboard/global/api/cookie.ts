"use server";

import { cookies } from "next/headers";
import { mdlGlobalAllusrCookie } from "../model/params";
import { redirect } from "next/navigation";

// Handle Cookie
export async function ApiGlobalCookieGetdta() {
  const fnccok = cookies();
  const tknnme = process.env.NEXT_PUBLIC_TKN_COOKIE || "x";
  const tokenx = (await fnccok).get(tknnme)?.value || "";
  const Objusr: mdlGlobalAllusrCookie = {
    stfnme: "",
    usrnme: "",
    stfeml: "",
    access: ["null"],
    keywrd: ["null"],
  };
  if (tokenx == "" || !tokenx) return Objusr;

  // Try hit API
  try {
    const rspnse = await fetch(
      `${process.env.NEXT_PUBLIC_URL_SERVER}/allusr/tokenx`,
      {
        method: "GET",
        headers: {
          Authorization: tokenx,
        },
        credentials: "include",
      },
    );
    if (!rspnse.ok) throw new Error("Failed to fetch user data");
    const fnlobj: mdlGlobalAllusrCookie = await rspnse.json();
    return fnlobj;
  } catch (error) {
    console.log(error);
  }

  // Return empty object
  return Objusr;
}

// Logout
export async function ApiGlobalCookieLogout() {
  const cookieStore = await cookies();
  cookieStore.delete(`${process.env.NEXT_PUBLIC_TKN_COOKIE}`);
  redirect("/loginp");
}
