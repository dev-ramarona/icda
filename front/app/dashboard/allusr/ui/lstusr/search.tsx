"use client";
import { useEffect, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";
import { MdlAllusrSearchParams } from "../../model/params";


export default function UixCrtusrAllusrSearch({
    prmAllusr,
}: {
    prmAllusr: MdlAllusrSearchParams;
}) {
    const [params, paramsSet] = useState<MdlAllusrSearchParams>({
        usredt: prmAllusr.usredt || "",
        stfnme: prmAllusr.stfnme || "",
        usrnme: prmAllusr.usrnme || "",
        stfeml: prmAllusr.stfeml || "",
        limitp: prmAllusr.limitp || 1,
        pagenw: prmAllusr.pagenw || 15,
        update: prmAllusr.update || "",
    });

    // Monitor change
    const [chnged, chngedSet] = useState<boolean>(false);
    useEffect(() => {
        chngedSet(false);
        paramsSet({
            usredt: prmAllusr.usredt || "",
            stfnme: prmAllusr.stfnme || "",
            usrnme: prmAllusr.usrnme || "",
            stfeml: prmAllusr.stfeml || "",
            limitp: prmAllusr.limitp || 1,
            pagenw: prmAllusr.pagenw || 15,
            update: prmAllusr.update || "",
        });
    }, [prmAllusr]);

    // Replace params
    const rplprm = FncGlobalQuerysEdlink();
    const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
        chngedSet(true);
        const namefl = e.currentTarget.id;
        let valuef = e.currentTarget.value;
        paramsSet({
            ...params,
            [namefl]: valuef,
        });
        rplprm([namefl, "pagenw"], [valuef, ""]);
    };

    // Reset function
    const resetx = () => {
        chngedSet(true);
        rplprm(["stfnme", "usrnme", "stfeml", "limitp", "usredt"], "");
    };
    return (
        <div className="w-full h-24 min-h-fit py-3 flexctr relative">
            <div className={`${chnged ? "w-16 h-10 translate-y-0" : "w-0 h-0 opacity-0 -translate-y-10"} z-10 absolute bg-white ring-2 ring-sky-300 px-5 py-2 rounded-xl flexctr duration-300`}>
                <div>Wait</div>
                <div className="animate-spin"><UixGlobalIconvcRfresh bold={2} color="black" size={1} /></div>
            </div>
            <div className={`afull flexstr flex-wrap gap-y-3 ${chnged ? "animate-pulse select-none" : ""} duration-300`}>
                <div className="w-1/2 md:w-28 h-11 flexctr relative">
                    <UixGlobalInputxFormdt
                        typipt={"text"}
                        length={undefined}
                        queryx={"stfnme"}
                        params={params.stfnme}
                        plchdr="Staff Name"
                        repprm={repprm}
                        labelx=""
                    />
                </div>
                <div className="w-1/2 md:w-28 h-11 flexctr relative">
                    <UixGlobalInputxFormdt
                        typipt={"text"}
                        length={undefined}
                        queryx={"usrnme"}
                        params={params.usrnme}
                        plchdr="Username"
                        repprm={repprm}
                        labelx=""
                    />
                </div>
                <div className="w-1/2 md:w-28 h-11 flexctr relative">
                    <UixGlobalInputxFormdt
                        typipt={"text"}
                        length={undefined}
                        queryx={"stfeml"}
                        params={params.stfeml}
                        plchdr="Staff Email"
                        repprm={repprm}
                        labelx=""
                    />
                </div>
            </div>
            <div className={`w-1/3 flexend flex-wrap gap-3 px-3 ${chnged ? "animate-pulse select-none" : ""} duration-300`}>
                <div className="w-full md:w-28 h-10 flexctr relative">
                    <div className="afull btnwrn flexctr" onClick={() => resetx()}>
                        Reset
                    </div>
                </div>
            </div>
        </div>
    );
}
