import { Dispatch, SetStateAction } from "react";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import {
  UixGlobalIconvcAddpls,
  UixGlobalIconvcCancel,
  UixGlobalIconvcCeklis,
} from "../../../global/ui/server/iconvc";

export default function UixGlobalTfootxTablex({
  actadd,
  objdef,
  objdta,
  objset,
  exclde,
  datefm,
  cnfupd,
  okeupd,
  cxlupd,
}: {
  actadd: (e: React.ChangeEvent<HTMLInputElement>, actmod: "add" | "add") => void;
  objdef: any;
  objdta: any;
  objset: Dispatch<SetStateAction<any>>;
  exclde: string[];
  datefm: string[];
  cnfupd: () => void;
  okeupd: string;
  cxlupd: string;
}) {
  const timeip = datefm.filter((v) => v.includes("time"));
  const dateip = datefm.filter((v) => v.includes("date"));
  return (
    <tfoot>
      <tr className="sticky bottom-0 z-20">
        <td
          className={`sticky left-0 z-10 w-20 text-center drop-shadow-md ${
            cxlupd === "add"
              ? "shkeit bg-red-200"
              : okeupd === "add"
                ? "shkeit bg-green-200"
                : objdta?.prmkey === "add"
                  ? "bg-sky-200"
                  : "bg-cyan-600"
          }`}
        >
          <div className="afull flexctr relative gap-x-1.5">
            <div
              className={`flexctr btnscs cursor-pointer duration-300 ${
                objdta?.prmkey === "add"
                  ? "opacity-100"
                  : "pointer-events-none opacity-0 select-none"
              }`}
              onClick={() => cnfupd()}
            >
              <UixGlobalIconvcCeklis bold={4} color="#fff" size={1.4} />
            </div>
            <div
              className={`flexctr btnwrn cursor-pointer duration-300 ${
                objdta?.prmkey === "add"
                  ? "opacity-100"
                  : "pointer-events-none opacity-0 select-none"
              }`}
              onClick={() => objset({ ...objdef, prmkey: "" })}
            >
              <UixGlobalIconvcCancel bold={4} color="#fff" size={1.4} />
            </div>
            <div
              className={`flexctr btnstb absolute cursor-pointer duration-300 ${
                objdta?.prmkey === "add"
                  ? "pointer-events-none opacity-0 select-none"
                  : "opacity-100"
              }`}
              onClick={() => objset({ ...objdef, prmkey: "add" })}
            >
              <UixGlobalIconvcAddpls bold={2.7} color="#0092b8" size={1.4} />
            </div>
          </div>
        </td>
        {Object.keys(objdef).map((val, key) => (
          <td
            className={`z-0 text-center ${
              cxlupd === "add"
                ? "shkeit h-12 bg-red-200"
                : okeupd === "add"
                  ? "shkeit h-12 bg-green-200"
                  : objdta?.prmkey === "add"
                    ? "h-12 bg-sky-200"
                    : "h-0 bg-cyan-600"
            } duration-300`}
            key={key}
          >
            <div className="flexctr relative">
              {objdta.prmkey == "add" && !exclde.includes(val) ? (
                <div className="flexctr relative h-0">
                  <div className="min-w-20">
                    <UixGlobalInputxFormdt
                      typipt={
                        timeip.includes(val)
                          ? "datetime-local"
                          : dateip.includes(val)
                            ? "date"
                            : "text"
                      }
                      length={24}
                      queryx={val.toString()}
                      params={String(objdta[val])}
                      plchdr={val}
                      repprm={(e) => actadd(e, "add")}
                      labelx=""
                    />
                  </div>
                </div>
              ) : (
                <div className="afull bg-red-300 shadow-2xl"></div>
              )}
            </div>
          </td>
        ))}
      </tr>
    </tfoot>
  );
}
