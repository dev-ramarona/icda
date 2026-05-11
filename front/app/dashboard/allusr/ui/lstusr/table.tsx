"use client";

import { useRef, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { ApiAllusrHandleDelete } from "../../api/user";
import { MdlAllusrFrntndParams } from "../../model/params";
import UixGlobalTheadxTablex from "../../../global/ui/tablex/theadx";
import UixGlobalTbodyrTablex from "../../../global/ui/tablex/tbodyr";
import UixGlobalConfrmAction from "../../../global/ui/action/confrm";
import { MdlGlobalConfrmAction } from "../../../global/model/params";

export default function ApiAllusrUsrlstTablex({ allusr }: { allusr: MdlAllusrFrntndParams[] }) {
  // Hit the database and get interval status
  const rplprm = FncGlobalQuerysEdlink();
  const [usrnme, usrnmeSet] = useState("");

  // Delete user
  const [confrm, confrmSet] = useState<boolean>(false);
  const [confdt, confdtSet] = useState<MdlGlobalConfrmAction[]>([]);
  const confdl = async (usrnme: MdlAllusrFrntndParams) => {
    usrnmeSet(usrnme.prmkey);
    const confst = [];
    Object.entries(usrnme).map(([k, v]) => confst.push({ paramx: k, valuex: v }));
    confrmSet(true);
    confdtSet(confst);
  };
  const godelt = async () => {
    const res = await ApiAllusrHandleDelete(usrnme);
    if (res) {
      usrnmeSet("");
      confrmSet(false);
      rplprm(["update"], String(Math.random()));
    } else {
    }
  };

  // Const edit user
  const edtusr = (params: MdlAllusrFrntndParams) => {
    const usredt = JSON.stringify({ params });
    rplprm(["usredt"], usredt);
  };

  return (
    <>
      <UixGlobalConfrmAction
        confrm={confrm}
        confdt={confdt}
        action={"update"}
        goupdt={godelt}
        confrmSet={confrmSet}
      />
      <div className="ctable">
        <table>
          <UixGlobalTheadxTablex firsth="action" mainhd={Object.keys(allusr[0])} />
          <UixGlobalTbodyrTablex
            arrdta={allusr}
            objdta={usrnme}
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
            spclfm={["access", "keywrd"]}
            rfresh={null}
            ignore={null}
            tolink={null}
            editdt={edtusr}
            trashx={confdl}
            okeupd={null}
            cxlupd={null}
          />
        </table>
      </div>
    </>
    // <>
    //   <div
    //     className={`flexctr absolute duration-300 ${
    //       usrnme ? "afull opacity-100" : "h-0 w-0 opacity-0 select-none"
    //     } `}
    //   >
    //     <div
    //       className={`flexctr z-30 flex-col rounded-lg bg-white p-3 text-base text-gray-600 shadow-md ring ring-gray-300 ${
    //         usrnme && "pointer-events-none -translate-y-full overflow-hidden opacity-0 select-none"
    //       } duration-300`}
    //     >
    //       <div>
    //         Confirm <span className="font-semibold text-red-500">Delete</span> this data?
    //       </div>
    //       <div>
    //         Username:<span className="px-1.5 text-green-700">{usrnme}</span>
    //       </div>
    //       <div className="flexctr gap-1.5 pt-3">
    //         <div
    //           className="flexctr btncxl w-10 cursor-pointer p-1 duration-300"
    //           onClick={() => confdl(usrnme)}
    //         >
    //           <UixGlobalIconvcTrashx bold={3} color="#fff" size={1.2} />
    //         </div>
    //         <div
    //           className="flexctr btnsbm w-10 cursor-pointer p-1 duration-300"
    //           onClick={() => usrnmeSet("")}
    //         >
    //           <UixGlobalIconvcCancel bold={3} color="#fff" size={1.2} />
    //         </div>
    //       </div>
    //     </div>
    //   </div>
    //   <div className="ctable">
    //     <table>
    //       <thead>
    //         <tr>
    //           <th className="sticky left-0">Action</th>
    //           {allusr && allusr.length > 0
    //             ? Object.entries(allusr[0]).map(([key]) => <th key={key}>{key}</th>)
    //             : ""}
    //         </tr>
    //       </thead>
    //       <tbody>
    //         {allusr.map((log, idx) => (
    //           <tr className={`${usrnme == log.usrnme ? "bg-cyan-100" : ""}`} key={idx}>
    //             <td className="sticky left-0 z-10 bg-white text-center drop-shadow-lg">
    //               <div className="afull flexctr gap-x-1.5">
    //                 <div
    //                   className="flexctr btnscs w-1/2 cursor-pointer p-1 duration-300"
    //                   onClick={() => edtusr(log)}
    //                 >
    //                   <UixGlobalIconvcEditdt bold={3} color="#fff" size={1.2} />
    //                 </div>
    //                 <div
    //                   className="flexctr btnwrn w-1/2 cursor-pointer p-1 duration-300"
    //                   onClick={() => usrnmeSet(log.usrnme)}
    //                 >
    //                   <UixGlobalIconvcTrashx bold={3} color="#fff" size={1.2} />
    //                 </div>
    //               </div>
    //             </td>
    //             {Object.entries(log).map(([key, val]) => (
    //               <td className="text-center" key={key}>
    //                 {["access", "keywrd"].includes(key) ? (
    //                   val ? (
    //                     <div className="flex flex-wrap justify-center gap-1">
    //                       {val.map((item: string, idx: number) => (
    //                         <div key={idx} className="rounded-md bg-gray-200 p-1.5 font-semibold">
    //                           {item}
    //                         </div>
    //                       ))}
    //                     </div>
    //                   ) : (
    //                     ""
    //                   )
    //                 ) : (
    //                   val
    //                 )}
    //               </td>
    //             ))}
    //           </tr>
    //         ))}
    //       </tbody>
    //     </table>
    //   </div>
    // </>
  );
}
