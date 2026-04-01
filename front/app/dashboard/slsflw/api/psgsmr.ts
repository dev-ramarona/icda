import { MdlSlsflwPsgsmrSrcprm } from "../model/params";

// Function get psglst database
export async function ApiSlsflwPsgsmrGetall(prmPsgsmr: MdlSlsflwPsgsmrSrcprm) {
  const tag = [
    "psgsmr",
    prmPsgsmr.update_global,
    prmPsgsmr.mnthfl_psgsmr,
    prmPsgsmr.datefl_psgsmr,
    prmPsgsmr.airlfl_psgsmr,
    prmPsgsmr.flnbfl_psgsmr,
    prmPsgsmr.depart_psgsmr,
    prmPsgsmr.routfl_psgsmr,
    prmPsgsmr.pagenw_psgsmr,
    prmPsgsmr.limitp_psgsmr,
  ]
    .filter(Boolean)
    .join(":");
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/psglst/psgsmr/getall`, {
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
