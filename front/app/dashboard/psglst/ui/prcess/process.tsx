"use client";

import { useEffect, useState } from "react";
import { MdlPsglstErrlogDtbase } from "../../model/params";
import { ApiPsglstPrcessManual } from "../../api/prcess";
import { mdlGlobalAllusrCookie } from "../../../global/model/params";
import { FncGlobalQuerysEdlink, FncGlobalParamsHminfr } from "../../../global/function/querys";
import { ApiGlobalStatusIntrvl, ApiGlobalStatusPrcess } from "../../../global/api/status";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import { FncGlobalFormatDatefm } from "../../../global/function/format";
import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";


export default function UixPsglstPrcessManual({ cookie, update }:
  { cookie: mdlGlobalAllusrCookie; update: string; }) {

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
  const [statfn, statfnSet] = useState("Done");

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

  // Hit the database and get interval status
  const prcess = async (params: MdlPsglstErrlogDtbase) => {
    const status = await ApiGlobalStatusPrcess();
    const nowParams = { ...params };
    if (status.sbrapi == 0) {
      if (params.flnbfl == "") {
        nowParams.worker = 3;
        if (params.depart == "") {
          nowParams.worker = 5;
          if (params.airlfl == "")
            nowParams.worker = 8;
        }
      }

      // Cek is admin or not
      if ((cookie.keywrd && (cookie.keywrd).includes("psglst")) || nowParams.worker == 1) {
        statfnSet("Wait");
        rplprm(["update_global"], String(Math.random()));
        ApiPsglstPrcessManual(nowParams);
        await ApiGlobalStatusIntrvl(statfnSet, "sbrapi");
      } else {
        statfnSet("Only admin can process ALL");
        return setTimeout(() => statfnSet("Done"), 2000);
      }
    } else statfnSet(`Wait ${status.sbrapi}%`);
  };

  // Monitor process status
  useEffect(() => {
    const gtstat = async () => {
      const status = await ApiGlobalStatusPrcess();
      statfnSet(status.sbrapi == 0 ? "Done" : `Wait ${status.sbrapi}%`);
      if (status.sbrapi != 0) {
        await ApiGlobalStatusIntrvl(statfnSet, "sbrapi");
      } else statfnSet("Done");
    };
    gtstat();
  }, [update]);

  // refresh page
  useEffect(() => {
    if (statfn == "Process Done") setTimeout(() => {
      rplprm(["update_global"], String(Math.random()));
    }, 1000);
  }, [statfn, rplprm])

  return (
    <div className="w-full h-24 min-h-fit py-3 flexctr relative">
      <div className={`${statfn != "Done" ? "w-16 h-10 translate-y-0" : "w-0 h-0 opacity-0 -translate-y-10"} z-10 absolute bg-white ring-2 ring-sky-300 px-5 py-2 rounded-xl flexctr duration-300`}>
        <div>Wait</div>
        <div className="animate-spin"><UixGlobalIconvcRfresh bold={2} color="black" size={1} /></div>
      </div>
      <div className={`afull flexstr flex-wrap gap-y-3 ${statfn != "Done" ? "animate-pulse select-none" : ""} duration-300`}>
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
      <div className={`w-full flexend flex-wrap gap-3 px-3 ${statfn != "Done" ? "animate-pulse select-none pointer-events-none" : ""} duration-300`}>
        {hminfr.map((val, idx) => (
          <div className="w-full md:w-40 h-12 flexctr relative" key={idx}>
            <button
              className={`afull flexctr 
                ${nwhour > 11 && idx == hminfr.length - 1 ? "btnoff select-none pointer-events-none" : "btnsbm"} 
                ${statfn.includes("admin") ? "shkeit btncxl" : ""}`}
              onClick={() => prcess({ ...params, datefl: val })}
            >
              {statfn == "Done" ? `Process Manual ${FncGlobalFormatDatefm(String(val))}` : statfn}
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}
