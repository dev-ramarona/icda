import { ApiAllusrApplstDtbase } from "./allusr/api/applst";
import { ApiAllusrCookieGetdta } from "./allusr/api/cookie";
import { MdlAllusrApplstParams } from "./allusr/model/params";
import UixGlobalAppbarClient from "./global/ui/action/appbar";
import UixGlobalHeaderClient from "./global/ui/action/header";

export default async function Layout({ children }: { children: React.ReactNode }) {
  const applst: MdlAllusrApplstParams[] = await ApiAllusrApplstDtbase([0]);
  const cookie = await ApiAllusrCookieGetdta();
  return (
    <div>
      <div className="flex h-20 w-full items-end justify-start px-3">
        <UixGlobalHeaderClient applst={applst} />
      </div>
      {children}
      <UixGlobalAppbarClient cookie={cookie} applst={applst} />
    </div>
  );
}
