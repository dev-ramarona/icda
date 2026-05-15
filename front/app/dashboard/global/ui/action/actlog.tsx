"use client";
import { useEffect, useState } from "react";
import { FncGlobalFormatDatefm } from "../../function/format";
import { ApiAllusrStatusPrcess } from "../../../allusr/api/status";
import { MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { FncGlobalQuerysEdlink } from "../../function/querys";
import { UixGlobalIconvcRfresh } from "../server/iconvc";

export default function UixGlobalActlogTablex({
  actlog,
  status,
  update,
}: {
  actlog: Record<string, any>[];
  status: MdlAllusrStatusPrcess;
  update: string;
}) {
  // Monitor process status
  const rplprm = FncGlobalQuerysEdlink();
  const [statfn, statfnSet] = useState("Process");
  useEffect(() => {
    if (status.sbrapi == 0) statfnSet("");
    const gtstat = async () => {
      if (status.sbrapi != 0) {
        const intrvl = setInterval(async () => {
          const instat = await ApiAllusrStatusPrcess();
          if (instat.sbrapi == 0) {
            statfnSet("");
            rplprm(["update_global"], String(Math.random()));
            clearInterval(intrvl);
          } else statfnSet(`${instat.sbrapi}%`);
        }, 5000);
      }
    };
    gtstat();
  }, [update]);

  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-gray-200">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              {actlog && actlog.length > 0
                ? Object.entries(actlog[0]).map(([key]) =>
                    key != "prmkey" ? (
                      <th key={key} className="thhead">
                        {key}
                      </th>
                    ) : (
                      ""
                    ),
                  )
                : ""}
            </tr>
          </thead>
          <tbody className="bg-white text-slate-700">
            {statfn != "" && (
              <tr className="group h-8">
                <td className="tdbody" colSpan={4}>
                  <div className="flexstr px-1">
                    <div>Data on progress {statfn}</div>
                    <div className="w-fit animate-spin">
                      <UixGlobalIconvcRfresh bold={3} color="gray " size={1.4} />
                    </div>
                  </div>
                </td>
              </tr>
            )}
            {actlog.map((log, idx) => (
              <tr className="group h-8" key={idx}>
                {Object.entries(log).map(([key, val]) => (
                  <td className="tdbody text-center" key={key}>
                    {["dateup", "datefl", "timeup"].includes(key)
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
