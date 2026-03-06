import { Suspense } from "react";
import { ApiGlobalCookieGetdta } from "../global/api/cookie";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import UixPsglstActlogMainpg from "./ui/actlog/main";
import UixPsglstErrlogMainpg from "./ui/errlog/main";
import UixPsglstDetailMainpg from "./ui/psgdtl/main";
import { MdlPsglstActlogDtbase, MdlPsglstGlobalSrcprm } from "./model/params";
import UixPsglstPrcessMainpg from "./ui/prcess/main";
import { FncPsglstErrlogSrcprm, FncPsglstPsgdtlSrcprm } from "./function/params";
import { ApiGlobalActlogDtbase } from "../global/api/dtbase";
import { ApiGlobalStatusPrcess } from "../global/api/status";

export default async function Page(props: {
  searchParams: Promise<MdlPsglstGlobalSrcprm>;
}) {
  const cookie = await ApiGlobalCookieGetdta();
  const qryprm = await props.searchParams;
  const actobj = await ApiGlobalActlogDtbase("psglst");
  const actlog: MdlPsglstActlogDtbase[] = actobj.actlog;
  const actdte: string[] = actobj.datefl;
  const status = await ApiGlobalStatusPrcess();
  const prmErrlog = FncPsglstErrlogSrcprm(qryprm);
  const prmPsgdtl = FncPsglstPsgdtlSrcprm(qryprm, actdte);
  return (
    <div className="afull flex justify-start items-start flex-wrap p-1.5 md:p-6">
      <div className="w-full md:w-40 min-w-1/5 h-60 md:h-80 max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Log Action
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstActlogMainpg actlog={actlog} />
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
            Passangger detail
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstDetailMainpg prmPsgdtl={prmPsgdtl} datefl={actdte} cookie={cookie} />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[20rem] min-w-full h-100 md:h-100 max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col ring-2 ring-gray-200">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Process manual
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstPrcessMainpg cookie={cookie} update={prmPsgdtl.update_global} status={status} />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
