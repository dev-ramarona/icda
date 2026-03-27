"use server";
import { MdlApndixSearchQueryx } from "../model/parmas";

// API Applist
export async function ApiApndixApplstDtbase() {
  const fnlObj: string[] = [];
  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/apndix/applst/getall`, {
      method: "GET",
      credentials: "include",
    });
    if (!response.ok) throw new Error("Failed to fetch app list");
    const fnlObj: string[] = await response.json();
    return fnlObj;
  } catch (error) {
    console.log(error);
  }
  return fnlObj;
}

// API accepted edit coloumn
export async function ApiApndixGetallDtbase(params: MdlApndixSearchQueryx) {
  // await new Promise((r) => setTimeout(r, 2000));
  // const tag = [
  //   params.pagedb_apndix,
  //   params.datefl_apndix,
  //   params.airlfl_apndix,
  //   params.depart_apndix,
  //   params.flnbfl_apndix,
  //   params.routfl_apndix,
  //   params.clssfl_apndix,
  //   params.pagenw_apndix,
  // ]
  //   .filter(Boolean)
  //   .join(":");
  try {
    const rspnse = await fetch(
      `${process.env.NEXT_PUBLIC_URL_SERVER}/apndix/${params.pagedb_apndix}/getall`,
      {
        method: "POST",
        body: JSON.stringify(params),
        headers: { "Content-Type": "application/json" },
        // next: { revalidate: 30, tags: [tag] },
      },
    );
    if (!rspnse.ok) throw new Error("Failed to fetch accepted edit data");
    const jsdata = await rspnse.json();
    const arrdta = jsdata.arrdta.map((row: any) =>
      Object.fromEntries(Object.entries(row).map(([k, v]) => [k, String(v)])),
    );
    return { arrdta: arrdta, totdta: jsdata.totdta };
  } catch (error) {
    console.log(error);
  }
  return { arrdta: [], totdta: 0 };
}

// Api update apendix data
export async function ApiApndixUpdateDtbase(objupd: any, target: string): Promise<string> {
  // Call API
  console.log(JSON.stringify(objupd));

  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/apndix/${target}/update`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(objupd),
      cache: "no-store",
    });
    if (!res.ok) {
      console.log(res);
      throw new Error("Failed update psgdtl");
    }
    const data = await res.json();
    return await data;
  } catch (error) {
    console.error(error);
    return "update failed";
  }
}

// API accepted edit coloumn
export async function ApiApndixAcpedtDtbase(params: string) {
  let fnlobj: any[] = [];
  try {
    const rspnse = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/apndix/acpedt/${params}`, {
      method: "GET",
    });
    if (!rspnse.ok) throw new Error("Failed to fetch accepted edit data");
    fnlobj = await rspnse.json();
  } catch (error) {
    console.log(error);
  }
  return fnlobj;
}
