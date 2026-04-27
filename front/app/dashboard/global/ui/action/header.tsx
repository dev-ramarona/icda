"use client";

import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";
import { MdlAllusrApplstParams } from "../../../allusr/model/params";

export default function UixGlobalHeaderClient({
  applst,
}: {
  applst: MdlAllusrApplstParams[];
}) {
  const pthnme = usePathname(); // misal: "/opclss"
  const [nowpth, nowpthSet] = useState("");
  const [dtlpth, dtlpthSet] = useState("Wellcome");
  useEffect(() => {
    const sgment = pthnme.split("/").filter(Boolean).pop();
    nowpthSet(sgment || "");
    applst.forEach((app) => {
      if (app.prmkey == sgment) {
        dtlpthSet(app.detail);
      }
    });
  }, [pthnme, applst]);
  if (nowpth == "" || nowpth == "global") return;
  return (
    <>
      <div className="font-bold text-2xl md:text-4xl text-slate-600 tracking-wide">
        {dtlpth.toUpperCase()}
      </div>
      <div className="text-lg text-slate-800 px-1.5">Page</div>
    </>
  );
}
