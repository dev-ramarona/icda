"use client";
import { useEffect, useRef, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatDefault } from "../../../global/function/format";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import { MdlApndixSearchQueryx } from "../../model/parmas";
import UixGlobalWraperSearch from "../../../public/ui/search/wraper";

export default function UixApndixProvncSearch({ qryprm }: { qryprm: MdlApndixSearchQueryx }) {
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
    const namefl = e.currentTarget.id;
    let valuef: string | number = e.currentTarget.value;
    valuef = FncGlobalFormatDefault(namefl, valuef);
    paramsSet({ ...params, [namefl]: valuef });
    if (timerf.current) clearTimeout(timerf.current);
    timerf.current = setTimeout(async () => {
      chngedSet(true);
      rplprm([namefl, "pagenw_psgdtl"], [valuef as string, ""]);
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
        lnk: `${process.env.NEXT_PUBLIC_URL_SERVER}/apndix/provnc/downld`,
        prm: params,
      }}
      resetx={resetx}
    >
      <div className="flexctr relative h-11 w-1/2 md:w-28">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={null}
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
