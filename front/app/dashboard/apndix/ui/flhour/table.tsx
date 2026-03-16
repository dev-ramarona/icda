"use client";

import { act, useEffect, useRef, useState } from "react";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import { MdlApndixAcpedtDtbase, MdlApndixFlhourDtbase } from "../../model/parmas";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { FncGlobalFormatDefault } from "../../../global/function/format";
import { ApiApndixUpdateDtbase } from "../../api/dtbase";
import { UixGlobalIconvcCancel, UixGlobalIconvcCeklis, UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import UixGlobalTheadxTablex from "../../../ui/tablex/thead";
import UixGlobalTbodyxTablex from "../../../ui/tablex/tbody";
import UixGlobalTfootxTablex from "../../../ui/tablex/tfoot";



export default function UixPsglstFlhourTablex({
    detail,
    pagedb,
    acpedt,
    cookie,
    isload,
}: {
    detail: MdlApndixFlhourDtbase[];
    pagedb: string
    acpedt: MdlApndixAcpedtDtbase[];
    cookie: mdlAllusrCookieObjson;
    isload: boolean
}) {

    // Dinamis
    const exclde = ["prmkey", "hstory", "updtby"]
    const rawobj: MdlApndixFlhourDtbase = {
        prmkey: "", airlfl: "", routfl: "", flnbfl: "", flhour: 0,
        timefl: 0, timerv: 0, timeup: 0, dateup: 0, datend: 0,
        airtyp: "", airmls: 0, hstory: "", updtby: ""
    }

    // Variable default
    const [edtobj, edtobjSet] = useState(rawobj);
    const [arrdtl, arrdtlSet] = useState([{ ...rawobj, prmkey: "-" }]);
    const [okeupd, okeupdSet] = useState<string>("");
    const [cxlupd, cxlupdSet] = useState<string>("");
    const [confrm, confrmSet] = useState<boolean>(false);
    const [confdt, confdtSet] = useState<{ idx: string, str: string }[]>([])

    // edit params
    const rplprm = FncGlobalQuerysEdlink();
    const actedt = (e: React.ChangeEvent<HTMLInputElement>) => {
        const key = e.currentTarget.id;
        let val: string | number = e.currentTarget.value;
        val = FncGlobalFormatDefault(key, val);
        edtobjSet({ ...edtobj, [key]: val, });
    };

    // 
    const refedt = useRef<NodeJS.Timeout | null>(null)
    const cnfupd = () => {
        if (refedt.current) clearTimeout(refedt.current)
        const confst = []
        let emptys = false
        Object.entries(edtobj).map(([k, v]) => {
            if (!exclde.includes(k)) {
                confst.push({ idx: k, str: v })
                if ((v == "" || v == 0)) {
                    cxlupdSet(edtobj.prmkey)
                    emptys = true
                }
            }
        })
        if (!emptys) {
            confrmSet(true)
            confdtSet(confst)
        } else refedt.current = setTimeout(() => cxlupdSet(""), 1000);

    }

    // Confirm update retail or series
    const refupd = useRef<NodeJS.Timeout | null>(null)
    const goupdt = async () => {
        confrmSet(false)
        if (refupd.current) clearTimeout(refupd.current)
        refupd.current = setTimeout(async () => {
            edtobj.updtby = cookie.usrnme
            const rspupd: string = await ApiApndixUpdateDtbase(edtobj, pagedb);
            edtobjSet({ ...edtobj, prmkey: "" })
            if (rspupd == "success") {
                okeupdSet(edtobj.prmkey);
            } else cxlupdSet(edtobj.prmkey);
            setTimeout(() => {
                okeupdSet("");
                cxlupdSet("");
                rplprm(["update"], String(Math.random()));
            }, 1000);
        }, 1000)
    }

    // Monitoring detail data
    useEffect(() => {
        if (isload && rawobj.prmkey != "-") {
            // arrdtlSet([{ ...rawobj, prmkey: "-" }])
        } else arrdtlSet(detail)
    }, [isload])

    return (
        <>
            <div className="w-full h-full flexctr absolute pointer-events-none">
                <div className={`${isload ? "w-16 h-10 translate-y-0" : "w-0 h-0 opacity-0 -translate-y-10"} z-10 absolute bg-white ring-2 ring-sky-300 px-5 py-2 rounded-xl flexctr duration-300`}>
                    <div>Wait</div>
                    <div className="animate-spin"><UixGlobalIconvcRfresh bold={2} color="black" size={1} /></div>
                </div>
            </div>
            <div className={`w-full h-full flexctr absolute z-30 ${confrm ? "backdrop-blur-[2px]" : "pointer-events-none"}`}>
                <div className={`bg-white ring-2 ring-gray-200 text-gray-700 rounded-xl flexctr flex-col gap-y-6
                    ${confrm ? "w-60 h-60 translate-y-0 p-3"
                        : "w-0 h-0 opacity-0 -translate-y-10"} overflow-hidden duration-300`}>
                    <div className="text-center text-base text-nowrap">Confirm process this data</div>
                    <div className="overflow-auto">
                        {confdt.map((val, idx) => (
                            <div className="flexctr gap-1.5" key={idx}>
                                <div className="w-20 whitespace-nowrap font-medium">{val.idx}</div>
                                <div className="w-3 flexctr">:</div>
                                <div className="w-full whitespace-nowrap">{val.str}</div>
                            </div>
                        ))}
                    </div>
                    <div className="flexctr gap-3">
                        <div className="w-10 h-8 btnsbm flexctr" onClick={() => goupdt()}>
                            <UixGlobalIconvcCeklis bold={4} color="#53eafd" size={1.4} />
                        </div>
                        <div className="w-10 h-8 btnsbm flexctr" onClick={() => confrmSet(false)}>
                            <UixGlobalIconvcCancel bold={4} color="#fb2c36" size={1.4} />
                        </div>
                    </div>
                </div>
            </div>
            <div className={`ctable ${isload ? "animate-pulse" : ""}`}>
                <table>
                    <UixGlobalTheadxTablex firsth="action" mainhd={Object.keys(rawobj)} />
                    <UixGlobalTbodyxTablex
                        actedt={actedt}
                        arrobj={arrdtl}
                        edtobj={edtobj}
                        edtset={edtobjSet}
                        acpedt={acpedt}
                        datefm={["datefl", "daterv", "datevc", "dateup", "datend", "timefl",
                            "timevc", "timerv", "timeis", "timecr", "timeup", "mnthfl"]}
                        nmbrfm={["ntaffl", "ntafvc", "yqtxfl", "yqtxvc", "qsrcrw", "qsrcvc"]}
                        cnfupd={cnfupd}
                        okeupd={okeupd}
                        cxlupd={cxlupd}
                    />
                    <UixGlobalTfootxTablex
                        actadd={actedt}
                        defobj={rawobj}
                        edtobj={edtobj}
                        edtset={edtobjSet}
                        exclde={exclde}
                        datefm={["datefl", "daterv", "datevc", "dateup", "datend", "timefl",
                            "timevc", "timerv", "timeis", "timecr", "timeup", "mnthfl"]}
                        cnfupd={cnfupd}
                        okeupd={okeupd}
                        cxlupd={cxlupd}
                    />
                </table>
            </div>
        </>
    );
}
