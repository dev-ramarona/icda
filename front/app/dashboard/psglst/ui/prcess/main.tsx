import { mdlGlobalAllusrCookie, MdlGlobalStatusPrcess } from "../../../global/model/params";
import UixPsglstPrcessManual from "./process";

export default function UixPsglstPrcessMainpg({ cookie, update, status }:
    { cookie: mdlGlobalAllusrCookie; update: string; status: MdlGlobalStatusPrcess }) {
    return (
        <>
            <UixPsglstPrcessManual cookie={cookie} update={update} status={status} />
        </>
    );
}