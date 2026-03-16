import { MdlGlobalActlogDtbase } from "../model/params";

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
