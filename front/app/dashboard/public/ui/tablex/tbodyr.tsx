import {
  UixGlobalIconvcTrashx,
  UixGlobalIconvcEditdt,
  UixGlobalIconvcIgnore,
  UixGlobalIconvcRfresh,
  UixGlobalIconvcTolink,
} from "../../../global/ui/server/iconvc";
import { FncGlobalFormatDatefm } from "../../../global/function/format";

export default function UixGlobalTbodyrTablex({
  arrdta,
  objdta,
  datefm,
  nmbrfm,
  rfresh,
  ignore,
  tolink,
  editdt,
  trashx,
  okeupd,
  cxlupd,
}: {
  arrdta: any[];
  objdta: any;
  datefm: string[];
  nmbrfm: string[];
  rfresh: (log: any) => void | null;
  ignore: (log: any) => void | null;
  tolink: (log: any) => void | null;
  editdt: (log: any) => void | null;
  trashx: (log: any) => void | null;
  okeupd: string;
  cxlupd: string;
}) {
  return (
    <tbody>
      {arrdta.map((log, idx) => (
        <tr key={idx}>
          {/* Button action */}
          {(rfresh || ignore || tolink || editdt || trashx) && (
            <td
              className={`sticky left-0 z-10 w-20 text-center drop-shadow-md ${
                cxlupd === log.prmkey
                  ? "shkeit h-12 bg-red-200"
                  : okeupd === log.prmkey
                    ? "shkeit h-12 bg-green-200"
                    : objdta?.prmkey === log.prmkey
                      ? "h-12 bg-sky-200"
                      : "bg-white"
              } duration-300`}
            >
              <div className="afull flexctr relative gap-x-1.5">
                {rfresh && (
                  <div
                    className={`flexctr cursor-pointer duration-300 ${
                      objdta?.prmkey === "all" ? "btnoff" : "btnsbm"
                    }`}
                    onClick={() => rfresh(log)}
                  >
                    <div
                      className={`${(objdta?.prmkey === log.prmkey || objdta?.prmkey === "all") && "animate-spin"}`}
                    >
                      <UixGlobalIconvcRfresh bold={3} color="#fff" size={1.4} />
                    </div>
                  </div>
                )}
                {ignore && (
                  <div
                    className={`flexctr cursor-pointer duration-300 ${
                      objdta?.prmkey === "all" ? "btnoff" : "btncxl"
                    }`}
                    onClick={() => ignore(log)}
                  >
                    <div
                      className={`${(objdta?.prmkey === log.prmkey || objdta?.prmkey === "all") && "animate-spin"}`}
                    >
                      <UixGlobalIconvcIgnore bold={3} color="#fff" size={1.4} />
                    </div>
                  </div>
                )}
                {tolink && (
                  <div
                    className={`flexctr btnwrn cursor-pointer duration-300`}
                    onClick={() => tolink(log)}
                  >
                    <UixGlobalIconvcTolink bold={3} color="#fff" size={1.4} />
                  </div>
                )}
                {editdt && (
                  <div
                    className={`flexctr btnwrn cursor-pointer duration-300`}
                    onClick={() => editdt(log)}
                  >
                    <UixGlobalIconvcEditdt bold={3} color="#fff" size={1.4} />
                  </div>
                )}
                {trashx && (
                  <div
                    className={`flexctr btnwrn cursor-pointer duration-300`}
                    onClick={() => trashx(log)}
                  >
                    <UixGlobalIconvcTrashx bold={3} color="#fff" size={1.4} />
                  </div>
                )}
              </div>
            </td>
          )}

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
                <div>
                  {datefm.includes(key) ? (
                    FncGlobalFormatDatefm(String(val))
                  ) : nmbrfm.includes(key) ? (
                    <div className="text-right">{val.toLocaleString("en-US")}</div>
                  ) : (
                    val
                  )}
                </div>
              </td>
            ))}
        </tr>
      ))}
    </tbody>
  );
}
