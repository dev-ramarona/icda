import { Dispatch, JSX, SetStateAction } from "react";
import { MdlApndixAcpedtDtbase } from "../../../apndix/model/parmas";
import {
  UixGlobalIconvcCancel,
  UixGlobalIconvcCeklis,
  UixGlobalIconvcEditdt,
} from "../../../global/ui/server/iconvc";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import { FncGlobalFormatDatefm } from "../../../global/function/format";

export default function UixGlobalTbodyxTablex({
  action,
  arrdta,
  objdta,
  objset,
  acpedt,
  datefm,
  nmbrfm,
  cnfupd,
  okeupd,
  cxlupd,
}: {
  action: (e: React.ChangeEvent<HTMLInputElement>) => void;
  arrdta: any[];
  objdta: any;
  objset: Dispatch<SetStateAction<any>>;
  acpedt: MdlApndixAcpedtDtbase[];
  datefm: string[];
  nmbrfm: string[];
  cnfupd: () => void;
  okeupd: string;
  cxlupd: string;
}) {
  const timeip = datefm.filter((v) => v.includes("time"));
  const dateip = datefm.filter((v) => v.includes("date"));
  return (
    <tbody>
      {arrdta.map((log, idx) => (
        <tr key={idx}>
          {/* Button action */}
          <td
            className={`sticky left-0 z-10 w-20 text-center drop-shadow-md ${
              cxlupd === log.prmkey
                ? "shkeit h-12 bg-red-200"
                : okeupd === log.prmkey
                  ? "shkeit h-12 bg-green-200"
                  : objdta?.prmkey === log.prmkey
                    ? "h-12 bg-sky-200"
                    : "bg-white"
            }`}
          >
            <div className="afull flexctr relative gap-x-1.5">
              <div
                className={`flexctr btnscs cursor-pointer duration-300 ${
                  objdta?.prmkey === log.prmkey
                    ? "opacity-100"
                    : "pointer-events-none opacity-0 select-none"
                }`}
                onClick={() => cnfupd()}
              >
                <UixGlobalIconvcCeklis bold={4} color="#fff" size={1.4} />
              </div>
              <div
                className={`flexctr btnwrn cursor-pointer duration-300 ${
                  objdta?.prmkey === log.prmkey
                    ? "opacity-100"
                    : "pointer-events-none opacity-0 select-none"
                }`}
                onClick={() => objset({ ...log, prmkey: "" })}
              >
                <UixGlobalIconvcCancel bold={4} color="#fff" size={1.4} />
              </div>
              <div
                className={`flexctr btnsbm absolute cursor-pointer duration-300 ${
                  objdta?.prmkey === log.prmkey
                    ? "pointer-events-none opacity-0 select-none"
                    : "opacity-100"
                }`}
                onClick={() => objset({ ...log, prmkey: log.prmkey })}
              >
                <UixGlobalIconvcEditdt bold={2.7} color="white" size={1.4} />
              </div>
            </div>
          </td>

          {/* Format edit or not edit */}
          {typeof log === "object" &&
            log !== null &&
            Object.entries(log as object).map(([key, val]) => (
              <td
                className={`z-0 text-center ${
                  cxlupd === log.prmkey
                    ? "shkeit h-12 bg-red-200"
                    : okeupd === log.prmkey
                      ? "shkeit h-12 bg-green-200"
                      : objdta?.prmkey === log.prmkey
                        ? "h-12 bg-sky-200"
                        : "h-0 bg-white"
                } duration-300`}
                key={key}
              >
                {objdta?.prmkey === log.prmkey && acpedt.some((item) => item.params === key) ? (
                  <div className="flexctr relative h-0">
                    <div className="min-w-20">
                      <UixGlobalInputxFormdt
                        typipt={
                          timeip.includes(key)
                            ? "datetime-local"
                            : dateip.includes(key)
                              ? "date"
                              : "text"
                        }
                        length={acpedt.find((item) => item.params === key)?.length}
                        queryx={key.toString()}
                        params={String(objdta[key])}
                        plchdr={""}
                        repprm={(e) => action(e)}
                        labelx=""
                      />
                    </div>
                  </div>
                ) : (
                  <div>
                    {datefm.includes(key) ? (
                      FncGlobalFormatDatefm(String(val))
                    ) : nmbrfm.includes(key) ? (
                      <div className="text-right">{val.toLocaleString("en-US")}</div>
                    ) : (
                      val
                    )}
                  </div>
                )}
              </td>
            ))}
        </tr>
      ))}
    </tbody>
  );
}
