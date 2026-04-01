import { Suspense } from "react";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import UixPsglstActlogMainpg from "./ui/actlog/main";
import UixPsglstErrlogMainpg from "./ui/errlog/main";
import UixPsglstDetailMainpg from "./ui/psgdtl/main";
import { MdlPsglstActlogDtbase, MdlPsglstGlobalSrcprm } from "./model/params";
import UixPsglstPrcessMainpg from "./ui/prcess/main";
import { FncPsglstErrlogSrcprm, FncPsglstPsgdtlSrcprm } from "./function/params";
import { ApiGlobalActlogDtbase } from "../global/api/dtbase";
import { ApiAllusrCookieGetdta } from "../allusr/api/cookie";
import { ApiAllusrStatusPrcess } from "../allusr/api/status";

export default async function Page(props: { searchParams: Promise<MdlPsglstGlobalSrcprm> }) {
  const cookie = await ApiAllusrCookieGetdta();
  const qryprm = await props.searchParams;
  const actobj = await ApiGlobalActlogDtbase("psglst");
  const actlog: MdlPsglstActlogDtbase[] = actobj.actlog;
  const actdte: string[] = actobj.datefl;
  const status = await ApiAllusrStatusPrcess();
  const dfault = "MNFERR";
  const fmtdef = dfault == qryprm.format_psgdtl || !qryprm.format_psgdtl;
  const prmErrlog = FncPsglstErrlogSrcprm(qryprm);
  const prmPsgdtl = FncPsglstPsgdtlSrcprm(qryprm, actdte, dfault);
  return (
    <div className="afull flex flex-wrap items-start justify-start p-1.5 md:p-6">
      <div className="h-60 max-h-fit w-full min-w-1/5 p-3 md:h-80 md:w-40">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Log Action
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstActlogMainpg actlog={actlog} />
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
            <UixPsglstErrlogMainpg prmErrlog={prmErrlog} status={status} />
          </Suspense>
        </div>
      </div>
      <div className="h-180 max-h-fit w-full min-w-full p-3 md:h-160 md:w-[20rem]">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Passangger detail
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstDetailMainpg
              prmPsgdtl={prmPsgdtl}
              datefl={actdte}
              cookie={cookie}
              fmtdef={fmtdef}
            />
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
            <UixPsglstPrcessMainpg
              cookie={cookie}
              update={prmPsgdtl.update_global}
              status={status}
            />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
