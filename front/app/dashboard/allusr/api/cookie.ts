"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { mdlAllusrCookieObjson } from "../model/params";

// Handle Cookie
export async function ApiAllusrCookieGetdta() {
  const fnccok = cookies();
  const tknnme = process.env.NEXT_PUBLIC_TKN_COOKIE || "x";
  const tokenx = (await fnccok).get(tknnme)?.value || "";
  const Objusr: mdlAllusrCookieObjson = {
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
    const fnlobj: mdlAllusrCookieObjson = await rspnse.json();
    return fnlobj;
  } catch (error) {
    console.log(error);
  }

  // Return empty object
  return Objusr;
}

// Logout
export async function ApiAllusrCookieLogout() {
  const cookieStore = await cookies();
  cookieStore.delete(`${process.env.NEXT_PUBLIC_TKN_COOKIE}`);
  redirect("/loginp");
}
