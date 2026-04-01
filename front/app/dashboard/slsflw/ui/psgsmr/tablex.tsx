"use client";

import UixGlobalTheadxTablex from "../../../public/ui/tablex/theadx";
import UixGlobalTbodyrTablex from "../../../public/ui/tablex/tbodyr";
import { MdlSlsflwPsgsmrFrntnd } from "../../model/params";
import { FncGlobalIntialObject } from "../../../global/function/format";

export default function UixSlsflwPsgsmrTablex({ arrdta }: { arrdta: MdlSlsflwPsgsmrFrntnd[] }) {
  // Dinamis
  const rawobj: MdlSlsflwPsgsmrFrntnd = FncGlobalIntialObject(arrdta[0]);

  return (
    <>
      <div className="ctable">
        <table>
          <UixGlobalTheadxTablex firsth="" mainhd={Object.keys(rawobj)} />
          <UixGlobalTbodyrTablex
            arrdta={arrdta}
            objdta={rawobj}
            datefm={[
              "datefl",
              "daterv",
              "datevc",
              "dateup",
              "datend",
              "timefl",
              "timevc",
              "timerv",
              "timeis",
              "timecr",
              "timeup",
              "mnthfl",
            ]}
            nmbrfm={["flhour", "totnta", "tottyq", "totpax", "totfae", "totqfr", "totrph"]}
            rfresh={null}
            ignore={null}
            tolink={null}
            editdt={null}
            trashx={null}
            okeupd={""}
            cxlupd={""}
          />
        </table>
      </div>
    </>
  );
}
