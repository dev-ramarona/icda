import { MdlAllusrSearchParams } from "../model/params";

// Function get all user
export async function ApiAllusrUsrlstGetall(prmAllusr: MdlAllusrSearchParams) {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/allusr/getall`, {
      method: "POST",
      body: JSON.stringify(prmAllusr),
      headers: { "Content-Type": "application/json" },
    });
    if (!res.ok) throw new Error("Failed fetch allusr");
    return await res.json();
  } catch (err) {
    console.error(err);
    return { arrdta: [], totdta: 0 };
  }
}

// Function get all user
export async function ApiAllusrHandleDelete(usrnme: string) {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/allusr/delete/${usrnme}`, {
      method: "GET",
    });
    if (!res.ok) throw new Error("Failed fetch allusr");
    return true;
  } catch (err) {
    console.error(err);
    return false;
  }
}
