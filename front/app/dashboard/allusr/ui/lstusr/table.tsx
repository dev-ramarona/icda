'use client'

import { useState } from "react";
import { FncGlobalQuerysEdlink } from "../../../global/function/querys";
import { UixGlobalIconvcCancel, UixGlobalIconvcDelete, UixGlobalIconvcEditdt } from "../../../global/ui/server/iconvc";
import { ApiAllusrHandleDelete } from "../../api/user";
import { MdlAllusrFrntndParams } from "../../model/params";

export default function UixCrtusrAllusrTablex({
    allusr,
}: {
    allusr: MdlAllusrFrntndParams[];
}) {

    // Hit the database and get interval status
    const rplprm = FncGlobalQuerysEdlink();
    const [usrnme, usrnmeSet] = useState("");

    // Delete user
    const delusr = async (usrnme: string) => {
        const res = await ApiAllusrHandleDelete(usrnme);
        if (res) {
            usrnmeSet("");
            rplprm(["update"], String(Math.random()));
        } else {
        }
    }

    // Const edit user
    const edtusr = (params: MdlAllusrFrntndParams) => {
        const usredt = JSON.stringify({ params })
        rplprm(["usredt"], usredt);
    }

    return (
        <>
            <div className={`absolute flexctr ${usrnme ? "afull opacity-100" : "w-0 h-0 opacity-0 select-none"} duration-300`}>
                <div className={`bg-white ring ring-gray-300 rounded-lg z-30 flexctr p-3 flexctr flex-col text-base text-gray-600 
                ${usrnme ? "" : "opacity-0 select-none pointer-events-none overflow-hidden -translate-y-full"} shadow-md duration-300`}>
                    <div>Confirm <span className="text-red-500 font-semibold">Delete</span> this data?</div>
                    <div>Username:<span className="text-green-700 px-1.5">{usrnme}</span></div>
                    <div className="flexctr gap-1.5 pt-3">
                        <div className="w-10 flexctr btncxl duration-300 cursor-pointer p-1" onClick={() => delusr(usrnme)}>
                            <UixGlobalIconvcDelete
                                bold={3}
                                color="#fff"
                                size={1.2}
                            />
                        </div>
                        <div className="w-10 flexctr btnsbm duration-300 cursor-pointer p-1" onClick={() => usrnmeSet("")}>
                            <UixGlobalIconvcCancel
                                bold={3}
                                color="#fff"
                                size={1.2}
                            />
                        </div>
                    </div>
                </div>
            </div>
            <div className="ctable">
                <table>
                    <thead>
                        <tr>
                            <th className="sticky left-0">Action</th>
                            {allusr && allusr.length > 0
                                ? Object.entries(allusr[0]).map(([key]) => (
                                    <th key={key}>
                                        {key}
                                    </th>
                                ))
                                : ""}
                        </tr>
                    </thead>
                    <tbody>
                        {allusr.map((log, idx) => (
                            <tr className={`${usrnme == log.usrnme ? "bg-cyan-100" : ""}`} key={idx}>
                                <td className="text-center sticky left-0 z-10 drop-shadow-lg bg-white">
                                    <div className="afull flexctr gap-x-1.5">
                                        <div className="w-1/2 flexctr btnwrn duration-300 cursor-pointer p-1"
                                            onClick={() => edtusr(log)}
                                        >
                                            <UixGlobalIconvcEditdt
                                                bold={3}
                                                color="#fff"
                                                size={1.2}
                                            />
                                        </div>
                                        <div className="w-1/2 flexctr btncxl duration-300 cursor-pointer p-1"
                                            onClick={() => usrnmeSet(log.usrnme)}
                                        >
                                            <UixGlobalIconvcDelete
                                                bold={3}
                                                color="#fff"
                                                size={1.2}
                                            />
                                        </div>
                                    </div>
                                </td>
                                {Object.entries(log).map(([key, val]) => (
                                    <td className="text-center" key={key}>
                                        {["access", "keywrd"].includes(key)
                                            ? (
                                                val ?
                                                    <div className="flex flex-wrap justify-center gap-1">
                                                        {val.map((item: string, idx: number) => (
                                                            <div key={idx} className="bg-gray-200 p-1.5 rounded-md font-semibold">
                                                                {item}
                                                            </div>
                                                        ))}
                                                    </div>
                                                    : ""
                                            )
                                            : val}
                                    </td>
                                ))}
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </>
    );
}
