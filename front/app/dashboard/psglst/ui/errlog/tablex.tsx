"use client";
import { MdlPsglstErrlogDtbase } from "../../model/params";
import { useEffect, useRef, useState } from "react";
import { ApiPsglstPrcessManual } from "../../api/prcess";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { ApiAllusrStatusPrcess } from "../../../allusr/api/status";
import UixGlobalTheadxTablex from "../../../public/ui/tablex/theadx";
import UixGlobalTbodyrTablex from "../../../public/ui/tablex/tbodyr";

export default function UixPsglstErrlogTablex({
  errlog,
  update,
  status,
}: {
  errlog: MdlPsglstErrlogDtbase[];
  update: string;
  status: MdlAllusrStatusPrcess;
}) {
  // Dinamis
  const rawobj: MdlPsglstErrlogDtbase = {
    prmkey: status.action,
    erstat: "",
    erpart: "",
    ersrce: "",
    erdtil: "",
    erdvsn: "",
    erignr: "",
    dateup: 0,
    timeup: 0,
    datefl: 0,
    airlfl: "",
    depart: "",
    flnbfl: "",
    Paxdif: "",
    flstat: "",
    flhour: 0,
    routfl: "",
    updtby: "",
    worker: 0,
  };

  // Variable default
  const [objdta, objdtaSet] = useState(rawobj);
  const [okeupd, okeupdSet] = useState<string>("");
  const [cxlupd, cxlupdSet] = useState<string>("");
  const rplprm = FncGlobalQuerysEdlink();

  // Refresh data
  const cxlref = useRef<NodeJS.Timeout | null>(null);
  const rfresh = async (objnow: MdlPsglstErrlogDtbase) => {
    if (cxlref.current) clearTimeout(cxlref.current);
    if (status.sbrapi == 0) {
      objdtaSet(objnow);
      if (objnow.worker == 0) {
        if (objnow.flnbfl == "") {
          objnow.worker = 3;
          if (objnow.depart == "") {
            objnow.worker = 5;
            if (objnow.airlfl == "") objnow.worker = 8;
          }
        }
      }

      // Start process data
      const rsp = ApiPsglstPrcessManual(objnow);
      setTimeout(() => {
        rplprm(["update_global"], String(Math.random()));
      }, 1000);
      if ((await rsp) == "Success") {
        okeupdSet(objnow.prmkey);
      } else cxlupdSet(objnow.prmkey);
      setTimeout(() => {
        (okeupdSet(""), cxlupdSet(""), objdtaSet(rawobj));
        rplprm(["update_global"], String(Math.random()));
      }, 2500);
    }

    // Reject other proses
    else {
      cxlupdSet(objnow.prmkey);
      cxlref.current = setTimeout(() => {
        cxlupdSet("");
      }, 1000);
    }
  };

  // Ignore
  const ignore = (objnow: MdlPsglstErrlogDtbase) => {
    objnow.erignr = objnow.prmkey;
    rfresh(objnow);
  };

  // Montitoring process
  const itvref = useRef<NodeJS.Timeout | null>(null);
  useEffect(() => {
    objdtaSet({ ...rawobj, prmkey: "all" });
    if (itvref.current) clearInterval(itvref.current);
    if (status.action != "") objdtaSet({ ...rawobj, prmkey: status.action });
    if (status.sbrapi != 0) {
      itvref.current = setInterval(async () => {
        const statnw = await ApiAllusrStatusPrcess();
        if (statnw.action == "") objdtaSet({ ...rawobj, prmkey: "all" });
        if (statnw.sbrapi == 0) {
          objdtaSet(rawobj);
          rplprm(["update_global"], String(Math.random()));
          clearInterval(itvref.current);
        }
      }, 5000);
    } else objdtaSet(rawobj);
  }, [update]);

  return (
    <>
      <div className="ctable">
        <table>
          <UixGlobalTheadxTablex firsth="action" mainhd={Object.keys(rawobj)} />
          <UixGlobalTbodyrTablex
            arrdta={errlog}
            objdta={objdta}
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
            rfresh={rfresh}
            ignore={ignore}
            tolink={null}
            editdt={null}
            trashx={null}
            okeupd={okeupd}
            cxlupd={cxlupd}
          />
        </table>
      </div>
    </>
  );
}
