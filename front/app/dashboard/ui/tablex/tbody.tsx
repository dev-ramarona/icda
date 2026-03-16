import { Dispatch, SetStateAction } from "react";
import { MdlApndixAcpedtDtbase } from "../../apndix/model/parmas";
import { FncGlobalFormatDatefm } from "../../global/function/format";
import UixGlobalInputxFormdt from "../../global/ui/client/inputx";
import { UixGlobalIconvcCancel, UixGlobalIconvcCeklis, UixGlobalIconvcEditdt } from "../../global/ui/server/iconvc";

export default function UixGlobalTbodyxTablex({
    actedt, arrobj, edtobj, edtset, acpedt, datefm, nmbrfm, cnfupd, okeupd, cxlupd }: {
        actedt: (e: React.ChangeEvent<HTMLInputElement>, actmod: "add" | "edt") => void,
        arrobj: any[],
        edtobj: any, edtset: Dispatch<SetStateAction<any>>,
        acpedt: MdlApndixAcpedtDtbase[]; datefm: string[], nmbrfm: string[]
        cnfupd: () => void, okeupd: string, cxlupd: string
    }) {
    const timeip = datefm.filter(v => v.includes("time"))
    const dateip = datefm.filter(v => v.includes("date"))
    return (
        <tbody>
            {arrobj.map((log, idx) => (
                <tr key={idx}>

                    {/* Button action */}
                    <td className={`text-center sticky left-0 z-10 w-20 drop-shadow-md
                        ${cxlupd === log.prmkey ? "bg-red-200 shkeit h-12"
                            : okeupd === log.prmkey ? "bg-green-200 shkeit h-12"
                                : edtobj?.prmkey === log.prmkey ? "bg-sky-200 h-12" : "bg-white"}`}>
                        <div className="afull flexctr gap-x-1.5 relative">
                            <div className={`flexctr btnsbm duration-300 cursor-pointer 
                            ${edtobj?.prmkey === log.prmkey ? "opacity-100"
                                    : "opacity-0 select-none pointer-events-none"}`}
                                onClick={() => cnfupd()}>
                                <UixGlobalIconvcCeklis bold={4} color="#53eafd" size={1.4} />
                            </div>
                            <div className={`flexctr btnsbm duration-300 cursor-pointer 
                            ${edtobj?.prmkey === log.prmkey ? "opacity-100"
                                    : "opacity-0 select-none pointer-events-none"}`}
                                onClick={() => edtset({ ...log, prmkey: "" })}>
                                <UixGlobalIconvcCancel bold={4} color="#fb2c36" size={1.4} />
                            </div>
                            <div
                                className={`absolute flexctr btnsbm duration-300 cursor-pointer 
                                    ${edtobj?.prmkey === log.prmkey ? "opacity-0 select-none pointer-events-none"
                                        : "opacity-100"}`}
                                onClick={() => edtset({ ...log, prmkey: log.prmkey })}>
                                <UixGlobalIconvcEditdt bold={2.7} color="white" size={1.4} />
                            </div>
                        </div>
                    </td>

                    {/* Format edit or not edit */}
                    {(typeof log === "object" && log !== null) && Object.entries(log as object).map(([key, val]) => (
                        <td className={`text-center z-0
                            ${cxlupd === log.prmkey ? "bg-red-200 shkeit h-12"
                                : okeupd === log.prmkey ? "bg-green-200 shkeit h-12"
                                    : edtobj?.prmkey === log.prmkey ? "bg-sky-200 h-12"
                                        : "bg-white h-0"} duration-300`} key={key}>
                            {edtobj?.prmkey === log.prmkey && acpedt.some((item) => item.params === key) ?
                                (
                                    <div className="relative flexctr h-0">
                                        <div className="min-w-20">
                                            <UixGlobalInputxFormdt
                                                typipt={timeip.includes(key) ? "datetime-local"
                                                    : dateip.includes(key) ? "date"
                                                        : "text"}
                                                length={acpedt.find((item) => item.params === key)?.length}
                                                queryx={key.toString()}
                                                params={String(edtobj[key])}
                                                plchdr={""}
                                                repprm={(e) => actedt(e, "edt")}
                                                labelx=""
                                            />
                                        </div>
                                    </div>
                                ) : (
                                    <div>
                                        {datefm.includes(key) ? FncGlobalFormatDatefm(String(val))
                                            : nmbrfm.includes(key) ? <div className="text-right">{val.toLocaleString("en-US")}</div>
                                                : val}
                                    </div>
                                )
                            }
                        </td>
                    ))}
                </tr>
            ))
            }
        </tbody >
    );
}