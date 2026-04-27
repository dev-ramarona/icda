"use client";
import { useEffect, useRef, useState } from "react";
import { MdlPsglstErrlogDtbase } from "../../../psglst/model/params";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { ApiPsglstPrcessManual } from "../../../psglst/api/prcess";
import { MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { ApiAllusrStatusPrcess } from "../../../allusr/api/status";
import UixGlobalTheadxTablex from "../../../global/ui/tablex/theadx";
import UixGlobalTbodyrTablex from "../../../global/ui/tablex/tbodyr";

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
  const [random, randomSet] = useState("");
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
  const tolink = (objnow: MdlPsglstErrlogDtbase) => {
    window.open(
      `/dashboard/apndix?pagedb_apndix=${objnow.erpart}&update_apndix=${random}`,
      "_blank",
    );
  };

  // Montitoring process
  const itvref = useRef<NodeJS.Timeout | null>(null);
  useEffect(() => {
    randomSet(String(Math.random()));
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
            ignore={null}
            tolink={tolink}
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

//   // Hit the database and get interval status
//   const rplprm = FncGlobalQuerysEdlink();
//   const [statfn, statfnSet] = useState("Done");
//   const [cxlupd, cxlupdSet] = useState<string>("");

//   // Process function
//   const prcess = async (params: MdlPsglstErrlogDtbase) => {
//     rplprm(["update_global"], String(Math.random()));
//     statfnSet("Wait");
//     const nowprm = { ...params };
//     if (status.sbrapi == 0) {
//       if (params.flnbfl == "") {
//         nowprm.worker = 3;
//         if (params.depart == "") {
//           nowprm.worker = 5;
//           if (params.airlfl == "")
//             nowprm.worker = 8;
//         }
//       }

//       // Set interval to check status
//       const rsp = ApiPsglstPrcessManual(nowprm);
//       setTimeout(() => rplprm(["update_global"], String(Math.random())), 1000);
//       statfnSet(await rsp);
//       if (await rsp == "Failed") cxlupdSet(params.prmkey)
//       setTimeout(() => { statfnSet(""); cxlupdSet(""); }, 2000);
//     }
//   }

//   // Monitor process status
//   useEffect(() => {
//     if (status.sbrapi == 0) statfnSet("");
//     const gtstat = async () => {
//       if (status.sbrapi != 0) {
//         const intrvl = setInterval(async () => {
//           console.log("action interval");
//           const instat = await ApiAllusrStatusPrcess();
//           if (instat.sbrapi == 0) {
//             statfnSet("");
//             rplprm(["update_global"], String(Math.random()));
//             clearInterval(intrvl);
//           } else statfnSet(`${instat.sbrapi}%`);
//         }, 2000);
//       }
//     };
//     gtstat();
//   }, [update]);

//   return (
//     <>
//       <div className="ctable">
//         <table>
//           <thead>
//             <tr>
//               <th className="sticky left-0">Action</th>
//               {errlog && errlog.length > 0
//                 ? Object.entries(errlog[0]).map(([key]) => (
//                   <th key={key}>
//                     {key}
//                   </th>
//                 ))
//                 : ""}
//             </tr>
//           </thead>
//           <tbody>
//             {errlog.map((log, idx) => (
//               <tr key={idx}>
//                 <td className={`text-center sticky left-0 z-10 drop-shadow-lg
//                   ${cxlupd == log.prmkey ? "bg-red-300 shkeit" : "bg-white"}`}>
//                   <div className="afull flexctr gap-x-1.5">
//                     <div className="w-1/2 flexctr btnsbm duration-300 cursor-pointer"
//                       onClick={() => prcess(log)}>
//                       <div className={`absolute text-gray-300 font-bold text-xs z-10`}>
//                         {statfn.includes("%") ? statfn : ""}
//                       </div>
//                       <div className={`${statfn != "" ? "animate-spin" : ""}`}>
//                         <UixGlobalIconvcRfresh
//                           bold={3}
//                           color="#fff"
//                           size={1.4}
//                         />
//                       </div>
//                     </div>
//                     <div className="w-1/2 flexctr btncxl duration-300 cursor-pointer"
//                       onClick={() => prcess({ ...log, erignr: log.prmkey })}>
//                       <UixGlobalIconvcIgnore
//                         bold={3}
//                         color="#fff"
//                         size={1.4}
//                       />
//                     </div>
//                     <div className="w-1/2 flexctr btnwrn duration-300 cursor-pointer"
//                       onClick={() => prcess({ ...log, erignr: log.prmkey })}>
//                       <UixGlobalIconvcTolink
//                         bold={3}
//                         color="#fff"
//                         size={1.4}
//                       />
//                     </div>
//                   </div>
//                 </td>
//                 {Object.entries(log).map(([key, val]) => (
//                   <td className={`text-center ${cxlupd == log.prmkey ? "bg-red-300 shkeit" : "bg-white"}`} key={key}>
//                     {["datefl", "timeup"].includes(key)
//                       ? FncGlobalFormatDatefm(String(val))
//                       : val}
//                   </td>
//                 ))}
//               </tr>
//             ))}
//           </tbody>
//         </table>
//       </div>
//     </>
//   );
// }
