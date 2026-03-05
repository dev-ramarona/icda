import { MdlGlobalAcpedtDtbase, MdlGlobalActlogDtbase } from "../model/params";

// API accepted edit coloumn
export async function ApiGlobalAcpedtDtbase(params: string) {
  let fnlobj: MdlGlobalAcpedtDtbase[] = [];
  try {
    const rspnse = await fetch(
      `${process.env.NEXT_PUBLIC_URL_SERVER}/global/acpedt/${params}`,
      { method: "GET" },
    );
    if (!rspnse.ok) throw new Error("Failed to fetch accepted edit data");
    fnlobj = await rspnse.json();
  } catch (error) {
    console.log(error);
  }
  return fnlobj;
}

// API action log database
export async function ApiGlobalActlogDtbase(params: string) {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_URL_SERVER}/${params}/actlog/getall`,
      { method: "GET" },
    );
    if (!res.ok) throw new Error("Failed fetch actlog");
    return await res.json();
  } catch (err) {
    console.error(err);
    return { actlog: <MdlGlobalActlogDtbase[]>[], datefl: <string[]>[] };
  }
}
