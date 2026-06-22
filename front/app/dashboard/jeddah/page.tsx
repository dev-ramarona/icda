import { Suspense } from "react";
import { ApiAllusrCookieGetdta } from "../allusr/api/cookie";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import { MdlJeddahGlobalSrcprm } from "./model/params";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import UixJeddahPrcessManual from "./ui/prcess/process";
import { MdlAllusrStatusPrcess } from "../allusr/model/params";
import UixJeddahPnrsmrMainpg from "./ui/pnrsmr/main";
import { FncJeddahPnrsmrSrcprm } from "./function/params";

export default async function Page(props: { searchParams: Promise<MdlJeddahGlobalSrcprm> }) {
  const cookie = await ApiAllusrCookieGetdta();
  const qryprm = await props.searchParams;
  const prmPnrsmr = FncJeddahPnrsmrSrcprm(qryprm);
  const status: MdlAllusrStatusPrcess = { action: "", sbrapi: 0 };

  return (
    <div className="afull flex flex-wrap items-start justify-start p-1.5 md:p-6">
      <div className="h-100 max-h-fit w-full min-w-full p-3 md:h-100 md:w-[20rem]">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Process manual
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixJeddahPrcessManual
              cookie={cookie}
              update={prmPnrsmr.update_global}
              status={status}
            />
          </Suspense>
        </div>
      </div>
      <div className="h-180 max-h-fit w-full min-w-full p-3 md:h-160 md:w-[20rem]">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Table
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixJeddahPnrsmrMainpg
              cookie={cookie}
              status={status}
              update={prmPnrsmr.update_global}
              prmPnrsmr={prmPnrsmr}
            />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
