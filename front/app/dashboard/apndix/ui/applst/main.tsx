import { ApiApndixApplstDtbase } from "../../api/dtbase";
import UixApndixApplstSelect from "./select";

export default async function UixApndixApplstMainpg({ pagedb }: { pagedb: string }) {
    const apndix = await ApiApndixApplstDtbase();
    return (
        <>
            <UixApndixApplstSelect apndix={apndix} pagedb={pagedb} />
        </>
    )
}