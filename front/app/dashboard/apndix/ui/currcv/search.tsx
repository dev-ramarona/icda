"use client";
import { useEffect, useRef, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import { MdlApndixSearchQueryx } from "../../model/parmas";
import UixGlobalWraperSearch from "../../../public/ui/search/wraper";
import { FncGlobalFormatDfault } from "../../../global/function/format";

export default function UixApndixCurrcvSearch({
  qryprm,
  datefl,
}: {
  qryprm: MdlApndixSearchQueryx;
  datefl: string[];
}) {
  const [params, paramsSet] = useState<MdlApndixSearchQueryx>({
    update_apndix: qryprm.update_apndix || "",
    pagedb_apndix: qryprm.pagedb_apndix || "",
    datefl_apndix: qryprm.datefl_apndix || "",
    airlfl_apndix: qryprm.airlfl_apndix || "",
    depart_apndix: qryprm.depart_apndix || "",
    flnbfl_apndix: qryprm.flnbfl_apndix || "",
    routfl_apndix: qryprm.routfl_apndix || "",
    clssfl_apndix: qryprm.clssfl_apndix || "",
    pagenw_apndix: Number(qryprm.pagenw_apndix) || 1,
    limitp_apndix: Number(qryprm.limitp_apndix) || 15,
  });

  // Monitor change
  const [chnged, chngedSet] = useState<boolean>(false);
  useEffect(() => {
    chngedSet(false);
    paramsSet({
      update_apndix: qryprm.update_apndix || "",
      pagedb_apndix: qryprm.pagedb_apndix || "",
      datefl_apndix: qryprm.datefl_apndix || "",
      airlfl_apndix: qryprm.airlfl_apndix || "",
      depart_apndix: qryprm.depart_apndix || "",
      flnbfl_apndix: qryprm.flnbfl_apndix || "",
      routfl_apndix: qryprm.routfl_apndix || "",
      clssfl_apndix: qryprm.clssfl_apndix || "",
      pagenw_apndix: Number(qryprm.pagenw_apndix) || 1,
      limitp_apndix: Number(qryprm.limitp_apndix) || 15,
    });
  }, [qryprm]);

  // Replace params
  const timerf = useRef<NodeJS.Timeout | null>(null);
  const rplprm = FncGlobalQuerysEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    const namefl = e.currentTarget.name;
    let valuef: string | number = e.currentTarget.value;
    valuef = FncGlobalFormatDfault(namefl, valuef);
    paramsSet({ ...params, [namefl]: valuef });
    if (timerf.current) clearTimeout(timerf.current);
    timerf.current = setTimeout(async () => {
      chngedSet(true);
      rplprm([namefl, "pagenw_apndix"], [valuef as string, ""]);
    }, 1000);
  };

  // Reset function
  const resetx = () => {
    chngedSet(true);
    rplprm(
      [
        "update_apndix",
        "datefl_apndix",
        "airlfl_apndix",
        "depart_apndix",
        "flnbfl_apndix",
        "routfl_apndix",
        "clssfl_apndix",
        "pagenw_apndix",
      ],
      [String(Math.random()), "", "", "", "", "", "", ""],
    );
  };
  return (
    <UixGlobalWraperSearch
      chnged={chnged}
      downld={{
        lnk: `/apndix/currcv/downld`,
        prm: params,
      }}
      upload={null}
      lblupl=""
      namefl=""
      updtfl={null}
      resetx={resetx}
    >
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"airlfl_apndix"}
          params={params.airlfl_apndix}
          plchdr="Airline"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"flnbfl_apndix"}
          params={params.flnbfl_apndix}
          plchdr="Flight number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"routfl_apndix"}
          params={params.routfl_apndix}
          plchdr="Route"
          repprm={repprm}
          labelx=""
        />
      </div>
    </UixGlobalWraperSearch>
  );
}
