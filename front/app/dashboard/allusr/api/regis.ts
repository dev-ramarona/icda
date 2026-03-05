"use server";

// import { cookies } from "next/headers";
export async function apiAllusrFormipRegist(
  prevState: any,
  formData: FormData,
) {
  const action = formData.get("action") as string;
  const stfnme = formData.get("stfnme") as string;
  const stfeml = formData.get("stfeml") as string;
  const usrnme = formData.get("usrnme") as string;
  const psswrd = formData.get("psswrd") as string;
  const confrm = formData.get("confrm") as string;
  const access = formData.getAll("access") as string[];
  const keywrd = formData.getAll("keywrd") as string[];
  const rawobj = {
    stfnme_errrsp: "",
    stfnme_dfault: "",
    stfeml_errrsp: "",
    stfeml_dfault: "",
    usrnme_errrsp: "",
    usrnme_dfault: "",
    psswrd_errrsp: "",
    psswrd_dfault: "",
    access_errrsp: "",
    access_dfault: [],
    keywrd_dfault: [],
    finals_errrsp: "",
  };

  // Check validation data
  const errobj = { ...rawobj };
  errobj.stfnme_dfault = stfnme;
  errobj.stfeml_dfault = stfeml;
  errobj.usrnme_dfault = usrnme;
  errobj.psswrd_dfault = psswrd;
  errobj.access_dfault = access;
  errobj.keywrd_dfault = keywrd;
  if (!stfnme || !stfeml || !usrnme || !psswrd || !confrm || !access.length) {
    if (!stfnme) errobj.stfnme_errrsp = "Staff name is empty";
    if (!stfeml) errobj.stfeml_errrsp = "Staff email is empty";
    if (!usrnme) errobj.usrnme_errrsp = "Username is empty";
    if (!psswrd || !confrm) errobj.psswrd_errrsp = "Password is empty";
    if (!access.length) errobj.access_errrsp = "Access is empty";
    if (!keywrd.length) errobj.access_errrsp = "Keyword is empty";
    errobj.finals_errrsp = "Check the data";
    return errobj;
  }

  // Hit api
  let res;
  try {
    res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/allusr/regist`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        action: action,
        usrnme: usrnme,
        stfnme: stfnme,
        stfeml: stfeml,
        psswrd: psswrd,
        access: access,
        keywrd: keywrd,
      }),
    });
  } catch (error) {
    console.error("Fetch failed:", error);
    errobj.finals_errrsp = "Server is not available.";
    return errobj;
  }

  // Get token credential
  const data = await res.json();
  if (!res.ok) {
    if (data.error === "user") {
      errobj.finals_errrsp = "Username already exists";
    } else errobj.finals_errrsp = "Invalid credential";
    return errobj;
  } else {
    rawobj.finals_errrsp = "Success";
    return rawobj;
  }
}
