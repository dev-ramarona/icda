"use server";
import { MdlPsglstErrlogDtbase } from "../model/params";

export async function ApiPsglstPrcessManual(params: MdlPsglstErrlogDtbase) {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_URL_SERVER}/psglst/prcess`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(params),
        cache: "no-store",
      },
    );
    if (!res.ok) return "failed";
    return "success";
  } catch (error) {
    console.error(error);
    return "update failed";
  }
}
