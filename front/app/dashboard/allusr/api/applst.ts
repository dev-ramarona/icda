import { MdlAllusrApplstParams } from "../model/params";

// API Applist
export async function ApiAllusrApplstDtbase(ignore: number[]) {
  const fnlObj: MdlAllusrApplstParams[] = [];
  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_URL_SERVER}/allusr/applst`,
      {
        method: "GET",
        credentials: "include",
      },
    );
    if (!response.ok) throw new Error("Failed to fetch app list");
    const tmpObj: MdlAllusrApplstParams[] = await response.json();
    const fnlObj = tmpObj.filter((item) => !ignore.includes(item.pagenb));
    return fnlObj;
  } catch (error) {
    console.log(error);
  }
  return fnlObj;
}
