import { JSX, Suspense } from "react";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import { ApiAllusrCookieGetdta } from "../allusr/api/cookie";
import { MdlApndixAcpedtDtbase, MdlApndixSearchQueryx } from "./model/parmas";
import { FncApndixSearchQueryx } from "./func/params";
import UixApndixApplstMainpg from "./ui/applst/main";
import UixApndixFlhourMainpg from "./ui/flhour/main";
import { ApiApndixAcpedtDtbase } from "./api/dtbase";
import UixApndixProvncMainpg from "./ui/provnc/main";


export default async function Page(props: {
    searchParams: Promise<MdlApndixSearchQueryx>;
}) {
    const cookie = await ApiAllusrCookieGetdta();
    const rawprm = await props.searchParams;
    const qryprm = FncApndixSearchQueryx(rawprm);
    const acpedt: MdlApndixAcpedtDtbase[] = await ApiApndixAcpedtDtbase(qryprm.pagedb);
    const lsmenu: Record<string, JSX.Element> = {
        flhour: <UixApndixFlhourMainpg cookie={cookie} qryprm={qryprm} acpedt={acpedt} />,
        provnc: <UixApndixProvncMainpg cookie={cookie} qryprm={qryprm} acpedt={acpedt} />,
        // frbase: <FrbaseTable />,
        // frtaxs: <FrtaxsTable />,
    }

    return (
        <div className="afull flex justify-start items-start flex-wrap p-1.5 md:p-6">
            <div className="w-full md:w-[20rem] min-w-full h-180 md:h-160 max-h-fit p-3">
                <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200 relative">
                    <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
                        Passangger detail
                        <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
                    </div>
                    <UixApndixApplstMainpg pagedb={qryprm.pagedb} />
                    <Suspense fallback={<UixGlobalLoadngAnmate />}>
                        {lsmenu[qryprm.pagedb ?? ""] ?? (
                            <div className="w-full h-16 flexctr text-base font-semibold text-sky-800">
                                Select menu
                            </div>
                        )}
                    </Suspense>
                </div>
            </div>
        </div>
    );
}