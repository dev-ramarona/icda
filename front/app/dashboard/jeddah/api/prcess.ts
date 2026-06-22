"use server";

import { MdlJeddahGlobalSrcprm } from "../model/params";

export async function ApiJeddahPrcessManual(params: MdlJeddahGlobalSrcprm) {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/jeddah/prcess`, {
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

// Function get psglst database
export async function ApiJeddahPnrsmrGetall(prmPsgsmr: MdlJeddahGlobalSrcprm) {
  const tag = [
    "psgsmr",
    prmPsgsmr.update_global,
    prmPsgsmr.airlfl_jeddah,
    prmPsgsmr.flnbfl_jeddah,
    prmPsgsmr.depart_jeddah,
    prmPsgsmr.routfl_jeddah,
    prmPsgsmr.pnrcde_jeddah,
  ]
    .filter(Boolean)
    .join(":");
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/jeddah/pnrsmr/getall`, {
      method: "POST",
      body: JSON.stringify(prmPsgsmr),
      headers: { "Content-Type": "application/json" },
      next: { revalidate: 30, tags: [tag] },
    });
    if (!res.ok) throw new Error("Failed fetch psgsmr");
    return await res.json();
  } catch (err) {
    console.error(err);
    return { arrdta: [], totdta: 0 };
  }
}
