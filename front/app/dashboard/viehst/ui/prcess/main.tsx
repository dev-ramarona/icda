import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { MdlViehstGlobalSrcprm } from "../../model/params";
import UixViehstPrcessManual from "./process";

export default function UixViehstPrcessMainpg({
  cookie,
  update,
  status,
  prmGlobal,
}: {
  cookie: mdlAllusrCookieObjson;
  update: string;
  status: MdlAllusrStatusPrcess;
  prmGlobal: MdlViehstGlobalSrcprm;
}) {
  return (
    <>
      {cookie.keywrd && cookie.keywrd.includes("viehst") ? (
        <UixViehstPrcessManual cookie={cookie} update={update} status={status} queryx={prmGlobal} />
      ) : (
        <div className="flexctr h-10 w-full text-gray-600">Only accepted users</div>
      )}
    </>
  );
}
