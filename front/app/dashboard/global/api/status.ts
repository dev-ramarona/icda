"use server";
import { MdlGlobalStatusPrcess } from "../model/params";

// Hit status sabre api
export async function ApiGlobalStatusPrcess() {
  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_URL_SERVER}/global/status`,
      {
        method: "GET",
        credentials: "include",
      },
    );
    if (!response.ok) throw new Error("Failed to fetch");
    const rawstr: MdlGlobalStatusPrcess = await response.json();
    const fnlstr: MdlGlobalStatusPrcess = {
      action: Number(rawstr.action.toFixed(2)),
      sbrapi: Number(rawstr.sbrapi.toFixed(2)),
    };
    return fnlstr;
  } catch (error) {
    console.log(error);
  }

  return { action: 0, sbrapi: 0 };
}

// Hit status api with interval time
// export async function ApiGlobalStatusIntrvl(
//   statfnSet: (v: string) => void,
//   strVarble: "action" | "sbrapi",
// ) {
//   const strtiv = setInterval(async () => {
//     const status = await ApiGlobalStatusPrcess();
//     const rawval = strVarble == "action" ? status.action : status.sbrapi;
//     const nowval = Number(rawval.toFixed(2));
//     const nowstr = nowval == 0 ? "Done" : `${nowval}%`;
//     statfnSet(nowstr);
//     if (nowstr === "Done") {
//       clearInterval(strtiv);
//       statfnSet("Process Done");
//       setTimeout(() => statfnSet("Done"), 1000);
//     }
//   }, 3000);
// }
