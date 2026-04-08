"use client";
import { useEffect, useState } from "react";
import { MdlPsglstPsgdtlSrcprm } from "../../model/params";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatDfault } from "../../../global/function/format";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import UixGlobalWraperSearch from "../../../public/ui/search/wraper";

export default function UixPsglstDetailSearch({
  prmPsgdtl,
  datefl,
  fmtdef,
}: {
  prmPsgdtl: MdlPsglstPsgdtlSrcprm;
  datefl: string[];
  fmtdef: boolean;
}) {
  const [params, paramsSet] = useState<MdlPsglstPsgdtlSrcprm>({
    update_global: prmPsgdtl.update_global || "",
    mnthfl_psgdtl: prmPsgdtl.mnthfl_psgdtl || "",
    datefl_psgdtl: prmPsgdtl.datefl_psgdtl || "",
    airlfl_psgdtl: prmPsgdtl.airlfl_psgdtl || "",
    flnbfl_psgdtl: prmPsgdtl.flnbfl_psgdtl || "",
    depart_psgdtl: prmPsgdtl.depart_psgdtl || "",
    routfl_psgdtl: prmPsgdtl.routfl_psgdtl || "",
    pnrcde_psgdtl: prmPsgdtl.pnrcde_psgdtl || "",
    tktnfl_psgdtl: prmPsgdtl.tktnfl_psgdtl || "",
    isitfl_psgdtl: prmPsgdtl.isitfl_psgdtl || "",
    isittx_psgdtl: prmPsgdtl.isittx_psgdtl || "",
    isitir_psgdtl: prmPsgdtl.isitir_psgdtl || "",
    nclear_psgdtl: prmPsgdtl.nclear_psgdtl || "",
    format_psgdtl: prmPsgdtl.format_psgdtl || "",
    pagenw_psgdtl: prmPsgdtl.pagenw_psgdtl || 1,
    limitp_psgdtl: prmPsgdtl.limitp_psgdtl || 15,
  });

  // Monitor change
  const [chnged, chngedSet] = useState<boolean>(false);
  useEffect(() => {
    chngedSet(false);
    paramsSet({
      update_global: prmPsgdtl.update_global || "",
      mnthfl_psgdtl: prmPsgdtl.mnthfl_psgdtl || "",
      datefl_psgdtl: prmPsgdtl.datefl_psgdtl || "",
      airlfl_psgdtl: prmPsgdtl.airlfl_psgdtl || "",
      flnbfl_psgdtl: prmPsgdtl.flnbfl_psgdtl || "",
      depart_psgdtl: prmPsgdtl.depart_psgdtl || "",
      routfl_psgdtl: prmPsgdtl.routfl_psgdtl || "",
      pnrcde_psgdtl: prmPsgdtl.pnrcde_psgdtl || "",
      tktnfl_psgdtl: prmPsgdtl.tktnfl_psgdtl || "",
      isitfl_psgdtl: prmPsgdtl.isitfl_psgdtl || "",
      isittx_psgdtl: prmPsgdtl.isittx_psgdtl || "",
      isitir_psgdtl: prmPsgdtl.isitir_psgdtl || "",
      nclear_psgdtl: prmPsgdtl.nclear_psgdtl || "",
      format_psgdtl: prmPsgdtl.format_psgdtl || "",
      pagenw_psgdtl: prmPsgdtl.pagenw_psgdtl || 1,
      limitp_psgdtl: prmPsgdtl.limitp_psgdtl || 15,
    });
  }, [prmPsgdtl]);

  // Replace params
  const rplprm = FncGlobalQuerysEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    chngedSet(true);
    const namefl = e.currentTarget.name;
    const valuef = FncGlobalFormatDfault(namefl, e.currentTarget.value);
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
    rplprm([namefl, "pagenw_psgdtl"], [valuef as string, ""]);
  };

  // Reset function
  const resetx = () => {
    chngedSet(true);
    rplprm(
      [
        "prmkey_psgdtl",
        "mnthfl_psgdtl",
        "datefl_psgdtl",
        "airlfl_psgdtl",
        "flnbfl_psgdtl",
        "depart_psgdtl",
        "routfl_psgdtl",
        "pnrcde_psgdtl",
        "tktnfl_psgdtl",
        "isitfl_psgdtl",
        "isittx_psgdtl",
        "isitir_psgdtl",
        "nclear_psgdtl",
        "format_psgdtl",
        "pagenw_psgdtl",
      ],
      "",
    );
  };

  // Update file
  const [filedt, filedtSet] = useState<FileList | null>(null);
  const [filenm, filenmSet] = useState<string>("");
  const filefn = (e: React.ChangeEvent<HTMLInputElement>) => {
    filedtSet(e.target.files);
    if (e.target.files.length > 1) {
      filenmSet(`${e.target.files.length} files selected`);
    } else if (e.target.files.length == 1) {
      filenmSet(e.target.files[0].name);
    } else filenmSet("");
  };
  return (
    <UixGlobalWraperSearch
      chnged={chnged}
      lblupl={fmtdef ? "Upload error" : ""}
      downld={{
        lnk: `/psglst/psgdtl/downld`,
        prm: params,
      }}
      upload={{
        lnk: `/psglst/psgdtl/upload`,
        prm: filedt,
      }}
      resetx={resetx}
      updtfl={filefn}
      namefl={filenm}
    >
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"select"}
          length={["ALL", "MNFERR"]}
          queryx={"nclear_psgdtl"}
          params={params.nclear_psgdtl}
          plchdr="Not Clear"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"select"}
          length={["MNFERR", "EBTFMT", "TKTFMT", "ALL"]}
          queryx={"format_psgdtl"}
          params={params.format_psgdtl}
          plchdr="Format data"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"select"}
          length={["F", "N"]}
          queryx={"isitfl_psgdtl"}
          params={params.isitfl_psgdtl}
          plchdr="Flown Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"select"}
          length={["TX"]}
          queryx={"isittx_psgdtl"}
          params={params.isittx_psgdtl}
          plchdr="Transit Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"select"}
          length={["IR"]}
          queryx={"isitir_psgdtl"}
          params={params.isitir_psgdtl}
          plchdr="Irreg Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"month"}
          length={undefined}
          queryx={"mnthfl_psgdtl"}
          params={params.mnthfl_psgdtl}
          plchdr="Flight Month"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"date"}
          length={datefl}
          queryx={"datefl_psgdtl"}
          params={params.datefl_psgdtl}
          plchdr="Flight Date"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"airlfl_psgdtl"}
          params={params.airlfl_psgdtl}
          plchdr="Airline"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"flnbfl_psgdtl"}
          params={params.flnbfl_psgdtl}
          plchdr="Flight Number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"depart_psgdtl"}
          params={params.depart_psgdtl}
          plchdr="Departure"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"routfl_psgdtl"}
          params={params.routfl_psgdtl}
          plchdr="Route"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"pnrcde_psgdtl"}
          params={params.pnrcde_psgdtl}
          plchdr="PNR Code"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"tktnfl_psgdtl"}
          params={params.tktnfl_psgdtl}
          plchdr="Ticket Number"
          repprm={repprm}
          labelx=""
        />
      </div>
    </UixGlobalWraperSearch>
  );
}
