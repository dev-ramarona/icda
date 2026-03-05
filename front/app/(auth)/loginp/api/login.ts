"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

// import { cookies } from "next/headers";
export async function apiLoginpFormpgLoginx(
  prevState: any,
  formData: FormData,
) {
  const usrnme = formData.get("usrnme") as string;
  const psswrd = formData.get("psswrd") as string;
  const errobj = { rspnse: "", dfault: usrnme };
  if (!usrnme || !psswrd) {
    errobj.rspnse = "User or password is empty";
    return errobj;
  }

  // Hit api
  let res;
  try {
    res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/allusr/loginx`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({ usrnme: usrnme, psswrd: psswrd }),
    });
  } catch (error) {
    console.error("Fetch failed:", error);
    errobj.rspnse = "Server is not available.";
    return errobj;
  }

  // Get token credential
  const data = await res.json();
  if (!res.ok) {
    if (data.error === "user") {
      errobj.rspnse = "User or password is false";
    } else errobj.rspnse = "Invalid credential";
    return errobj;
  }

  // Store token
  const cookieStore = await cookies();
  const tknnme = process.env.NEXT_PUBLIC_TKN_COOKIE || "x";
  cookieStore.set(tknnme, data.ok, {
    httpOnly: true,
    path: "/",
    maxAge: 10800,
  });
  redirect("/dashboard/global");
}
