import { mdlGlobalAllusrCookie } from "../../../global/model/params";
import UixPsglstPrcessManual from "./process";

export default function UixPsglstPrcessMainpg({ cookie, update }:
    { cookie: mdlGlobalAllusrCookie; update: string; }) {
    return (
        <>
            <UixPsglstPrcessManual cookie={cookie} update={update} />
        </>
    );
}