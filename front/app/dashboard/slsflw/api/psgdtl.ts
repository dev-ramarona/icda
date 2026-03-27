"use server";
import { MdlSlsflwPsgdtlFrntnd, MdlSlsflwPsgdtlSrcprm } from "../model/params";

// Function get psglst database
export async function ApiSlsflwPsgdtlGetall(prmPsgdtl: MdlSlsflwPsgdtlSrcprm) {
  const tag = [
    "psgdtl",
    prmPsgdtl.update_global,
    prmPsgdtl.mnthfl_psgdtl,
    prmPsgdtl.datefl_psgdtl,
    prmPsgdtl.airlfl_psgdtl,
    prmPsgdtl.flnbfl_psgdtl,
    prmPsgdtl.depart_psgdtl,
    prmPsgdtl.routfl_psgdtl,
    prmPsgdtl.pnrcde_psgdtl,
    prmPsgdtl.tktnfl_psgdtl,
    prmPsgdtl.nclear_psgdtl,
    prmPsgdtl.pagenw_psgdtl,
  ]
    .filter(Boolean)
    .join(":");
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/psglst/psgdtl/getall/slsflw`, {
      method: "POST",
      body: JSON.stringify(prmPsgdtl),
      headers: { "Content-Type": "application/json" },
      next: { revalidate: 30, tags: [tag] },
    });
    if (!res.ok) throw new Error("Failed fetch psgdtl");
    return await res.json();
  } catch (err) {
    console.error(err);
    return { arrdta: [], totdta: 0 };
  }
}

// Function get psglst database
export async function ApiSlsflwPsgdtlUpdate(params: MdlSlsflwPsgdtlFrntnd): Promise<string> {
  // Validation
  if (params.ntafvc === 0) return "NTA empty";
  if (params.curncy === "") return "curncy empty";

  // Call API
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/psglst/psgdtl/update/slsrpt`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(params),
      cache: "no-store",
    });
    if (!res.ok) {
      throw new Error("Failed update psgdtl");
    }
    const data = await res.json();
    return await data;
  } catch (error) {
    console.error(error);
    return "update failed";
  }
}
