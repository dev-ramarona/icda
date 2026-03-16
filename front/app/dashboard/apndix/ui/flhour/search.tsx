"use client";
import { useEffect, useRef, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatDefault } from "../../../global/function/format";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import { MdlApndixSearchQueryx } from "../../model/parmas";


export default function UixApndixFlhourSearch({
  qryprm,
  datefl,
}: {
  qryprm: MdlApndixSearchQueryx;
  datefl: string[];

}) {
  const [params, paramsSet] = useState<MdlApndixSearchQueryx>({
    update: qryprm.update || "",
    pagedb: qryprm.pagedb || "",
    datefl: qryprm.datefl || "",
    airlfl: qryprm.airlfl || "",
    depart: qryprm.depart || "",
    flnbfl: qryprm.flnbfl || "",
    routfl: qryprm.routfl || "",
    clssfl: qryprm.clssfl || "",
    pagenw: Number(qryprm.pagenw) || 1,
    limitp: Number(qryprm.limitp) || 15,
  });

  // Monitor change
  const [chnged, chngedSet] = useState<boolean>(false);
  useEffect(() => {
    chngedSet(false);
    paramsSet({
      update: qryprm.update || "",
      pagedb: qryprm.pagedb || "",
      datefl: qryprm.datefl || "",
      airlfl: qryprm.airlfl || "",
      depart: qryprm.depart || "",
      flnbfl: qryprm.flnbfl || "",
      routfl: qryprm.routfl || "",
      clssfl: qryprm.clssfl || "",
      pagenw: Number(qryprm.pagenw) || 1,
      limitp: Number(qryprm.limitp) || 15,
    });
  }, [qryprm]);

  // Replace params
  const timerf = useRef<NodeJS.Timeout | null>(null)
  const rplprm = FncGlobalQuerysEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    const namefl = e.currentTarget.id;
    let valuef: string | number = e.currentTarget.value;
    valuef = FncGlobalFormatDefault(namefl, valuef);
    paramsSet({ ...params, [namefl]: valuef, });
    if (timerf.current) clearTimeout(timerf.current)
    timerf.current = setTimeout(async () => {
      chngedSet(true);
      rplprm([namefl, "pagenw_psgdtl"], [valuef as string, ""]);
    }, 1000)
  };

  // Reset function
  const resetx = () => {
    chngedSet(true);
    rplprm(
      [
        "update",
        "pagedb",
        "datefl",
        "airlfl",
        "depart",
        "flnbfl",
        "routfl",
        "clssfl",
        "pagenw",
      ],
      ""
    );
  };
  return (
    <div className="w-full h-24 min-h-fit py-3 flexctr relative">
      <div className={`${chnged ? "w-16 h-10 translate-y-0" : "w-0 h-0 opacity-0 -translate-y-10"} z-10 absolute bg-white ring-2 ring-sky-300 px-5 py-2 rounded-xl flexctr duration-300`}>
        <div>Wait</div>
        <div className="animate-spin"><UixGlobalIconvcRfresh bold={2} color="black" size={1} /></div>
      </div>
      <div className={`afull flexstr flex-wrap gap-y-3 ${chnged ? "animate-pulse select-none" : ""} duration-300`}>
        <div className="w-1/2 md:w-28 h-11 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"date"}
            length={undefined}
            queryx={"datefl"}
            params={params.datefl}
            plchdr="Date flown"
            repprm={repprm}
            labelx={'a:"ALL"(All Data)|SPT:"SLSRPT"(Sales Report)|MNF:"MNFEST"(Manifest)'}
          />
        </div>
        <div className="w-1/2 md:w-28 h-11 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"airlfl"}
            params={params.airlfl}
            plchdr="Airline"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-11 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={datefl}
            queryx={"depart"}
            params={params.depart}
            plchdr="Departure"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-11 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"flnbfl"}
            params={params.flnbfl}
            plchdr="Flight number"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-11 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"routfl"}
            params={params.routfl}
            plchdr="Route"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-11 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"clssfl"}
            params={params.clssfl}
            plchdr="Class"
            repprm={repprm}
            labelx=""
          />
        </div>
      </div>
      <div className={`w-1/3 flexend flex-wrap gap-3 px-3 ${chnged ? "animate-pulse select-none" : ""} duration-300`}>
        <form className="w-full md:w-28 h-10 flexctr relative"
          method="POST"
          action={`${process.env.NEXT_PUBLIC_URL_SERVER}/psglst/psgdtl/getall/downld`}>
          <input type="hidden" name="data" value={JSON.stringify(params)} />
          <button type="submit" className="afull btnsbm flexctr">
            Download
          </button>
        </form>
        <div className="w-full md:w-28 h-10 flexctr relative">
          <div className="afull btnwrn flexctr" onClick={() => resetx()}>
            Reset
          </div>
        </div>
      </div>
    </div>
  );
}
