import { Dispatch, SetStateAction } from "react";
import UixGlobalInputxFormdt from "../../global/ui/client/inputx";
import { UixGlobalIconvcAddpls, UixGlobalIconvcCancel, UixGlobalIconvcCeklis } from "../../global/ui/server/iconvc";

export default function UixGlobalTfootxTablex({
    actadd, defobj, edtobj, edtset, exclde, datefm, cnfupd, okeupd, cxlupd }: {
        actadd: (e: React.ChangeEvent<HTMLInputElement>, actmod: "add" | "add") => void,
        defobj: any, edtobj: any, edtset: Dispatch<SetStateAction<any>>,
        exclde: string[], datefm: string[], cnfupd: () => void, okeupd: string, cxlupd: string
    }) {
    const timeip = datefm.filter(v => v.includes("time"))
    const dateip = datefm.filter(v => v.includes("date"))
    return (
        <tfoot>
            <tr className="sticky bottom-0 z-20">
                <td className={`text-center sticky left-0 z-10 w-20 drop-shadow-md 
                    ${cxlupd === "add" ? "bg-red-200 shkeit"
                        : okeupd === "add" ? "bg-green-200 shkeit"
                            : edtobj?.prmkey === "add" ? "bg-sky-200" : "bg-white"}`}>
                    <div className="afull flexctr gap-x-1.5 relative">
                        <div className={`flexctr btnsbm duration-300 cursor-pointer 
                            ${edtobj?.prmkey === "add" ? "opacity-100"
                                : "opacity-0 select-none pointer-events-none"}`}
                            onClick={() => cnfupd()}>
                            <UixGlobalIconvcCeklis bold={4} color="#53eafd" size={1.4} />
                        </div>
                        <div className={`flexctr btnsbm duration-300 cursor-pointer 
                            ${edtobj?.prmkey === "add" ? "opacity-100"
                                : "opacity-0 select-none pointer-events-none"}`}
                            onClick={() => edtset({ ...defobj, prmkey: "" })}>
                            <UixGlobalIconvcCancel bold={4} color="#fb2c36" size={1.4} />
                        </div>
                        <div className={`absolute flexctr btnsbm duration-300 cursor-pointer 
                            ${edtobj?.prmkey === "add" ? "opacity-0 select-none pointer-events-none"
                                : "opacity-100"}`}
                            onClick={() => edtset({ ...defobj, prmkey: "add" })}>
                            <UixGlobalIconvcAddpls bold={2.7} color="white" size={1.4} />
                        </div>
                    </div>
                </td>
                {Object.keys(defobj).map((val, key) => (
                    <td className={`text-center z-0
                        ${cxlupd === "add" ? "bg-red-200 shkeit h-12"
                            : okeupd === "add" ? "bg-green-200 shkeit h-12"
                                : edtobj?.prmkey === "add" ? "bg-sky-200 h-12"
                                    : "bg-white h-0"} duration-300`} key={key}>
                        <div className="relative flexctr">
                            {edtobj.prmkey == "add" && !exclde.includes(val) ?
                                (
                                    <div className="relative flexctr h-0">
                                        <div className="min-w-20">
                                            <UixGlobalInputxFormdt
                                                typipt={timeip.includes(val) ? "datetime-local"
                                                    : dateip.includes(val) ? "date"
                                                        : "text"}
                                                length={24}
                                                queryx={val.toString()}
                                                params={String(edtobj[val])}
                                                plchdr={""}
                                                repprm={(e) => actadd(e, "add")}
                                                labelx=""
                                            />
                                        </div>
                                    </div>
                                ) : (
                                    <div>-</div>
                                )
                            }
                        </div>
                    </td>
                ))}
            </tr>
        </tfoot>
    );
}