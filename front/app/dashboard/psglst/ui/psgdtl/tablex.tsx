"use client";

import { MdlPsglstPsgdtlFrntnd } from "../../model/params";
import { useRef, useState } from "react";
import { ApiPsglstPsgdtlUpdate } from "../../api/psgdtl";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import {
  FncGlobalFormatDfault,
  FncGlobalFormatInptdt,
  FncGlobalIntialObject,
} from "../../../global/function/format";
import { MdlApndixAcpedtDtbase } from "../../../apndix/model/parmas";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import UixGlobalTbodyxTablex from "../../../global/ui/tablex/tbodyx";
import UixGlobalTheadxTablex from "../../../global/ui/tablex/theadx";
import UixGlobalConfrmAction from "../../../global/ui/action/confrm";
import { MdlGlobalConfrmAction } from "../../../global/model/params";

export default function UixPsglstDetailTablex({
  arrdta,
  acpedt,
  cookie,
  fmtdef,
}: {
  arrdta: MdlPsglstPsgdtlFrntnd[];
  acpedt: MdlApndixAcpedtDtbase[];
  cookie: mdlAllusrCookieObjson;
  fmtdef: boolean;
}) {
  // Dinamis
  const exclde = ["prmkey", "hstory", "updtby"];
  const inclde = acpedt.map((item) => item.params);
  const rawobj: MdlPsglstPsgdtlFrntnd = FncGlobalIntialObject(arrdta[0]);
  let othrfn: Function | undefined;
  othrfn = (objdta: MdlPsglstPsgdtlFrntnd) => {
    Object.entries(objdta).map(([k, v]) => {
      if (["cpnbvc"].includes(k)) {
        objdta[k] = Number(v);
      } else if (["timeis"].includes(k)) {
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
    val = FncGlobalFormatDfault(key, val);
    objdtaSet({ ...objdta, [key]: val });
  };

  // Action
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
      const rspupd: string = await ApiPsglstPsgdtlUpdate(copydt);
      objdtaSet({ ...copydt, prmkey: "" });
      if (rspupd == "success") {
        okeupdSet(copydt.prmkey);
      } else cxlupdSet(copydt.prmkey);
      setTimeout(() => {
        okeupdSet("");
        cxlupdSet("");
        rplprm(["update_global"], String(Math.random()));
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
          <UixGlobalTheadxTablex firsth={fmtdef ? "action" : null} mainhd={Object.keys(rawobj)} />
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
            cnfupd={fmtdef ? cnfupd : null}
            okeupd={okeupd}
            cxlupd={cxlupd}
          />
        </table>
      </div>
    </>
  );
}
