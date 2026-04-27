"use client";

import { useEffect, useState } from "react";
import { MdlPsglstErrlogDtbase } from "../../model/params";
import { ApiPsglstPrcessManual } from "../../api/prcess";
import { FncGlobalQuerysEdlink, FncGlobalParamsHminfr } from "../../../global/function/querys";
import UixGlobalInputxFormdt from "../../../global/ui/action/inputx";
import { FncGlobalFormatDatefm } from "../../../global/function/format";
import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { ApiAllusrStatusPrcess } from "../../../allusr/api/status";

export default function UixPsglstPrcessManual({
  cookie,
  update,
  status,
}: {
  cookie: mdlAllusrCookieObjson;
  update: string;
  status: MdlAllusrStatusPrcess;
}) {
  // Get status first
  const rplprm = FncGlobalQuerysEdlink();
  const hminfr = FncGlobalParamsHminfr(4);
  const dfault: MdlPsglstErrlogDtbase = {
    prmkey: "",
    erstat: "",
    erpart: "",
    ersrce: "",
    erdtil: "",
    erdvsn: "",
    erignr: "",
    dateup: 0,
    timeup: 0,
    datefl: 0,
    airlfl: "",
    depart: "",
    flnbfl: "",
    Paxdif: "",
    flstat: "",
    flhour: 0,
    routfl: "",
    updtby: "",
    worker: 1,
  };
  const nwhour = Number(new Date().getHours().toString().padStart(2, "0"));
  const [params, paramsSet] = useState<MdlPsglstErrlogDtbase>(dfault);
  const [statfn, statfnSet] = useState("Wait");

  // Edit parameter
  const onchge = (e: React.ChangeEvent<HTMLInputElement>) => {
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value.toUpperCase();
    if (namefl == "flnbfl") valuef = valuef.replace(/[^0-9]/g, "");
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
  };

  // Process function
  const prcess = async (params: MdlPsglstErrlogDtbase) => {
    rplprm(["update_global"], String(Math.random()));
    statfnSet("Wait");
    const nowprm = { ...params };
    if ((cookie.keywrd && cookie.keywrd.includes("psglst")) || nowprm.worker == 1)
      if (status.sbrapi == 0) {
        if (params.flnbfl == "") {
          nowprm.worker = 3;
          if (params.depart == "") {
            nowprm.worker = 5;
            if (params.airlfl == "") nowprm.worker = 8;
          }
        }

        // Set interval to check status
        const rsp = ApiPsglstPrcessManual(nowprm);
        setTimeout(() => {
          rplprm(["update_global"], String(Math.random()));
        }, 1000);
        statfnSet(await rsp);
        setTimeout(() => statfnSet(""), 2000);
      }
  };

  // Monitor process status
  useEffect(() => {
    if (status.sbrapi == 0) statfnSet("");
    const gtstat = async () => {
      if (status.sbrapi != 0) {
        const intrvl = setInterval(async () => {
          const instat = await ApiAllusrStatusPrcess();
          if (instat.sbrapi == 0) {
            statfnSet("");
            rplprm(["update_global"], String(Math.random()));
            clearInterval(intrvl);
          } else statfnSet(`${instat.sbrapi}%`);
        }, 5000);
      }
    };
    gtstat();
  }, [update]);

  return (
    <div className="flexctr relative h-24 min-h-fit w-full py-3">
      <div
        className={`${statfn != "" ? "h-10 w-16 translate-y-0" : "h-0 w-0 -translate-y-10 opacity-0"} flexctr absolute z-10 rounded-xl bg-white px-5 py-2 ring-2 ring-sky-300 duration-300`}
      >
        <div>Wait</div>
        <div className="animate-spin">
          <UixGlobalIconvcRfresh bold={2} color="black" size={1} />
        </div>
      </div>
      <div
        className={`afull flexstr flex-wrap gap-y-3 ${statfn != "" ? "animate-pulse select-none" : ""} duration-300`}
      >
        <div className="flexctr relative h-10 w-full md:w-28">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={2}
            queryx={"airlfl"}
            params={params.airlfl}
            plchdr="Airline"
            repprm={onchge}
            labelx=""
          />
        </div>
        <div className="flexctr relative h-10 w-full md:w-28">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={3}
            queryx={"depart"}
            params={params.depart}
            plchdr="Departure"
            repprm={onchge}
            labelx=""
          />
        </div>
        <div className="flexctr relative h-10 w-full md:w-28">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={4}
            queryx={"flnbfl"}
            params={params.flnbfl}
            plchdr="Flight Number"
            repprm={onchge}
            labelx=""
          />
        </div>
      </div>
      <div
        className={`flexend w-full flex-wrap gap-3 px-3 ${statfn != "" ? "pointer-events-none animate-pulse select-none" : ""} duration-300`}
      >
        {hminfr.map((val, idx) => (
          <div className="flexctr relative h-12 w-full md:w-40" key={idx}>
            <button
              className={`afull flexctr ${nwhour > 11 && idx == hminfr.length - 1 ? "btnoff pointer-events-none select-none" : "btnsbm"} ${statfn.includes("admin") ? "shkeit btncxl" : ""}`}
              onClick={() => prcess({ ...params, datefl: val })}
            >
              {statfn == "" ? `Process Manual ${FncGlobalFormatDatefm(String(val))}` : statfn}
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}
