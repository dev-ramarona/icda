import { Suspense } from "react";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import UixPsglstErrlogMainpg from "./ui/errlog/main";
import { FncSlsflwErrlogSrcprm, FncSlsflwPsgdtlSrcprm, FncSlsflwPsgsmrSrcprm } from "./function/params";
import UixSlsflwDetailMainpg from "./ui/psgdtl/main";
import UixSlsflwActlogMainpg from "./ui/actlog/main";
import { MdlSlsflwActlogDtbase, MdlSlsflwGlobalSrcprm } from "./model/params";
import { ApiGlobalActlogDtbase } from "../global/api/dtbase";
import UixSlsflwPsgsmrMainpg from "./ui/psgsmr/main";
import UixSlsflwSmmry1Mainpg from "./ui/smmry1/main";
import { ApiAllusrCookieGetdta } from "../allusr/api/cookie";
import { ApiAllusrStatusPrcess } from "../allusr/api/status";


export default async function Page(props: {
  searchParams: Promise<MdlSlsflwGlobalSrcprm>;
}) {
  const cookie = await ApiAllusrCookieGetdta();
  const qryprm = await props.searchParams;
  const actobj = await ApiGlobalActlogDtbase("psglst");
  const actlog: MdlSlsflwActlogDtbase[] = actobj.actlog;
  const actdte: string[] = actobj.datefl;
  const status = await ApiAllusrStatusPrcess();
  const prmErrlog = FncSlsflwErrlogSrcprm(qryprm);
  const prmPsgdtl = FncSlsflwPsgdtlSrcprm(qryprm, actdte);
  const prmPsgsmr = FncSlsflwPsgsmrSrcprm(qryprm, actdte);
  return (
    <div className="afull flex justify-start items-start flex-wrap p-1.5 md:p-6">
      <div className="w-full md:w-40 min-w-1/5 h-60 md:h-80 max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Log Action
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixSlsflwActlogMainpg actlog={actlog} />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-60 min-w-4/5 h-120 md:h-80 max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Log error
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstErrlogMainpg prmErrlog={prmErrlog} status={status} />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[20rem] min-w-full h-180 md:h-160 max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Passangger Detail
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixSlsflwDetailMainpg prmPsgdtl={prmPsgdtl} datefl={actdte} cookie={cookie} />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[20rem] min-w-full h-180 md:h-160 max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Summary 30 Day
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            {/* <UixSlsflwDetailMainpg prmPsgdtl={prmPsgdtl} datefl={actdte} cookie={cookie} /> */}
            <UixSlsflwSmmry1Mainpg />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[20rem] min-w-full h-180 md:h-160 max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Passangger Summary
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixSlsflwPsgsmrMainpg prmPsgsmr={prmPsgsmr} datefl={actdte} cookie={cookie} />
          </Suspense>
        </div>
      </div>
    </div >
  );
}
