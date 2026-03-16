'use client'
import { useEffect, useState } from "react";
import { ApiApndixGetallDtbase } from "../../api/dtbase";
import { MdlApndixAcpedtDtbase, MdlApndixFlhourDtbase, MdlApndixSearchQueryx } from "../../model/parmas";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/client/pagntn";
import UixPsglstFlhourTablex from "./table";
import UixApndixFlhourSearch from "./search";

export default function UixApndixFlhourMainpg({
    qryprm,
    cookie,
    acpedt,
}: {
    qryprm: MdlApndixSearchQueryx;
    cookie: mdlAllusrCookieObjson;
    acpedt: MdlApndixAcpedtDtbase[];
}) {
    // const objdta = await ApiApndixGetallDtbase(qryprm)
    const [objdta, objdtaSet] = useState({ arrdta: [], totdta: 1 })
    const [isload, isloadSet] = useState(true)
    const arrdta: MdlApndixFlhourDtbase[] = objdta.arrdta
    const totdta: number = objdta.totdta

    useEffect(() => {
        const getall = async () => {
            isloadSet(true)
            const objdta = await ApiApndixGetallDtbase(qryprm)
            objdtaSet(objdta)
            isloadSet(false)
        }
        getall()
    }, [qryprm])

    return (
        <>
            <UixApndixFlhourSearch qryprm={qryprm} datefl={[]} />
            <UixPsglstFlhourTablex acpedt={acpedt} detail={arrdta} pagedb={qryprm.pagedb} cookie={cookie} isload={isload} />
            <UixGlobalPagntnMainpg
                pgview={15}
                pgenbr={qryprm.pagenw}
                pgestr="pagenw"
                totdta={totdta}
            />
        </>
    )
}