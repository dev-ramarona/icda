"use client";
import { useEffect, useState } from "react";
import { MdlSlsflwPsgsmrSrcprm } from "../../model/params";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatDfault } from "../../../global/function/format";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import UixGlobalWraperSearch from "../../../public/ui/search/wraper";

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
    isitjn_psgsmr: prmPsgsmr.isitjn_psgsmr || "",
    depart_psgsmr: prmPsgsmr.depart_psgsmr || "",
    routfl_psgsmr: prmPsgsmr.routfl_psgsmr || "",
    keywrd_psgsmr: prmPsgsmr.keywrd_psgsmr || "",
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
      isitjn_psgsmr: prmPsgsmr.isitjn_psgsmr || "",
      depart_psgsmr: prmPsgsmr.depart_psgsmr || "",
      routfl_psgsmr: prmPsgsmr.routfl_psgsmr || "",
      keywrd_psgsmr: prmPsgsmr.keywrd_psgsmr || "",
      pagenw_psgsmr: prmPsgsmr.pagenw_psgsmr || 1,
      limitp_psgsmr: prmPsgsmr.limitp_psgsmr || 15,
    });
  }, [prmPsgsmr]);

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
        "mnthfl_psgsmr",
        "datefl_psgsmr",
        "airlfl_psgsmr",
        "flnbfl_psgsmr",
        "isitjn_psgsmr",
        "depart_psgsmr",
        "routfl_psgsmr",
        "keywrd_psgsmr",
        "pagenw_psgsmr",
        "limitp_psgsmr",
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
      lblupl="Upload File Join"
      downld={{
        lnk: `/psglst/psgsmr/downld`,
        prm: params,
      }}
      upload={{
        lnk: `/apndix/fljoin/upload`,
        prm: filedt,
      }}
      resetx={resetx}
      updtfl={filefn}
      namefl={filenm}
    >
      <div className="flexctr relative h-10 w-1/2 md:w-28">
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
      <div className="flexctr relative h-10 w-1/2 md:w-28">
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
      <div className="flexctr relative h-10 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"select"}
          length={["Combined", "Separated"]}
          queryx={"isitjn_psgsmr"}
          params={params.isitjn_psgsmr}
          plchdr="Combine Join"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-10 w-1/2 md:w-28">
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
      <div className="flexctr relative h-10 w-1/2 md:w-28">
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
      <div className="flexctr relative h-10 w-1/2 md:w-28">
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
      <div className="flexctr relative h-10 w-1/2 md:w-28">
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
    </UixGlobalWraperSearch>
  );
}
