"use client";
import { useEffect, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import UixGlobalInputxFormdt from "../../../global/ui/action/inputx";
import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import { MdlAllusrSearchParams } from "../../model/params";

export default function UixCrtusrAllusrSearch({ prmAllusr }: { prmAllusr: MdlAllusrSearchParams }) {
  const [params, paramsSet] = useState<MdlAllusrSearchParams>({
    usredt: prmAllusr.usredt || "",
    stfnme: prmAllusr.stfnme || "",
    usrnme: prmAllusr.usrnme || "",
    stfeml: prmAllusr.stfeml || "",
    limitp: prmAllusr.limitp || 1,
    pagenw: prmAllusr.pagenw || 15,
    update: prmAllusr.update || "",
  });

  // Monitor change
  const [chnged, chngedSet] = useState<boolean>(false);
  useEffect(() => {
    chngedSet(false);
    paramsSet({
      usredt: prmAllusr.usredt || "",
      stfnme: prmAllusr.stfnme || "",
      usrnme: prmAllusr.usrnme || "",
      stfeml: prmAllusr.stfeml || "",
      limitp: prmAllusr.limitp || 1,
      pagenw: prmAllusr.pagenw || 15,
      update: prmAllusr.update || "",
    });
  }, [prmAllusr]);

  // Replace params
  const rplprm = FncGlobalQuerysEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    chngedSet(true);
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
    rplprm([namefl, "pagenw"], [valuef, ""]);
  };

  // Reset function
  const resetx = () => {
    chngedSet(true);
    rplprm(["stfnme", "usrnme", "stfeml", "limitp", "usredt"], "");
  };
  return (
    <div className="flexctr relative h-24 min-h-fit w-full py-3">
      <div
        className={`${chnged ? "h-10 w-16 translate-y-0" : "h-0 w-0 -translate-y-10 opacity-0"} flexctr absolute z-10 rounded-xl bg-white px-5 py-2 ring-2 ring-sky-300 duration-300`}
      >
        <div>Wait</div>
        <div className="animate-spin">
          <UixGlobalIconvcRfresh bold={2} color="black" size={1} />
        </div>
      </div>
      <div
        className={`afull flexstr flex-wrap gap-y-3 ${chnged ? "animate-pulse select-none" : ""} duration-300`}
      >
        <div className="flexctr relative h-11 w-1/2 md:w-28">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"stfnme"}
            params={params.stfnme}
            plchdr="Staff Name"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="flexctr relative h-11 w-1/2 md:w-28">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"usrnme"}
            params={params.usrnme}
            plchdr="Username"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="flexctr relative h-11 w-1/2 md:w-28">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"stfeml"}
            params={params.stfeml}
            plchdr="Staff Email"
            repprm={repprm}
            labelx=""
          />
        </div>
      </div>
      <div
        className={`flexend w-1/3 flex-wrap gap-3 px-3 ${chnged ? "animate-pulse select-none" : ""} duration-300`}
      >
        <div className="flexctr relative h-10 w-full md:w-28">
          <div className="afull btnwrn flexctr" onClick={() => resetx()}>
            Reset
          </div>
        </div>
      </div>
    </div>
  );
}
