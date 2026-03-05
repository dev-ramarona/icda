import { Suspense } from "react";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import { FncAllusrSearchParams } from "./function/params";
import UixAllusrUsrlstMainpg from "./ui/lstusr/main";
import { MdlAllusrApplstParams, MdlAllusrSearchParams } from "./model/params";
import { ApiAllusrApplstDtbase } from "./api/applst";
import UixAllusrFormipMainpg from "./ui/create/form";

export default async function Page(props: {
    searchParams: Promise<MdlAllusrSearchParams>;
}) {
    const applst: MdlAllusrApplstParams[] = await ApiAllusrApplstDtbase();
    const qryprm = await props.searchParams;
    const prmAllusr = FncAllusrSearchParams(qryprm);
    return (
        <div className="afull flex justify-start items-start flex-wrap p-1.5 md:p-6">
            <div className="w-full md:w-60 min-w-1/2 h-250 md:h-250 max-h-fit p-3">
                <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200">
                    <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
                        Form create user
                        <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
                    </div>
                    <Suspense fallback={<UixGlobalLoadngAnmate />}>
                        <UixAllusrFormipMainpg applst={applst} prmAllusr={prmAllusr} />
                    </Suspense>
                </div>
            </div>
            <div className="w-full md:w-60 min-w-1/2 h-250 md:h-250 max-h-fit p-3">
                <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200 relative">
                    <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
                        List user
                        <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
                    </div>
                    <Suspense fallback={<UixGlobalLoadngAnmate />}>
                        <UixAllusrUsrlstMainpg prmAllusr={prmAllusr} />
                    </Suspense>
                </div>
            </div>
        </div>
    );
}