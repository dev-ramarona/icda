"use server";

import { MdlViehstPrcessSrcprm } from "../model/params";

export async function ApiViehstPrcessManual(params: MdlViehstPrcessSrcprm) {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/viehst/prcess`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(params),
      cache: "no-store",
    });
    if (!res.ok) return "Failed";
    return "Success";
  } catch (error) {
    console.error(error);
    return "Failed";
  }
}
