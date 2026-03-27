"use client";
import { useEffect, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import UixGlobalWaitngAction from "../../../public/ui/action/waitng";

export default function UixApndixApplstSelect({
  apndix,
  pagedb,
}: {
  apndix: string[];
  pagedb: string;
}) {
  const rplprm = FncGlobalQuerysEdlink();
  const [pageon, pageonSet] = useState("");
  const [chnged, chngedSet] = useState(false);
  const select = (val: string) => {
    if (pageon != val) {
      pageonSet(val);
      chngedSet(true);
      rplprm(
        [
          "update_apndix",
          "pagedb_apndix",
          "datefl_apndix",
          "airlfl_apndix",
          "depart_apndix",
          "flnbfl_apndix",
          "routfl_apndix",
          "clssfl_apndix",
          "pagenw_apndix",
        ],
        [String(Math.random()), val, "", "", "", "", "", "", "1"],
      );
    }
  };
  useEffect(() => {
    if (pagedb != "") pageonSet(pagedb);
    if (pagedb == pageon) chngedSet(false);
  }, [pagedb]);
  return (
    <div className="flexctr relative h-10 min-h-fit w-full py-3">
      <UixGlobalWaitngAction chnged={chnged} />
      <div className="afull flexstr flex-wrap gap-3">
        {apndix.map((val, idx) => (
          <div
            className={`flexctr w-fit px-3 py-1.5 text-center ${val == pageon ? "btnsbm" : "btnstb"} duration-300`}
            key={idx}
            onClick={() => select(val)}
          >
            <div>{val}</div>
            <div
              className={`flexctr ${val == pageon ? "w-8 pl-1.5" : "w-1 opacity-0"} duration-300`}
            >
              <div className="absolute h-1 w-6 rotate-45 rounded-full bg-white"></div>
              <div className="h-1.5 w-6"></div>
              <div className="absolute h-1 w-6 -rotate-45 rounded-full bg-white"></div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
