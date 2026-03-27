"use client";

import UixGlobalTheadxTablex from "../../../public/ui/tablex/theadx";
import UixGlobalTbodyrTablex from "../../../public/ui/tablex/tbodyr";
import { MdlSlsflwPsgsmrFrntnd } from "../../model/params";

export default function UixSlsflwPsgsmrTablex({ psgsmr }: { psgsmr: MdlSlsflwPsgsmrFrntnd[] }) {
  // Dinamis
  const rawobj: MdlSlsflwPsgsmrFrntnd = {
    prmkey: "",
    airlfl: "",
    provnc: "",
    depart: "",
    flnbfl: "",
    routfl: "",
    ndayfl: "",
    datefl: 0,
    mnthfl: 0,
    flstat: "",
    seatcn: "",
    airtyp: "",
    flhour: 0,
    totnta: 0,
    tottyq: 0,
    totpax: 0,
    totfae: 0,
    totqfr: 0,
    totrph: 0,
  };

  return (
    <>
      <div className="ctable">
        <table>
          <UixGlobalTheadxTablex firsth="" mainhd={Object.keys(rawobj)} />
          <UixGlobalTbodyrTablex
            arrdta={psgsmr}
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
            nmbrfm={["ntaffl", "ntafvc", "yqtxfl", "yqtxvc", "qsrcrw", "qsrcvc"]}
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

//   const [edtobj, edtobjSet] = useState<MdlSlsflwPsgsmrFrntnd>();
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
//     } as MdlSlsflwPsgsmrFrntnd);
//   };

//   return (
//     <>
//       <div className="ctable">
//         <table>
//           <thead>
//             <tr>
//               {Psgsmr && Psgsmr.length > 0
//                 ? Object.entries(Psgsmr[0]).map(([key]) => <th key={key}>{key}</th>)
//                 : ""}
//             </tr>
//           </thead>
//           <tbody>
//             {Psgsmr.map((log, idx) => (
//               <tr key={idx}>
//                 {Object.entries(log).map(([key, val]) => (
//                   <td className={`z-0 w-fit text-center`} key={key}>
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
//                         {["datefl", "mnthfl"].includes(key) ? (
//                           FncGlobalFormatDatefm(String(val))
//                         ) : ["totnta", "tottyq", "totpax", "totfae", "totqfr"].includes(key) ? (
//                           <div className="w-full text-right">{val.toLocaleString("en-US")}</div>
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
