import { Suspense } from "react";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import UixPsglstErrlogMainpg from "./ui/errlog/main";
import {
  FncSlsflwErrlogSrcprm,
  FncSlsflwPsgdtlSrcprm,
  FncSlsflwPsgsmrSrcprm,
} from "./function/params";
import UixSlsflwDetailMainpg from "./ui/psgdtl/main";
import UixSlsflwActlogMainpg from "./ui/actlog/main";
import { MdlSlsflwActlogDtbase, MdlSlsflwGlobalSrcprm } from "./model/params";
import { ApiGlobalActlogDtbase } from "../global/api/dtbase";
import UixSlsflwPsgsmrMainpg from "./ui/psgsmr/main";
import UixSlsflwSmmry1Mainpg from "./ui/smmry1/main";
import { ApiAllusrCookieGetdta } from "../allusr/api/cookie";
import { ApiAllusrStatusPrcess } from "../allusr/api/status";

export default async function Page(props: { searchParams: Promise<MdlSlsflwGlobalSrcprm> }) {
  const cookie = await ApiAllusrCookieGetdta();
  const qryprm = await props.searchParams;
  const actobj = await ApiGlobalActlogDtbase("psglst");
  const actlog: MdlSlsflwActlogDtbase[] = actobj.actlog;
  const actdte: string[] = actobj.datefl;
  const status = await ApiAllusrStatusPrcess();
  const dfault = "SLSERR";
  const fmtdef = dfault == qryprm.format_psgdtl || !qryprm.format_psgdtl;
  const prmErrlog = FncSlsflwErrlogSrcprm(qryprm);
  const prmPsgdtl = FncSlsflwPsgdtlSrcprm(qryprm, actdte, dfault);
  const prmPsgsmr = FncSlsflwPsgsmrSrcprm(qryprm, actdte);
  return (
    <div className="afull flex flex-wrap items-start justify-start p-1.5 md:p-6">
      <div className="h-60 max-h-fit w-full min-w-1/5 p-3 md:h-80 md:w-40">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Log Action
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixSlsflwActlogMainpg actlog={actlog} />
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
            <UixPsglstErrlogMainpg prmErrlog={prmErrlog} status={status} cookie={cookie} />
          </Suspense>
        </div>
      </div>
      <div className="h-180 max-h-fit w-full min-w-full p-3 md:h-160 md:w-[20rem]">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Passangger Detail
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixSlsflwDetailMainpg
              prmPsgdtl={prmPsgdtl}
              datefl={actdte}
              cookie={cookie}
              fmtdef={fmtdef}
            />
          </Suspense>
        </div>
      </div>
      <div className="h-180 max-h-fit w-full min-w-full p-3 md:h-160 md:w-[20rem]">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Summary 30 Day
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            {/* <UixSlsflwDetailMainpg prmPsgdtl={prmPsgdtl} datefl={actdte} cookie={cookie} /> */}
            <UixSlsflwSmmry1Mainpg />
          </Suspense>
        </div>
      </div>
      <div className="h-180 max-h-fit w-full min-w-full p-3 md:h-160 md:w-[20rem]">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Passangger Summary
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixSlsflwPsgsmrMainpg prmPsgsmr={prmPsgsmr} datefl={actdte} cookie={cookie} />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
