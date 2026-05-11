"use client";
import { useActionState, useEffect, useState } from "react";
import { mdlAllusrCookieObjson } from "../../model/params";
import { apiAllusrFormipRegist } from "../../api/regis";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { ApiAllusrCookieLogout } from "../../api/cookie";

export default function UixAllusrFormipChpass({ cookie }: { cookie: mdlAllusrCookieObjson }) {
  const rspdef = "Update Password";
  const rplprm = FncGlobalQuerysEdlink();
  const [formac, formacSet, pndingIst] = useActionState(apiAllusrFormipRegist, null);
  const [succss, succssSet] = useState<string>(rspdef);
  const [rspnse, rspnseSet] = useState({
    stfnme_errrsp: formac?.stfnme_errrsp || "",
    stfnme_dfault: formac?.stfnme_dfault || "",
    stfeml_errrsp: formac?.stfeml_errrsp || "",
    stfeml_dfault: formac?.stfeml_dfault || "",
    usrnme_errrsp: formac?.usrnme_errrsp || "",
    usrnme_dfault: formac?.usrnme_dfault || "",
    psswrd_errrsp: formac?.psswrd_errrsp || "",
    psswrd_dfault: formac?.psswrd_dfault || "",
    confrm_errrsp: formac?.confrm_errrsp || "",
    confrm_dfault: formac?.confrm_dfault || "",
    isited_dfault: formac?.isited_dfault || "",
    access_errrsp: formac?.access_errrsp || "",
    access_dfault: formac?.access_dfault || [],
    keywrd_dfault: formac?.keywrd_dfault || [],
    finals_errrsp: formac?.finals_errrsp || "",
  });

  // Monitor form response
  useEffect(() => {
    rspnseSet((prev) => ({
      ...prev,
      stfnme_errrsp: formac?.stfnme_errrsp || "",
      stfnme_dfault: formac?.stfnme_dfault || "",
      stfeml_errrsp: formac?.stfeml_errrsp || "",
      stfeml_dfault: formac?.stfeml_dfault || "",
      usrnme_errrsp: formac?.usrnme_errrsp || "",
      usrnme_dfault: formac?.usrnme_dfault || "",
      psswrd_errrsp: formac?.psswrd_errrsp || "",
      psswrd_dfault: formac?.psswrd_dfault || "",
      confrm_errrsp: formac?.confrm_errrsp || "",
      confrm_dfault: formac?.confrm_dfault || "",
      access_errrsp: formac?.access_errrsp || "",
      access_dfault: formac?.access_dfault || [],
      keywrd_dfault: formac?.keywrd_dfault || [],
      finals_errrsp: formac?.finals_errrsp || "",
    }));
  }, [formac]);

  // Monitor password and confirm password
  useEffect(() => {
    rspnseSet((prev) => ({
      ...prev,
      confrm_errrsp: rspnse.psswrd_dfault != rspnse.confrm_dfault ? "Password not match" : "",
    }));
  }, [rspnse.psswrd_dfault, rspnse.confrm_dfault]);

  // Monitor final output
  useEffect(() => {
    if (rspnse.finals_errrsp == "Success") {
      succssSet("Success update password");
      rplprm("update", String(Math.random()));
      setTimeout(() => ApiAllusrCookieLogout(), 1000);
    } else if (rspnse.finals_errrsp != "") {
      succssSet(rspnse.finals_errrsp);
      setTimeout(() => succssSet(rspdef), 1500);
    }
  }, [rspnse.finals_errrsp]);
  return (
    <form
      className="h-fit w-fit rounded-md bg-white p-6 shadow-md ring-2 ring-gray-200"
      action={formacSet}
    >
      <div className="flexctr h-full w-full flex-col gap-3 md:gap-6">
        <input type="hidden" name="stfnme" id="stfnme" defaultValue={cookie.stfnme} />
        <input type="hidden" name="stfeml" id="stfeml" defaultValue={cookie.stfeml} />
        <input type="hidden" name="usrnme" id="usrnme" defaultValue={cookie.usrnme} />
        <div className="relative flex w-full flex-col items-start justify-center gap-1.5">
          <label className="font-medium" htmlFor="psswrd">
            New Password
          </label>
          <input
            className={`w-full rounded-md px-3 py-1.5 ring ring-gray-300`}
            type="password"
            name="psswrd"
            id="psswrd"
            onChange={(e) =>
              rspnseSet((prev) => ({ ...prev, psswrd_dfault: e.target.value, psswrd_errrsp: "" }))
            }
          />
          <span className={`absolute right-2 -bottom-4 text-[10px] text-red-600`}>
            {rspnse.psswrd_errrsp}
          </span>
        </div>
        <div className="relative flex w-full flex-col items-start justify-center gap-1.5">
          <label className="font-medium" htmlFor="confrm">
            Confirm New Password
          </label>
          <input
            className={`w-full rounded-md px-3 py-1.5 ring ring-gray-300`}
            type="password"
            name="confrm"
            id="confrm"
            onChange={(e) => rspnseSet((prev) => ({ ...prev, confrm_dfault: e.target.value }))}
          />
          <span className={`absolute right-2 -bottom-4 text-[10px] text-red-600`}>
            {rspnse.confrm_errrsp}
          </span>
        </div>
        <input name="isited" id="isited" defaultValue="update" hidden />
        <input name="action" id="action" defaultValue="update" hidden />
        {cookie.access.map((item, index) => (
          <input key={index} type="hidden" name="access" id="access" defaultValue={item} />
        ))}
        {cookie.keywrd.map((item, index) => (
          <input key={index} type="hidden" name="keywrd" id="keywrd" defaultValue={item} />
        ))}
        <button
          // "Success update password"
          className={`${succss == rspdef ? "btnsbm" : succss == "Success update password" ? "btnscs" : "btncxl"} w-full py-1.5 text-center`}
          type="submit"
          disabled={pndingIst}
        >
          {pndingIst ? "Loading..." : succss}
        </button>
      </div>
    </form>
  );
}
