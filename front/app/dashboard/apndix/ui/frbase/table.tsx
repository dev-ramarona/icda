"use client";

import { useRef, useState } from "react";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import { MdlApndixAcpedtDtbase, MdlApndixFrbaseFrntnd } from "../../model/parmas";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatDefault, FncGlobalFormatInptdt } from "../../../global/function/format";
import { ApiApndixUpdateDtbase } from "../../api/dtbase";
import UixGlobalTheadxTablex from "../../../public/ui/tablex/theadx";
import UixGlobalTbodyxTablex from "../../../public/ui/tablex/tbodyx";
import UixGlobalTfootxTablex from "../../../public/ui/tablex/tfootx";
import UixGlobalConfrmAction from "../../../public/ui/action/confrm";
import { MdlGlobalConfrmAction } from "../../../public/model/params";

export default function UixPsglstFrbaseTablex({
  arrdta,
  pagedb,
  acpedt,
  cookie,
}: {
  arrdta: MdlApndixFrbaseFrntnd[];
  pagedb: string;
  acpedt: MdlApndixAcpedtDtbase[];
  cookie: mdlAllusrCookieObjson;
}) {
  // Dinamis
  const exclde = ["prmkey", "scdkey", "hstory", "updtby"];
  const inclde = acpedt.map((item) => item.params);
  const rawobj: MdlApndixFrbaseFrntnd = {
    prmkey: "",
    scdkey: "",
    airlfl: "",
    clssfl: "",
    routfl: "",
    frbcde: "",
    frbnta: "",
    frbsbr: "",
    datend: "",
    hstory: "",
    updtby: "",
  };
  let othrfn: Function | undefined;
  othrfn = (objdta: MdlApndixFrbaseFrntnd) => {
    Object.entries(objdta).map(([k, v]) => {
      if (["frbnta", "frbsbr"].includes(k)) {
        objdta[k] = Number(v);
      } else if (["datend"].includes(k)) {
        objdta[k] = Number(FncGlobalFormatInptdt(v));
      }
    });
    return objdta;
  };

  // Variable default
  const [objdta, objdtaSet] = useState(rawobj);
  const [okeupd, okeupdSet] = useState<string>("");
  const [cxlupd, cxlupdSet] = useState<string>("");
  const [confrm, confrmSet] = useState<boolean>(false);
  const [confdt, confdtSet] = useState<MdlGlobalConfrmAction[]>([]);

  // edit params
  const rplprm = FncGlobalQuerysEdlink();
  const actedt = (e: React.ChangeEvent<HTMLInputElement>) => {
    const key = e.currentTarget.id;
    let val: string | number = e.currentTarget.value;
    val = FncGlobalFormatDefault(key, val);
    objdtaSet({ ...objdta, [key]: val });
  };

  //
  const refedt = useRef<NodeJS.Timeout | null>(null);
  const cnfupd = () => {
    if (refedt.current) clearTimeout(refedt.current);
    const confst = [];
    let emptys = false;
    Object.entries(objdta).map(([k, v]) => {
      if (!exclde.includes(k)) {
        confst.push({ paramx: k, valuex: v });
        if ((v == "" || v == 0) && inclde.includes(k)) {
          cxlupdSet(objdta.prmkey);
          emptys = true;
        }
      }
    });
    if (!emptys) {
      confrmSet(true);
      confdtSet(confst);
    } else refedt.current = setTimeout(() => cxlupdSet(""), 1000);
  };

  // Confirm update retail or series
  const refupd = useRef<NodeJS.Timeout | null>(null);
  const goupdt = async () => {
    confrmSet(false);
    if (refupd.current) clearTimeout(refupd.current);
    refupd.current = setTimeout(async () => {
      const copydt = othrfn?.(objdta) ?? objdta;
      copydt.updtby = cookie.usrnme;
      objdta.updtby = cookie.usrnme;
      const rspupd: string = await ApiApndixUpdateDtbase(copydt, pagedb);
      objdtaSet({ ...copydt, prmkey: "" });
      if (rspupd == "success") {
        okeupdSet(copydt.prmkey);
      } else cxlupdSet(copydt.prmkey);
      setTimeout(() => {
        okeupdSet("");
        cxlupdSet("");
        rplprm(["update"], String(Math.random()));
      }, 1000);
    }, 1000);
  };

  return (
    <>
      <UixGlobalConfrmAction
        confrm={confrm}
        confdt={confdt}
        action={"update"}
        goupdt={goupdt}
        confrmSet={confrmSet}
      />
      <div className={`ctable`}>
        <table>
          <UixGlobalTheadxTablex firsth="action" mainhd={Object.keys(rawobj)} />
          <UixGlobalTbodyxTablex
            action={actedt}
            arrdta={arrdta}
            objdta={objdta}
            objset={objdtaSet}
            acpedt={acpedt}
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
            nmbrfm={["ntaffl", "ntafvc", "yqtxfl", "yqtxvc", "qsrcrw", "qsrcvc"]}
            cnfupd={cnfupd}
            okeupd={okeupd}
            cxlupd={cxlupd}
          />
          <UixGlobalTfootxTablex
            actadd={actedt}
            objdef={rawobj}
            objdta={objdta}
            objset={objdtaSet}
            exclde={exclde}
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
            cnfupd={cnfupd}
            okeupd={okeupd}
            cxlupd={cxlupd}
          />
        </table>
      </div>
    </>
  );
}
