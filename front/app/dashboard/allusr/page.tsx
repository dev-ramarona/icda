import { Suspense } from "react";
import { UixGlobalIconvcSeting } from "../global/ui/server/iconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/loadng";
import { FncAllusrSearchParams } from "./function/params";
import UixAllusrUsrlstMainpg from "./ui/lstusr/main";
import { MdlAllusrApplstParams, MdlAllusrSearchParams } from "./model/params";
import { ApiAllusrApplstDtbase } from "./api/applst";
import UixAllusrFormipMainpg from "./ui/create/form";

export default async function Page(props: { searchParams: Promise<MdlAllusrSearchParams> }) {
  const applst: MdlAllusrApplstParams[] = await ApiAllusrApplstDtbase([]);
  const qryprm = await props.searchParams;
  const prmAllusr = FncAllusrSearchParams(qryprm);
  return (
    <div className="afull flex flex-wrap items-start justify-start p-1.5 md:p-6">
      <div className="h-250 max-h-fit w-full min-w-1/2 p-3 md:h-250 md:w-60">
        <div className="afull flexstr max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Form create user
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixAllusrFormipMainpg applst={applst} prmAllusr={prmAllusr} />
          </Suspense>
        </div>
      </div>
      <div className="h-250 max-h-fit w-full min-w-1/2 p-3 md:h-250 md:w-60">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
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
