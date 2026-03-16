import { ApiAllusrApplstDtbase } from "./allusr/api/applst";
import { ApiAllusrCookieGetdta } from "./allusr/api/cookie";
import { MdlAllusrApplstParams } from "./allusr/model/params";
import UixGlobalAppbarClient from "./global/ui/client/appbar";
import UixGlobalHeaderClient from "./global/ui/client/header";

export default async function Layout({ children }: { children: React.ReactNode }) {
    const applst: MdlAllusrApplstParams[] = await ApiAllusrApplstDtbase([0]);
    const cookie = await ApiAllusrCookieGetdta();
    return (
        <div>
            <div className="w-full h-20 flex justify-start items-end px-3">
                <UixGlobalHeaderClient applst={applst} />
            </div>
            {children}
            <UixGlobalAppbarClient cookie={cookie} applst={applst} />
        </div>
    );
}