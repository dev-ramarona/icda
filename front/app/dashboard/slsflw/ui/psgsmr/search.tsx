"use client";
import { useEffect, useState } from "react";
import { MdlSlsflwPsgsmrSrcprm } from "../../model/params";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatFilter, FncGlobalFormatRoutfl } from "../../../global/function/format";
import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";

export default function UixSlsflwPsgsmrSearch({
  prmPsgsmr,
  datefl,
}: {
  prmPsgsmr: MdlSlsflwPsgsmrSrcprm;
  datefl: string[];
}) {
  const [params, paramsSet] = useState<MdlSlsflwPsgsmrSrcprm>({
    update_global: prmPsgsmr.update_global || "",
    mnthfl_psgsmr: prmPsgsmr.mnthfl_psgsmr || "",
    datefl_psgsmr: prmPsgsmr.datefl_psgsmr || "",
    airlfl_psgsmr: prmPsgsmr.airlfl_psgsmr || "",
    flnbfl_psgsmr: prmPsgsmr.flnbfl_psgsmr || "",
    depart_psgsmr: prmPsgsmr.depart_psgsmr || "",
    routfl_psgsmr: prmPsgsmr.routfl_psgsmr || "",
    pagenw_psgsmr: prmPsgsmr.pagenw_psgsmr || 1,
    limitp_psgsmr: prmPsgsmr.limitp_psgsmr || 15,
  });

  // Monitor change
  const [chnged, chngedSet] = useState<boolean>(false);
  useEffect(() => {
    chngedSet(false);
    paramsSet({
      update_global: prmPsgsmr.update_global || "",
      mnthfl_psgsmr: prmPsgsmr.mnthfl_psgsmr || "",
      datefl_psgsmr: prmPsgsmr.datefl_psgsmr || "",
      airlfl_psgsmr: prmPsgsmr.airlfl_psgsmr || "",
      flnbfl_psgsmr: prmPsgsmr.flnbfl_psgsmr || "",
      depart_psgsmr: prmPsgsmr.depart_psgsmr || "",
      routfl_psgsmr: prmPsgsmr.routfl_psgsmr || "",
      pagenw_psgsmr: prmPsgsmr.pagenw_psgsmr || 1,
      limitp_psgsmr: prmPsgsmr.limitp_psgsmr || 15,
    });
  }, [prmPsgsmr]);

  // Replace params
  const rplprm = FncGlobalQuerysEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    chngedSet(true);
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    if (["isittx_psgsmr", "isitfl_psgsmr", "isitir_psgsmr"].includes(namefl))
      valuef = FncGlobalFormatFilter(valuef,
        [{ keywrd: "fl", output: "Flown" }, { keywrd: "no", output: "Not flown" }]);
    else if (namefl == "nclear_psgsmr") valuef = FncGlobalFormatFilter(valuef,
      [{ keywrd: "", output: "SLSRPT" }, { keywrd: "a", output: "ALL" }]);
    else if (namefl == "format_psgsmr") valuef = FncGlobalFormatFilter(valuef,
      [{ keywrd: "d", output: "DFAULT" },
      { keywrd: "e", output: "EBTFMT" },
      { keywrd: "t", output: "TKTFMT" }]);
    else if (["flnbfl_psgsmr", "tktnfl_psgsmr"].includes(namefl))
      valuef = valuef.replace(/[^0-9]/g, "");
    else if (namefl == "routfl_psgsmr") valuef = FncGlobalFormatRoutfl(valuef);
    else valuef = valuef.toUpperCase();
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
    rplprm([namefl, "pagenw_psgsmr"], [valuef, ""]);
  };

  // Reset function
  const resetx = () => {
    chngedSet(true);
    rplprm(
      [
        "prmkey_psgsmr",
        "mnthfl_psgsmr",
        "datefl_psgsmr",
        "airlfl_psgsmr",
        "flnbfl_psgsmr",
        "depart_psgsmr",
        "routfl_psgsmr",
        "pnrcde_psgsmr",
        "tktnfl_psgsmr",
        "isitfl_psgsmr",
        "isittx_psgsmr",
        "isitir_psgsmr",
        "nclear_psgsmr",
        "format_psgsmr",
        "pagenw_psgsmr",
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
        <div className="w-1/2 md:w-28 h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"month"}
            length={undefined}
            queryx={"mnthfl_psgsmr"}
            params={params.mnthfl_psgsmr}
            plchdr="Flight Month"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"date"}
            length={datefl}
            queryx={"datefl_psgsmr"}
            params={params.datefl_psgsmr}
            plchdr="Flight Date"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"airlfl_psgsmr"}
            params={params.airlfl_psgsmr}
            plchdr="Airline"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"flnbfl_psgsmr"}
            params={params.flnbfl_psgsmr}
            plchdr="Flight Number"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"depart_psgsmr"}
            params={params.depart_psgsmr}
            plchdr="Departure"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-28 h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"routfl_psgsmr"}
            params={params.routfl_psgsmr}
            plchdr="Route"
            repprm={repprm}
            labelx=""
          />
        </div>
      </div>
      <div className={`w-1/3 flexend flex-wrap gap-3 px-3 ${chnged ? "animate-pulse select-none" : ""} duration-300`}>
        <form className="w-full md:w-28 h-10 flexctr relative"
          method="POST"
          action={`${process.env.NEXT_PUBLIC_URL_AXIOSB}/psglst/psgsmr/getall/downld`}>
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
