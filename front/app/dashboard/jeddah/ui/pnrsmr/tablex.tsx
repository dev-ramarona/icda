"use client";

import UixGlobalTheadxTablex from "../../../global/ui/tablex/theadx";
import UixGlobalTbodyrTablex from "../../../global/ui/tablex/tbodyr";
import { FncGlobalIntialObject } from "../../../global/function/format";
import { MdlJeddahPrcessDtbase } from "../../model/params";

export default function UixSlsflwPsgsmrTablex({ arrdta }: { arrdta: MdlJeddahPrcessDtbase[] }) {
  // Dinamis
  const rawobj: MdlJeddahPrcessDtbase = FncGlobalIntialObject(arrdta[0]);

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
            nmbrfm={[
              "flhour",
              "totnta",
              "tottyq",
              "tottyr",
              "totpax",
              "totfae",
              "totqfr",
              "totrph",
              "totrev",
              "totcph",
              "costph",
            ]}
            spclfm={[""]}
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
