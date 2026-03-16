'use client'
import { useEffect, useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";

export default function UixApndixApplstSelect({ apndix, pagedb }: { apndix: string[], pagedb: string }) {
    const rplprm = FncGlobalQuerysEdlink();
    const [pageon, pageonSet] = useState("");
    const select = (val: string) => {
        pageonSet(val)
        rplprm(["update", "pagedb", "datefl", "airlfl",
            "depart", "flnbfl", "routfl", "clssfl", "pagenw",
        ], [String(Math.random()), val, "", "", "", "", "", "", "1"])
    }
    useEffect(() => {
        if (pagedb == "") {
            pageonSet(apndix[0])
        } else pageonSet(pagedb)
    }, [pagedb])
    return (
        <div className="w-full h-10 min-h-fit py-3 flexctr relative">
            <div className={`afull flexstr flex-wrap gap-3`}>
                {apndix.map((val, idx) => (
                    <div className={`px-3 py-1.5 w-fit flexctr text-center ${val == pageon ? "btnsbm" : "btnstb"} duration-300`} key={idx} onClick={() => select(val)}>
                        <div>{val}</div>
                        <div className={`flexctr ${val == pageon ? "w-8 pl-1.5" : "w-1 opacity-0"} duration-300`}>
                            <div className="absolute w-6 rounded-full h-1 bg-white rotate-45"></div>
                            <div className="w-6 h-1.5"></div>
                            <div className="absolute w-6 rounded-full h-1 bg-white -rotate-45"></div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}