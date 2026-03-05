"use client";
import { MdlSlsflwAcpedtDtbase, MdlSlsflwPsgdtlFrntnd } from "../../model/params";
import { ApiSlsflwPsgdtlUpdate } from "../../api/psgdtl";
import { useState } from "react";
import { mdlGlobalAllusrCookie } from "../../../global/model/params";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatCpnfmt, FncGlobalFormatDatefm, FncGlobalFormatRoutfl } from "../../../global/function/format";
import { UixGlobalIconvcCancel, UixGlobalIconvcCeklis, UixGlobalIconvcEditdt } from "../../../global/ui/server/iconvc";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";

export default function UixSlsflwDetailTablex({
  detail,
  acpedt,
  cookie,
}: {
  detail: MdlSlsflwPsgdtlFrntnd[];
  acpedt: MdlSlsflwAcpedtDtbase[];
  cookie: mdlGlobalAllusrCookie;
}) {
  const [edtobj, edtobjSet] = useState<MdlSlsflwPsgdtlFrntnd>();
  const [okeupd, okeupdSet] = useState<string>("");
  const [cxlupd, cxlupdSet] = useState<string>("");
  const [cxlrsp, cxlrspSet] = useState<string>("");
  const rplprm = FncGlobalQuerysEdlink();
  const actedt = (e: React.ChangeEvent<HTMLInputElement>) => {
    const key = e.currentTarget.id;
    let val: string | number = e.currentTarget.value;
    if (key == "routvc") val = FncGlobalFormatRoutfl(val);
    else if (key == "cpnbvc") val = FncGlobalFormatCpnfmt(val);
    else if (["tktnbr", "flnbvc"].includes(key))
      val = val.replace(/[^0-9]/g, "");
    else if (["ntafvc", "ntaffl", "qsrcvc"].includes(key))
      val = (val);
    else val = val.toUpperCase();
    edtobjSet({
      ...edtobj,
      [key]: val,
    } as MdlSlsflwPsgdtlFrntnd);
  };

  // Confirm update retail or series
  const update = async (log: MdlSlsflwPsgdtlFrntnd) => {
    console.log(log);
    const rspupd: string = await ApiSlsflwPsgdtlUpdate({
      ...log,
      ntafvc: Number(log.ntafvc),
      ntaffl: Number(log.ntaffl),
      qsrcvc: Number(log.qsrcvc)
    });
    edtobjSet({ ...log, prmkey: "" })
    if (rspupd == "success") {
      okeupdSet(log.prmkey);
    } else {
      cxlupdSet(log.prmkey);
      cxlrspSet(rspupd);
    }
    setTimeout(() => {
      okeupdSet("");
      cxlupdSet("");
      cxlrspSet("");
      rplprm(["update_global"], String(Math.random()));
    }, 1000);
  };

  return (
    <>
      <div className="ctable">
        <table>
          <thead>
            <tr>
              <th className="thhead sticky left-0">Action</th>
              {detail && detail.length > 0
                ? Object.entries(detail[0]).map(([key]) => (
                  <th key={key} className="thhead">
                    {key}
                  </th>
                ))
                : ""}
            </tr>
          </thead>
          <tbody>
            {detail.map((log, idx) => (
              <tr key={idx}>
                <td
                  className={`text-center sticky left-0 z-10 shadow-md drop-shadow-lg 
                    ${edtobj?.prmkey === log.prmkey ? "bg-sky-200" :
                      okeupd === log.prmkey ? "bg-green-400 shkeit" :
                        cxlupd === log.prmkey ? "bg-red-400 shkeit" : "bg-white"}`}

                >
                  <div className="afull flexctr gap-x-1.5 relative">
                    <div
                      className={`flexctr btnsbm duration-300 cursor-pointer ${edtobj?.prmkey === log.prmkey
                        ? "opacity-100"
                        : "opacity-0 select-none pointer-events-none"
                        }`}
                      onClick={() => update(edtobj as MdlSlsflwPsgdtlFrntnd)}
                    >
                      <UixGlobalIconvcCeklis
                        bold={2.5}
                        color="#53eafd"
                        size={1.4}
                      />
                    </div>
                    <div
                      className={`flexctr btnsbm duration-300 cursor-pointer ${edtobj?.prmkey === log.prmkey
                        ? "opacity-100"
                        : "opacity-0 select-none pointer-events-none"
                        }`}
                      onClick={() => edtobjSet({ ...log, prmkey: "" })}
                    >
                      <UixGlobalIconvcCancel
                        bold={2.5}
                        color="#fb2c36"
                        size={1.4}
                      />
                    </div>
                    <div
                      className={`absolute flexctr btnsbm duration-300 cursor-pointer ${edtobj?.prmkey === log.prmkey
                        ? "opacity-0 select-none pointer-events-none"
                        : "opacity-100"
                        }`}
                      onClick={() => edtobjSet({ ...log, updtby: cookie.usrnme, prmkey: log.prmkey })}
                    >
                      <UixGlobalIconvcEditdt
                        bold={2.5}
                        color="white"
                        size={1.4}
                      />
                    </div>
                  </div>
                  <div className={`${cxlupd === log.prmkey ? "h-8 flexctr font-semibold text-white" :
                    "h-0 opacity-0"} duration-300`}>{cxlrsp}</div>
                </td>
                {Object.entries(log).map(([key, val]) => (
                  <td
                    className={`text-center z-0 h-8 w-fit ${edtobj?.prmkey === log.prmkey ? "bg-sky-200" :
                      okeupd === log.prmkey ? "bg-green-400 shkeit" :
                        cxlupd === log.prmkey ? "bg-red-400 shkeit" : "bg-white"}`}
                    key={key}
                  >
                    {edtobj?.prmkey === log.prmkey &&
                      acpedt.some((item) => item.params === key) ? (
                      <div className="relative flexctr">
                        <span className="invisible">
                          XXXXXXXXXXXXX{String(edtobj[key as keyof typeof edtobj])}
                        </span>
                        <div className="h-8 absolute">
                          <UixGlobalInputxFormdt
                            typipt={key == "timeis" ? "datetime-local" : key == "datevc" ? "date" : "text"}
                            length={
                              acpedt.find((item) => item.params === key)?.length
                            }
                            queryx={key.toString()}
                            params={String(edtobj[key as keyof typeof edtobj])}
                            plchdr=""
                            repprm={actedt}
                            labelx=""
                          />
                        </div>
                      </div>
                    ) : (
                      <div>
                        {["datefl", "daterv", "datevc", "timefl", "timevc", "timerv", "timeis", "timecr", "mnthfl"].includes(key)
                          ? FncGlobalFormatDatefm(String(val))
                          : ["ntaffl", "ntafvc", "yqtxfl", "yqtxvc", "qsrcrw", "qsrcvc"].includes(key)
                            ? <div className="text-right">{val.toLocaleString("en-US")}</div>
                            : val}
                      </div>
                    )}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}
