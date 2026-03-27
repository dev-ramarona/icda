"use client";

import { MdlPsglstPsgdtlFrntnd } from "../../model/params";
import { useRef, useState } from "react";
import { ApiPsglstPsgdtlUpdate } from "../../api/psgdtl";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatDefault, FncGlobalFormatInptdt } from "../../../global/function/format";
import { MdlApndixAcpedtDtbase } from "../../../apndix/model/parmas";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import UixGlobalTbodyxTablex from "../../../public/ui/tablex/tbodyx";
import UixGlobalTheadxTablex from "../../../public/ui/tablex/theadx";
import UixGlobalConfrmAction from "../../../public/ui/action/confrm";
import { MdlGlobalConfrmAction } from "../../../public/model/params";
import { FncPsglstRawdtaParams } from "../../function/params";

export default function UixPsglstDetailTablex({
  arrdta,
  acpedt,
  cookie,
}: {
  arrdta: MdlPsglstPsgdtlFrntnd[];
  acpedt: MdlApndixAcpedtDtbase[];
  cookie: mdlAllusrCookieObjson;
}) {
  // Dinamis
  const exclde = ["prmkey", "hstory", "updtby"];
  const inclde = acpedt.map((item) => item.params);
  const rawobj: MdlPsglstPsgdtlFrntnd = FncPsglstRawdtaParams();
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
    val = FncGlobalFormatDefault(key, val);
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
        </table>
      </div>
    </>
  );
}

//   const [edtobj, edtobjSet] = useState<MdlPsglstPsgdtlFrntnd>();
//   const [okeupd, okeupdSet] = useState<string>("");
//   const [cxlupd, cxlupdSet] = useState<string>("");
//   const [cxlrsp, cxlrspSet] = useState<string>("");
//   const rplprm = FncGlobalQuerysEdlink();
//   const actedt = (e: React.ChangeEvent<HTMLInputElement>) => {
//     const key = e.currentTarget.id;
//     let val: string | number = e.currentTarget.value;
//     if (key == "routvc") val = FncGlobalFormatRoutfl(val);
//     else if (key == "cpnbvc") val = FncGlobalFormatCpnfmt(val);
//     else if (["tktnbr", "flnbvc"].includes(key)) val = val.replace(/[^0-9]/g, "");
//     else if (["ntafvc", "ntaffl"].includes(key)) val = Number(val);
//     else val = val.toUpperCase();
//     edtobjSet({
//       ...edtobj,
//       [key]: val,
//     } as MdlPsglstPsgdtlFrntnd);
//   };

//   // Confirm update retail or series
//   const update = async (log: MdlPsglstPsgdtlFrntnd) => {
//     const rspupd: string = await ApiPsglstPsgdtlUpdate(log);
//     edtobjSet({ ...log, prmkey: "" });
//     if (rspupd == "success") {
//       okeupdSet(log.prmkey);
//     } else {
//       cxlupdSet(log.prmkey);
//       cxlrspSet(rspupd);
//     }
//     setTimeout(() => {
//       okeupdSet("");
//       cxlupdSet("");
//       cxlrspSet("");
//       rplprm(["update_global"], String(Math.random()));
//     }, 1000);
//   };

//   return (
//     <>
//       <div className="ctable">
//         <table>
//           <thead>
//             <tr>
//               <th className="sticky left-0">Action</th>
//               {arrdta && arrdta.length > 0
//                 ? Object.entries(arrdta[0]).map(([key]) => <th key={key}>{key}</th>)
//                 : ""}
//             </tr>
//           </thead>
//           <tbody>
//             {arrdta.map((log, idx) => (
//               <tr key={idx}>
//                 <td
//                   className={`sticky left-0 z-10 text-center drop-shadow-lg ${
//                     edtobj?.prmkey === log.prmkey
//                       ? "bg-sky-200"
//                       : okeupd === log.prmkey
//                         ? "shkeit bg-green-400"
//                         : cxlupd === log.prmkey
//                           ? "shkeit bg-red-400"
//                           : "bg-white"
//                   }`}
//                 >
//                   <div className="afull flexctr relative gap-x-1.5">
//                     <div
//                       className={`flexctr btnsbm cursor-pointer duration-300 ${
//                         edtobj?.prmkey === log.prmkey
//                           ? "opacity-100"
//                           : "pointer-events-none opacity-0 select-none"
//                       }`}
//                       onClick={() => update(edtobj as MdlPsglstPsgdtlFrntnd)}
//                     >
//                       <UixGlobalIconvcCeklis bold={2.5} color="#53eafd" size={1.4} />
//                     </div>
//                     <div
//                       className={`flexctr btnsbm cursor-pointer duration-300 ${
//                         edtobj?.prmkey === log.prmkey
//                           ? "opacity-100"
//                           : "pointer-events-none opacity-0 select-none"
//                       }`}
//                       onClick={() => edtobjSet({ ...log, prmkey: "" })}
//                     >
//                       <UixGlobalIconvcCancel bold={2.5} color="#fb2c36" size={1.4} />
//                     </div>
//                     <div
//                       className={`flexctr btnsbm absolute cursor-pointer duration-300 ${
//                         edtobj?.prmkey === log.prmkey
//                           ? "pointer-events-none opacity-0 select-none"
//                           : "opacity-100"
//                       }`}
//                       onClick={() =>
//                         edtobjSet({ ...log, updtby: cookie.usrnme, prmkey: log.prmkey })
//                       }
//                     >
//                       <UixGlobalIconvcEditdt bold={2.5} color="white" size={1.4} />
//                     </div>
//                   </div>
//                   <div
//                     className={`${
//                       cxlupd === log.prmkey ? "flexctr font-semibold text-white" : "h-0 opacity-0"
//                     } duration-300`}
//                   >
//                     {cxlrsp}
//                   </div>
//                 </td>
//                 {Object.entries(log).map(([key, val]) => (
//                   <td
//                     className={`z-0 w-fit text-center ${
//                       edtobj?.prmkey === log.prmkey
//                         ? "bg-sky-200"
//                         : okeupd === log.prmkey
//                           ? "shkeit bg-green-400"
//                           : cxlupd === log.prmkey
//                             ? "shkeit bg-red-400"
//                             : "bg-white"
//                     }`}
//                     key={key}
//                   >
//                     {edtobj?.prmkey === log.prmkey && acpedt.some((item) => item.params === key) ? (
//                       <div className="flexctr relative">
//                         <span className="invisible">
//                           XXXXXXXXXXXXX{String(edtobj[key as keyof typeof edtobj])}
//                         </span>
//                         <div className="absolute">
//                           <UixGlobalInputxFormdt
//                             typipt={
//                               key == "timeis" ? "datetime-local" : key == "datevc" ? "date" : "text"
//                             }
//                             length={acpedt.find((item) => item.params === key)?.length}
//                             queryx={key.toString()}
//                             params={String(edtobj[key as keyof typeof edtobj])}
//                             plchdr=""
//                             repprm={actedt}
//                             labelx=""
//                           />
//                         </div>
//                       </div>
//                     ) : (
//                       <div>
//                         {[
//                           "datefl",
//                           "daterv",
//                           "datevc",
//                           "timefl",
//                           "timevc",
//                           "timerv",
//                           "timeis",
//                           "timecr",
//                           "mnthfl",
//                         ].includes(key) ? (
//                           FncGlobalFormatDatefm(String(val))
//                         ) : ["ntaffl", "ntafvc", "yqtxfl", "yqtxvc", "qsrcrw", "qsrcvc"].includes(
//                             key,
//                           ) ? (
//                           <div className="text-right">{val.toLocaleString("en-US")}</div>
//                         ) : (
//                           val
//                         )}
//                       </div>
//                     )}
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
