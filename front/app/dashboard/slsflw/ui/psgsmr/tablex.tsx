"use client";

import { useState } from "react";
import { MdlSlsflwAcpedtDtbase, MdlSlsflwPsgsmrFrntnd } from "../../model/params";
import { FncGlobalFormatCpnfmt, FncGlobalFormatDatefm, FncGlobalFormatRoutfl } from "../../../global/function/format";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";

export default function UixSlsflwPsgsmrTablex({
  Psgsmr,
  acpedt,
  cookie,
}: {
  Psgsmr: MdlSlsflwPsgsmrFrntnd[];
  acpedt: MdlSlsflwAcpedtDtbase[];
  cookie: mdlAllusrCookieObjson;
}) {
  const [edtobj, edtobjSet] = useState<MdlSlsflwPsgsmrFrntnd>();
  const actedt = (e: React.ChangeEvent<HTMLInputElement>) => {
    const key = e.currentTarget.id;
    let val: string | number = e.currentTarget.value;
    if (key == "routvc") val = FncGlobalFormatRoutfl(val);
    else if (key == "cpnbvc") val = FncGlobalFormatCpnfmt(val);
    else if (["tktnbr", "flnbvc"].includes(key))
      val = val.replace(/[^0-9]/g, "");
    else if (["ntafvc", "ntaffl"].includes(key))
      val = Number(val);
    else val = val.toUpperCase();
    edtobjSet({
      ...edtobj,
      [key]: val,
    } as MdlSlsflwPsgsmrFrntnd);
  };


  return (
    <>
      <div className="ctable">
        <table>
          <thead>
            <tr>
              {Psgsmr && Psgsmr.length > 0
                ? Object.entries(Psgsmr[0]).map(([key]) => (
                  <th key={key}>
                    {key}
                  </th>
                ))
                : ""}
            </tr>
          </thead>
          <tbody>
            {Psgsmr.map((log, idx) => (
              <tr key={idx}>
                {
                  Object.entries(log).map(([key, val]) => (
                    <td
                      className={`text-center z-0 w-fit`}
                      key={key}
                    >
                      {edtobj?.prmkey === log.prmkey &&
                        acpedt.some((item) => item.params === key) ? (
                        <div className="relative flexctr">
                          <span className="invisible">
                            XXXXXXXXXXXXX{String(edtobj[key as keyof typeof edtobj])}
                          </span>
                          <div className="absolute">
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
                          {["datefl", "mnthfl"].includes(key)
                            ? FncGlobalFormatDatefm(String(val))
                            : ["totnta", "tottyq", "totpax", "totfae", "totqfr"].includes(key)
                              ? <div className="w-full text-right">{val.toLocaleString("en-US")}</div>
                              : val}
                        </div>
                      )}
                    </td>
                  ))
                }
              </tr>
            ))}
          </tbody>
        </table>
      </div >
    </>
  );
}
