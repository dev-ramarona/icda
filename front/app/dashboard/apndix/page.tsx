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
import UixApndixFrbaseMainpg from "./ui/frbase/main";
import UixApndixFrtaxsMainpg from "./ui/frtaxs/main";
import UixApndixFllistMainpg from "./ui/fllist/main";
import { ApiGlobalActlogDtbase } from "../global/api/dtbase";
import UixApndixMilegeMainpg from "./ui/milege/main";

export default async function Page(props: { searchParams: Promise<MdlApndixSearchQueryx> }) {
  const cookie = await ApiAllusrCookieGetdta();
  const rawprm = await props.searchParams;
  const actobj = await ApiGlobalActlogDtbase("psglst");
  const actdte: string[] = actobj.datefl;
  const qryprm = FncApndixSearchQueryx(rawprm, actdte);
  const acpedt: MdlApndixAcpedtDtbase[] = await ApiApndixAcpedtDtbase(qryprm.pagedb_apndix);
  const lsmenu: Record<string, JSX.Element> = {
    flhour: <UixApndixFlhourMainpg cookie={cookie} qryprm={qryprm} acpedt={acpedt} />,
    provnc: <UixApndixProvncMainpg cookie={cookie} qryprm={qryprm} acpedt={acpedt} />,
    frbase: <UixApndixFrbaseMainpg cookie={cookie} qryprm={qryprm} acpedt={acpedt} />,
    frtaxs: <UixApndixFrtaxsMainpg cookie={cookie} qryprm={qryprm} acpedt={acpedt} />,
    milege: <UixApndixMilegeMainpg cookie={cookie} qryprm={qryprm} acpedt={acpedt} />,
    fllist: (
      <UixApndixFllistMainpg cookie={cookie} qryprm={qryprm} acpedt={acpedt} actdte={actdte} />
    ),
    // frbase: <FrbaseTable />,
    // frtaxs: <FrtaxsTable />,
  };

  return (
    <div className="afull flex flex-wrap items-start justify-start p-1.5 md:p-6">
      <div className="h-180 max-h-fit w-full min-w-full p-3 md:h-160 md:w-[20rem]">
        <div className="afull flexstr relative max-h-fit flex-col rounded-xl px-3 py-1.5 ring-2 ring-gray-200">
          <div className="flexstr w-full py-1.5 text-base font-semibold text-slate-800">
            Passangger detail
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <UixApndixApplstMainpg pagedb={qryprm.pagedb_apndix} cookie={cookie} />
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            {(cookie.keywrd &&
              (cookie.keywrd.includes("apndix") || cookie.keywrd.includes(qryprm.pagedb_apndix)) &&
              lsmenu[qryprm.pagedb_apndix]) ?? (
              <div className="flexctr h-16 w-full text-base font-semibold text-sky-800">
                Select menu
              </div>
            )}
          </Suspense>
        </div>
      </div>
    </div>
  );
}
