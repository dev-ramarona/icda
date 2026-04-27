"use client";
import { useActionState, useEffect, useState } from "react";
import { UixGlobalIconvcCancel } from "../../../global/ui/server/iconvc";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { MdlAllusrApplstParams, MdlAllusrSearchParams } from "../../model/params";
import { apiAllusrFormipRegist } from "../../api/regis";

export default function UixAllusrFormipMainpg({
  applst,
  prmAllusr,
}: {
  applst: MdlAllusrApplstParams[];
  prmAllusr: MdlAllusrSearchParams;
}) {
  // Declare variable
  const rspscs = "Success create user";
  const rspdef = "Create user";
  const rplprm = FncGlobalQuerysEdlink();
  const [formac, formacSet, pndingIst] = useActionState(apiAllusrFormipRegist, null);
  const [keywrd, keywrdSet] = useState<string>("");
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

  // Function
  const keywrdRmv = (value: string) =>
    rspnseSet((prev) => ({
      ...prev,
      keywrd_dfault: prev.keywrd_dfault.filter((prv) => prv !== value),
    }));
  const keywrdAdd = () => {
    if (!keywrd.trim()) return;
    rspnseSet((prev) => ({
      ...prev,
      keywrd_dfault: [...prev.keywrd_dfault, keywrd.trim()],
    }));
    keywrdSet("");
  };

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

  // Monitor edit user params
  useEffect(() => {
    succssSet(rspdef);
    if (prmAllusr.usredt) {
      succssSet("Update user");
      const objUsredt = prmAllusr.usredt ? JSON.parse(prmAllusr.usredt) : null;
      rspnseSet((prev) => ({
        ...prev,
        stfnme_dfault: objUsredt?.params?.stfnme || "",
        stfeml_dfault: objUsredt?.params?.stfeml || "",
        usrnme_dfault: objUsredt?.params?.usrnme || "",
        psswrd_dfault: objUsredt?.params?.psswrd || "",
        confrm_dfault: objUsredt?.params?.confrm || "",
        access_dfault: objUsredt?.params?.access || [],
        keywrd_dfault: objUsredt?.params?.keywrd || [],
      }));
    }
  }, [prmAllusr.usredt]);

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
      succssSet(rspscs);
      rplprm(["update", "usredt"], [String(Math.random()), ""]);
      setTimeout(() => {
        succssSet("Create user");
        rspnseSet((prev) => ({ ...prev, finals_errrsp: "" }));
      }, 2000);
    } else if (rspnse.finals_errrsp != "") {
      succssSet(rspnse.finals_errrsp);
      setTimeout(() => {
        succssSet(prmAllusr.usredt ? "Update user" : "Create user");
        rspnseSet((prev) => ({ ...prev, finals_errrsp: "" }));
      }, 1500);
    }
  }, [rspnse.finals_errrsp]);

  // Reset function
  const resetx = () => {
    rplprm(["usredt"], "");
    rspnseSet((prev) => ({
      ...prev,
      stfnme_dfault: "",
      stfeml_dfault: "",
      usrnme_dfault: "",
      psswrd_dfault: "",
      confrm_dfault: "",
      access_dfault: [],
      keywrd_dfault: [],
    }));
  };

  return (
    <form className="flex h-24 min-h-fit w-full flex-col gap-6 py-3" action={formacSet}>
      <input type="hidden" name="action" value={prmAllusr.usredt ? "update" : "regist"} />
      <div className="flex w-full flex-col gap-3 md:flex-row">
        <div className="flexctr h-full w-full flex-col gap-3 md:gap-6">
          <div className="relative flex w-full flex-col items-start justify-center gap-1.5">
            <label className="font-medium" htmlFor="stfnme">
              Staff name
            </label>
            <input
              className="w-full rounded-md px-3 py-1.5 ring ring-gray-300"
              type="text"
              name="stfnme"
              id="stfnme"
              defaultValue={rspnse.stfnme_dfault}
              onChange={() => rspnseSet((prev) => ({ ...prev, stfnme_errrsp: "" }))}
            />
            <span className={`absolute right-2 -bottom-4 text-[10px] text-red-600`}>
              {rspnse.stfnme_errrsp}
            </span>
          </div>
          <div className="relative flex w-full flex-col items-start justify-center gap-1.5">
            <label className="font-medium" htmlFor="stfeml">
              Staff Email
            </label>
            <input
              className="w-full rounded-md px-3 py-1.5 ring ring-gray-300"
              type="email"
              name="stfeml"
              id="stfeml"
              defaultValue={rspnse.stfeml_dfault}
              onChange={() => rspnseSet((prev) => ({ ...prev, stfeml_errrsp: "" }))}
            />
            <span className={`absolute right-2 -bottom-4 text-[10px] text-red-600`}>
              {rspnse.stfeml_errrsp}
            </span>
          </div>
          <div className="relative flex w-full flex-col items-start justify-center gap-1.5">
            <label className="font-medium" htmlFor="usrnme">
              Username
            </label>
            <input
              className="w-full rounded-md px-3 py-1.5 ring ring-gray-300"
              type="text"
              name="usrnme"
              id="usrnme"
              defaultValue={rspnse.usrnme_dfault}
              onChange={() => rspnseSet((prev) => ({ ...prev, usrnme_errrsp: "" }))}
            />
            <span className={`absolute right-2 -bottom-4 text-[10px] text-red-600`}>
              {rspnse.usrnme_errrsp}
            </span>
          </div>
          <div className="relative flex w-full flex-col items-start justify-center gap-1.5">
            <label className="font-medium" htmlFor="psswrd">
              Password
            </label>
            <input
              className={`w-full rounded-md px-3 py-1.5 ring ring-gray-300 ${!(succss == rspdef) && "pointer-events-none bg-gray-200"}`}
              type="password"
              name="psswrd"
              id="psswrd"
              value={rspnse.psswrd_dfault}
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
              Confirm Password
            </label>
            <input
              className={`w-full rounded-md px-3 py-1.5 ring ring-gray-300 ${!(succss == rspdef) && "pointer-events-none bg-gray-200"}`}
              type="password"
              name="confrm"
              id="confrm"
              value={rspnse.confrm_dfault}
              onChange={(e) => rspnseSet((prev) => ({ ...prev, confrm_dfault: e.target.value }))}
            />
            <span className={`absolute right-2 -bottom-4 text-[10px] text-red-600`}>
              {rspnse.confrm_errrsp}
            </span>
          </div>
          <input name="isited" id="isited" defaultValue={succss} hidden />
        </div>
        <div className="flex h-full w-full flex-col items-start gap-6 justify-self-center">
          <div className="flex w-full flex-col items-start gap-3 justify-self-center">
            <div className="font-medium">Select Access</div>
            <div className="flexstr flex-wrap gap-1.5">
              {applst.map((val, idx) => (
                <div key={idx}>
                  <input
                    className="peer"
                    hidden
                    type="checkbox"
                    name="access"
                    id={val.prmkey}
                    value={val.prmkey}
                    onChange={(e) => {
                      if (e.target.checked)
                        rspnseSet((prev) => ({
                          ...prev,
                          access_dfault: [...prev.access_dfault, val.prmkey],
                        }));
                      else
                        rspnseSet((prev) => ({
                          ...prev,
                          access_dfault: prev.access_dfault.filter((v) => v !== val.prmkey),
                        }));
                    }}
                    checked={rspnse.access_dfault.includes(val.prmkey) || val.prmkey == "global"}
                  />
                  <label
                    className="flexctr w-fit cursor-pointer gap-0.5 rounded-md px-1.5 py-1 ring-2 ring-gray-200 select-none peer-checked:bg-cyan-600 peer-checked:text-white peer-checked:[&_div]:w-4 peer-checked:[&_div]:opacity-100"
                    htmlFor={val.prmkey}
                  >
                    <span>{val.prmkey}</span>
                    <div className="w-0 opacity-0 duration-300">
                      <UixGlobalIconvcCancel bold={4} color="#fff" size={1.3} />
                    </div>
                  </label>
                </div>
              ))}
            </div>
          </div>
          <div className="relative flex w-full flex-col items-start justify-center gap-1.5">
            <label className="font-medium" htmlFor="usrnme">
              Add Keyword
            </label>
            <div className="flex w-full gap-x-3 pb-3">
              <input
                className="w-full rounded-md px-3 py-1.5 ring ring-gray-300"
                type="text"
                value={keywrd}
                onChange={(e) => keywrdSet(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === "Enter") {
                    e.preventDefault();
                    keywrdAdd();
                  }
                }}
              />
              <div className="flexctr btnsbm w-16" onClick={() => keywrdAdd()}>
                add
              </div>
            </div>
            <div className="flexstr flex-wrap gap-1.5">
              {rspnse.keywrd_dfault.map((val, idx) => (
                <div key={idx}>
                  <input className="peer" type="hidden" name="keywrd" id={idx + val} value={val} />
                  <div
                    className="flexctr w-fit cursor-pointer gap-0.5 rounded-md bg-cyan-600 px-1.5 py-1 text-white ring-2 ring-gray-200 select-none"
                    onClick={() => keywrdRmv(val)}
                  >
                    <span>{val.substring(0, 6)}</span>
                    <div>
                      <UixGlobalIconvcCancel bold={4} color="#fff" size={1.3} />
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
      <div className="flexctr gap-3">
        <button className="btnsbm pointer-events-none w-0 px-0 opacity-0 select-none">
          xawdwad
        </button>
        <button
          className={`w-full py-1.5 text-center ${succss == rspscs ? "btnscs" : succss == rspdef ? "btnsbm" : "btnwrn"}`}
          type="submit"
          disabled={pndingIst}
        >
          {pndingIst ? "Loading..." : succss}
        </button>
        <div
          className={`${prmAllusr.usredt ? "w-20" : "w-0 opacity-0"} btncxl py-1.5 text-center whitespace-nowrap duration-300`}
          onClick={() => resetx()}
        >
          Reset
        </div>
      </div>
    </form>
  );
}
