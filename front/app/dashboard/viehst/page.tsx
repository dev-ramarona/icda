import { Suspense } from "react";
import { ApiAllusrCookieGetdta } from "../allusr/api/cookie";
import { ApiAllusrStatusPrcess } from "../allusr/api/status";
import { ApiGlobalActlogDtbase } from "../global/api/dtbase";
import { MdlGlobalActlogDtbase } from "../global/model/params";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import { MdlViehstGlobalSrcprm } from "./model/params";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import UixViehstActlogMainpg from "./ui/actlog/main";
import UixViehstPrcessMainpg from "./ui/prcess/main";
import { FncViehstPsgdtlSrcprm } from "./function/params";

export default async function Page(props: { searchParams: Promise<MdlViehstGlobalSrcprm> }) {
  const cookie = await ApiAllusrCookieGetdta();
  const qryprm = await props.searchParams;
  // const actobj = await ApiGlobalActlogDtbase("viehst");
  // const actlog: MdlGlobalActlogDtbase[] = actobj.actlog;
  // const actdte: string[] = actobj.datefl;
  const status = await ApiAllusrStatusPrcess();
  const dfault = "MNFERR";
  const prmGlobal = FncViehstPsgdtlSrcprm(qryprm);
  return (
    <div className="afull flex flex-wrap items-start justify-start p-1.5 md:p-6">
      <div className="h-60 max-h-fit w-full min-w-1/5 p-3 md:h-80 md:w-40">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Log Action
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            {/* <UixViehstActlogMainpg actlog={actlog} /> */}
          </Suspense>
        </div>
      </div>
      <div className="h-120 max-h-fit w-full min-w-4/5 p-3 md:h-80 md:w-60">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Log error
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            {/* <UixViehstErrlogMainpg prmErrlog={prmErrlog} status={status} /> */}
          </Suspense>
        </div>
      </div>
      <div className="h-100 max-h-fit w-full min-w-full p-3 md:h-100 md:w-[20rem]">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Process manual
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixViehstPrcessMainpg
              cookie={cookie}
              update={prmGlobal.update_global}
              status={status}
              prmGlobal={prmGlobal}
            />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
