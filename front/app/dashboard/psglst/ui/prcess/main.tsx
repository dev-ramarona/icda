import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import UixPsglstPrcessManual from "./process";

export default function UixPsglstPrcessMainpg({ cookie, update, status }:
    { cookie: mdlAllusrCookieObjson; update: string; status: MdlAllusrStatusPrcess }) {
    return (
        <>
            <UixPsglstPrcessManual cookie={cookie} update={update} status={status} />
        </>
    );
}