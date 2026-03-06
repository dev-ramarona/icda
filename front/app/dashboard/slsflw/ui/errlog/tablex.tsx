'use client'
import { useEffect, useState } from "react";
import { MdlPsglstErrlogDtbase } from "../../../psglst/model/params";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { ApiGlobalStatusPrcess } from "../../../global/api/status";
import { ApiPsglstPrcessManual } from "../../../psglst/api/prcess";
import { UixGlobalIconvcIgnore, UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import { FncGlobalFormatDatefm } from "../../../global/function/format";
import { MdlGlobalStatusPrcess } from "../../../global/model/params";

export default function UixPsglstErrlogTablex({
  errlog,
  update,
  status,
}: {
  errlog: MdlPsglstErrlogDtbase[];
  update: string;
  status: MdlGlobalStatusPrcess;
}) {

  // Hit the database and get interval status
  const rplprm = FncGlobalQuerysEdlink();
  const [statfn, statfnSet] = useState("Done");

  // Process function
  const prcess = async (params: MdlPsglstErrlogDtbase) => {
    rplprm(["update_global"], String(Math.random()));
    statfnSet("Wait");
    const nowprm = { ...params };
    if (status.sbrapi == 0) {
      if (params.flnbfl == "") {
        nowprm.worker = 3;
        if (params.depart == "") {
          nowprm.worker = 5;
          if (params.airlfl == "")
            nowprm.worker = 8;
        }
      }

      // Set interval to check status
      const rsp = ApiPsglstPrcessManual(nowprm);
      setTimeout(() => {
        rplprm(["update_global"], String(Math.random()));
      }, 1000);
      statfnSet(await rsp);
      setTimeout(() => statfnSet(""), 2000);
    }
  }

  // Monitor process status
  useEffect(() => {
    if (status.sbrapi == 0) statfnSet("");
    const gtstat = async () => {
      if (status.sbrapi != 0) {
        const intrvl = setInterval(async () => {
          console.log("action interval");
          const instat = await ApiGlobalStatusPrcess();
          if (instat.sbrapi == 0) {
            statfnSet("");
            rplprm(["update_global"], String(Math.random()));
            clearInterval(intrvl);
          } else statfnSet(`${instat.sbrapi}%`);
        }, 2000);
      }
    };
    gtstat();
  }, [update]);

  return (
    <>
      <div className="ctable">
        <table>
          <thead>
            <tr>
              <th className="sticky left-0">Action</th>
              {errlog && errlog.length > 0
                ? Object.entries(errlog[0]).map(([key]) => (
                  <th key={key}>
                    {key}
                  </th>
                ))
                : ""}
            </tr>
          </thead>
          <tbody>
            {errlog.map((log, idx) => (
              <tr key={idx}>
                <td className="text-center sticky left-0 z-10 drop-shadow-lg bg-white">
                  <div className="afull flexctr gap-x-1.5">
                    <div className="w-1/2 flexctr btnsbm duration-300 cursor-pointer"
                      onClick={() => prcess(log)}>
                      <div className={`absolute text-gray-300 font-bold text-xs z-10`}>
                        {statfn.includes("%") ? statfn : ""}
                      </div>
                      <div className={`${statfn != "" ? "animate-spin" : ""}`}>
                        <UixGlobalIconvcRfresh
                          bold={3}
                          color="#fff"
                          size={1.4}
                        />
                      </div>
                    </div>
                    <div className="w-1/2 flexctr btncxl duration-300 cursor-pointer"
                      onClick={() => prcess({ ...log, erignr: log.prmkey })}>
                      <UixGlobalIconvcIgnore
                        bold={3}
                        color="#fff"
                        size={1.4}
                      />
                    </div>
                  </div>
                </td>
                {Object.entries(log).map(([key, val]) => (
                  <td className="text-center" key={key}>
                    {["datefl", "timeup"].includes(key)
                      ? FncGlobalFormatDatefm(String(val))
                      : val}
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
