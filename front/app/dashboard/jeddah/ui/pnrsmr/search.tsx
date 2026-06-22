"use client";
import { useEffect, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatDfault } from "../../../global/function/format";
import UixGlobalInputxFormdt from "../../../global/ui/action/inputx";
import UixGlobalWraperSearch from "../../../global/ui/search/wraper";
import { ApiAllusrStatusPrcess } from "../../../allusr/api/status";
import { MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { MdlJeddahGlobalSrcprm } from "../../model/params";

export default function UixJeddahPnrsmrSearch({
  prmPnrsmr,
  status,
  update,
}: {
  prmPnrsmr: MdlJeddahGlobalSrcprm;
  status: MdlAllusrStatusPrcess;
  update: string;
}) {
  const [params, paramsSet] = useState<MdlJeddahGlobalSrcprm>({
    update_global: prmPnrsmr.update_global || "",
    airlfl_jeddah: prmPnrsmr.airlfl_jeddah || "",
    flnbfl_jeddah: prmPnrsmr.flnbfl_jeddah || "",
    depart_jeddah: prmPnrsmr.depart_jeddah || "",
    routfl_jeddah: prmPnrsmr.routfl_jeddah || "",
    pnrcde_jeddah: prmPnrsmr.pnrcde_jeddah || "",
    pagenw_jeddah: prmPnrsmr.pagenw_jeddah || 1,
    limitp_jeddah: prmPnrsmr.limitp_jeddah || 15,
  });

  // Monitor change
  const [chnged, chngedSet] = useState<boolean>(false);
  useEffect(() => {
    chngedSet(false);
    paramsSet({
      update_global: prmPnrsmr.update_global || "",
      airlfl_jeddah: prmPnrsmr.airlfl_jeddah || "",
      flnbfl_jeddah: prmPnrsmr.flnbfl_jeddah || "",
      depart_jeddah: prmPnrsmr.depart_jeddah || "",
      routfl_jeddah: prmPnrsmr.routfl_jeddah || "",
      pnrcde_jeddah: prmPnrsmr.pnrcde_jeddah || "",
      pagenw_jeddah: prmPnrsmr.pagenw_jeddah || 1,
      limitp_jeddah: prmPnrsmr.limitp_jeddah || 15,
    });
  }, [prmPnrsmr]);

  // Replace params
  const rplprm = FncGlobalQuerysEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    chngedSet(true);
    const namefl = e.currentTarget.name;
    const valuef = FncGlobalFormatDfault(namefl, e.currentTarget.value);
    const random = String(Math.random());
    paramsSet({ ...params, [namefl]: valuef });
    if (namefl.includes("date")) {
      rplprm([namefl, "pagenw_pnrsmr", "update_global"], [valuef as string, "", random]);
    } else rplprm([namefl, "pagenw_pnrsmr"], [valuef as string, ""]);
  };

  // Reset function
  const resetx = () => {
    chngedSet(true);
    rplprm(
      [
        "update_global",
        "airlfl_jeddah",
        "flnbfl_jeddah",
        "depart_jeddah",
        "routfl_jeddah",
        "pnrcde_jeddah",
        "pagenw_jeddah",
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

  // Monitor process status
  const [statfn, statfnSet] = useState(100);
  useEffect(() => {
    if (status.sbrapi == 0) statfnSet(0);
    const gtstat = async () => {
      if (status.sbrapi != 0) {
        const intrvl = setInterval(async () => {
          const instat = await ApiAllusrStatusPrcess();
          if (instat.sbrapi == 0) {
            statfnSet(0);
            rplprm(["update_global"], String(Math.random()));
            clearInterval(intrvl);
          } else statfnSet(instat.sbrapi);
        }, 5000);
      }
    };
    gtstat();
  }, [update]);

  return (
    <UixGlobalWraperSearch
      chnged={chnged}
      lblupl="Upload File Join"
      downld={{ lnk: `/jeddah/pnrsmr/downld`, prm: params }}
      upload={undefined}
      resetx={resetx}
      updtfl={filefn}
      namefl={filenm}
    >
      <div className="flexctr relative h-10 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"airlfl_jeddah"}
          params={params.airlfl_jeddah}
          plchdr="Airline"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-10 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"flnbfl_jeddah"}
          params={params.flnbfl_jeddah}
          plchdr="Flight number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-10 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"depart_jeddah"}
          params={params.depart_jeddah}
          plchdr="Departure"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-10 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"routfl_jeddah"}
          params={params.routfl_jeddah}
          plchdr="Route"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-10 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"pnrcde_jeddah"}
          params={params.pnrcde_jeddah}
          plchdr="PNR Code"
          repprm={repprm}
          labelx=""
        />
      </div>
    </UixGlobalWraperSearch>
  );
}
