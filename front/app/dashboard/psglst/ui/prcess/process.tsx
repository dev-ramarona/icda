"use client";

import { useEffect, useState } from "react";
import { MdlPsglstErrlogDtbase } from "../../model/params";
import { ApiPsglstPrcessManual } from "../../api/prcess";
import { mdlGlobalAllusrCookie, MdlGlobalStatusPrcess } from "../../../global/model/params";
import { FncGlobalQuerysEdlink, FncGlobalParamsHminfr } from "../../../global/function/querys";
import { ApiGlobalStatusPrcess } from "../../../global/api/status";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import { FncGlobalFormatDatefm } from "../../../global/function/format";
import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";


export default function UixPsglstPrcessManual({ cookie, update, status }:
  { cookie: mdlGlobalAllusrCookie; update: string; status: MdlGlobalStatusPrcess }) {

  // Get status first
  const rplprm = FncGlobalQuerysEdlink();
  const hminfr = FncGlobalParamsHminfr(4)
  const dfault: MdlPsglstErrlogDtbase = {
    prmkey: "", erstat: "", erpart: "",
    ersrce: "", erdtil: "", erdvsn: "",
    erignr: "", dateup: 0, timeup: 0,
    datefl: 0, airlfl: "", depart: "",
    flnbfl: "", Paxdif: "", flstat: "",
    flhour: 0, routfl: "", updtby: "", worker: 1,
  }
  const nwhour = (Number(new Date().getHours().toString().padStart(2, '0')))
  const [params, paramsSet] = useState<MdlPsglstErrlogDtbase>(dfault)
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
  }

  // Process function
  const prcess = async (params: MdlPsglstErrlogDtbase) => {
    rplprm(["update_global"], String(Math.random()));
    statfnSet("Wait");
    const nowprm = { ...params };
    if ((cookie.keywrd && (cookie.keywrd).includes("psglst")) || nowprm.worker == 1)
      if (status.sbrapi == 0) {
        if (params.flnbfl == "") {
          nowprm.worker = 3;
          if (params.depart == "") {
            nowprm.worker = 5;
            if (params.airlfl == "")
              nowprm.worker = 8;
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
  }

  // Monitor process status
  useEffect(() => {
    if (status.sbrapi == 0) statfnSet("");
    const gtstat = async () => {
      if (status.sbrapi != 0) {
        const intrvl = setInterval(async () => {
          console.log("action interval");
          const instat = await ApiGlobalStatusPrcess();
          if (instat.sbrapi == 0) {
            statfnSet("");
            rplprm(["update_global"], String(Math.random()));
            clearInterval(intrvl);
          } else statfnSet(`${instat.sbrapi}%`);
        }, 2000);
      }
    };
    gtstat();
  }, [update]);


  return (
    <div className="w-full h-24 min-h-fit py-3 flexctr relative">
      <div className={`${statfn != "" ? "w-16 h-10 translate-y-0" : "w-0 h-0 opacity-0 -translate-y-10"} 
      z-10 absolute bg-white ring-2 ring-sky-300 px-5 py-2 rounded-xl flexctr duration-300`}>
        <div>Wait</div>
        <div className="animate-spin"><UixGlobalIconvcRfresh bold={2} color="black" size={1} /></div>
      </div>
      <div className={`afull flexstr flex-wrap gap-y-3 ${statfn != "" ? "animate-pulse select-none" : ""} duration-300`}>
        <div className="w-full md:w-28 h-10 flexctr relative">
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
        <div className="w-full md:w-28 h-10 flexctr relative">
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
        <div className="w-full md:w-28 h-10 flexctr relative">
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
      <div className={`w-full flexend flex-wrap gap-3 px-3 ${statfn != "" ? "animate-pulse select-none pointer-events-none" : ""} duration-300`}>
        {hminfr.map((val, idx) => (
          <div className="w-full md:w-40 h-12 flexctr relative" key={idx}>
            <button
              className={`afull flexctr 
                ${nwhour > 11 && idx == hminfr.length - 1 ? "btnoff select-none pointer-events-none" : "btnsbm"} 
                ${statfn.includes("admin") ? "shkeit btncxl" : ""}`}
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
