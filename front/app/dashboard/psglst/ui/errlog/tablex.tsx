'use client'
import { MdlPsglstErrlogDtbase } from "../../model/params";
import { useEffect, useState } from "react";
import { ApiPsglstPrcessManual } from "../../api/prcess";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { UixGlobalIconvcIgnore, UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import { FncGlobalFormatDatefm } from "../../../global/function/format";
import { MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { ApiAllusrStatusPrcess } from "../../../allusr/api/status";

export default function UixPsglstErrlogTablex({
  errlog,
  update,
  status,
}: {
  errlog: MdlPsglstErrlogDtbase[];
  update: string;
  status: MdlAllusrStatusPrcess
}) {

  // Hit the database and get interval status
  const rplprm = FncGlobalQuerysEdlink();
  const [statfn, statfnSet] = useState("Done");
  const [onpkey, onpkeySet] = useState("");

  // Process function
  const prcess = async (params: MdlPsglstErrlogDtbase) => {
    rplprm(["update_global"], String(Math.random()));
    statfnSet("Wait");
    onpkeySet(params.prmkey);
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
      setTimeout(() => { statfnSet(""), onpkeySet(""); }, 2500);
    }
  }

  // Monitor process status
  useEffect(() => {
    if (status.sbrapi == 0) statfnSet("");
    const gtstat = async () => {
      if (status.sbrapi != 0) {
        const intrvl = setInterval(async () => {
          console.log("action interval");
          const instat = await ApiAllusrStatusPrcess();
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
                <td className={`text-center sticky left-0 z-10 drop-shadow-lg 
                ${onpkey === log.prmkey && statfn === "Success" ? "bg-green-200 shkeit" :
                    onpkey === log.prmkey && statfn === "Failed" ? "bg-red-200 shkeit" :
                      onpkey === log.prmkey ? "bg-cyan-200" : "bg-white"} duration-300`}>
                  <div className="afull flexctr gap-x-1.5">
                    <div className="w-1/2 flexctr btnsbm duration-300 cursor-pointer"
                      onClick={() => prcess(log)}>
                      <div className={`absolute text-gray-300 font-bold text-[0.5rem] z-10`}>
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
                  <td className={`text-center 
                  ${onpkey === log.prmkey && statfn === "Success" ? "bg-green-200 shkeit" :
                      onpkey === log.prmkey && statfn === "Failed" ? "bg-red-200 shkeit" :
                        onpkey === log.prmkey ? "bg-cyan-200" : "bg-white"} 
                      duration-300`} key={key}>
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
