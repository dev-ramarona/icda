'use client'
import { useEffect, useState } from "react";
import { MdlPsglstErrlogDtbase } from "../../../psglst/model/params";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { ApiGlobalStatusIntrvl, ApiGlobalStatusPrcess } from "../../../global/api/status";
import { ApiPsglstPrcessManual } from "../../../psglst/api/prcess";
import { UixGlobalIconvcIgnore, UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import { FncGlobalFormatDatefm } from "../../../global/function/format";

export default function UixPsglstErrlogTablex({
  errlog,
  update,
}: {
  errlog: MdlPsglstErrlogDtbase[];
  update: string;
}) {

  // Hit the database and get interval status
  const rplprm = FncGlobalQuerysEdlink();
  const [statfn, statfnSet] = useState("Done");
  const [onpkey, onpkeySet] = useState("Done");

  // Hit the database and get interval status
  const prcess = async (params: MdlPsglstErrlogDtbase) => {
    const status = await ApiGlobalStatusPrcess();
    const nowParams = { ...params };
    if (status.sbrapi == 0) {
      if (params.flnbfl == "") {
        nowParams.worker = 3;
        if (params.depart == "") {
          nowParams.worker = 5;
          if (params.airlfl == "")
            nowParams.worker = 8;
        }
      }

      // Cek is admin or not
      statfnSet("Wait");
      onpkeySet(params.prmkey);
      rplprm(["update_global"], String(Math.random()));
      ApiPsglstPrcessManual(nowParams);
      await ApiGlobalStatusIntrvl(statfnSet, "sbrapi");
    } else statfnSet(`Wait ${status.sbrapi}%`);
  };

  // Monitor process status
  useEffect(() => {
    const gtstat = async () => {
      const status = await ApiGlobalStatusPrcess();
      statfnSet(status.sbrapi == 0 ? "Done" : `Wait ${status.sbrapi}%`);
      if (status.sbrapi != 0) {
        await ApiGlobalStatusIntrvl(statfnSet, "sbrapi");
      } else statfnSet("Done");
    };
    gtstat();
  }, [update]);

  // refresh page
  useEffect(() => {
    if (statfn == "Process Done") setTimeout(() => {
      rplprm(["update_global"], String(Math.random()));
    }, 1000);
  }, [statfn, rplprm])

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
                      <div className={`${onpkey == log.prmkey ? "animate-spin" : ""}`}>
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
