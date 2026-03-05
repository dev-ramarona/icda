'use client'

import { useActionState, useEffect, useState } from "react";
import { apiLoginpFormpgLoginx } from "../api/login";

export default function UixLoginpInputdFormdt() {
    const [formac, formacSet, isPending] = useActionState(apiLoginpFormpgLoginx, null);
    const [formdt, formdtSet] = useState({ usrnme: "", psswrd: "" });
    const [rspnse, rspnseSet] = useState({
        dfault: formac?.dfault || "",
        rspnse: formac?.rspnse || "",
    });

    // Monitor
    useEffect(() => {
        rspnseSet((prev) => ({
            ...prev,
            dfault: formac?.dfault || "",
            rspnse: formac?.rspnse || "",
        }));
    }, [formac]);


    // Function onchange
    const onchng = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { id, value } = e.target;
        formdtSet((prev) => ({
            ...prev,
            [id as keyof typeof prev]: value,
        }));
    };
    return (
        <form
            className="w-80 md:w-96 min-w-fit h-80 min-h-fit"
            action={formacSet}
        >
            <div className="afull flexctr flex-col bg-linear-to-br from-zinc-100 via-white to-cyan-100 rounded-xl ring ring-gray-200">
                <div className="w-3/4 py-1.5 flexctr flex-col">
                    <div className="text-gray-700 font-black text-2xl">LOGIN</div>
                    <div className="font-semibold">Data Analyst IC</div>
                </div>
                <div className="w-1/2 py-1.5 flexctr flex-col relative">
                    <label
                        className={`px-1.5 afull font-semibold text-gray-500/50 relative flexstr`}
                        htmlFor="usrnme"
                    >
                        <div className="opacity-0">Username</div>
                        <div
                            className={`absolute opacity-100 cursor-text ${formdt.usrnme.length > 0
                                ? "translate-y-0 pt-0 pb-1"
                                : "translate-y-full pt-1 pb-0"
                                } duration-300`}
                        >
                            Username
                        </div>
                    </label>
                    <input
                        className="afull bg-white text-slate-800 p-1.5 rounded-md ring-2 ring-gray-100"
                        defaultValue={rspnse.dfault}
                        type="text"
                        name="usrnme"
                        id="usrnme"
                        onChange={(e) => onchng(e)}
                    />
                </div>
                <div className="w-1/2 py-1.5 flexctr flex-col relative">
                    <label
                        className={`px-1.5 afull font-semibold text-gray-500/50 relative flexstr`}
                        htmlFor="psswrd"
                    >
                        <div className="opacity-0">Password</div>
                        <div
                            className={`absolute opacity-100 cursor-text ${formdt.psswrd.length > 0
                                ? "translate-y-0 pt-0 pb-1"
                                : "translate-y-full pt-1 pb-0"
                                } duration-300`}
                        >
                            Password
                        </div>
                    </label>
                    <input
                        className="afull bg-white text-slate-800 p-1.5 rounded-md ring-2 ring-gray-100"
                        type="password"
                        name="psswrd"
                        id="psswrd"
                        onChange={(e) => onchng(e)}
                    />
                </div>
                <button
                    className="w-1/2 h-16 py-3 flexctr cursor-pointer group relative"
                    type="submit"
                    disabled={isPending}
                >
                    <div className="afull btnsbm flexctr">{isPending ? "Loading..." : "Login"}</div>
                    <div className="w-full absolute -bottom-2.5 flexctr text-red-500 text-[0.65rem] px-1.5 pt-0.5">
                        {rspnse.rspnse}
                    </div>
                </button>
            </div>
        </form>
    );
}