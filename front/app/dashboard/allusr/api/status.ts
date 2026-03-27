"use server";

import { MdlAllusrStatusPrcess } from "../model/params";

// Hit status sabre api
export async function ApiAllusrStatusPrcess() {
  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/allusr/status`, {
      method: "GET",
      credentials: "include",
    });
    if (!response.ok) throw new Error("Failed to fetch");
    const rawstr: MdlAllusrStatusPrcess = await response.json();
    const fnlstr: MdlAllusrStatusPrcess = {
      action: rawstr.action,
      sbrapi: Number(rawstr.sbrapi.toFixed(2)),
    };
    return fnlstr;
  } catch (error) {
    console.log(error);
  }

  return { action: "", sbrapi: 0 };
}
