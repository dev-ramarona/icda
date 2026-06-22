import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import UixPsglstPrcessManual from "./process";

export default function UixPsglstPrcessMainpg({
  cookie,
  update,
  status,
}: {
  cookie: mdlAllusrCookieObjson;
  update: string;
  status: MdlAllusrStatusPrcess;
}) {
  return (
    <>
      {cookie.keywrd && cookie.keywrd.includes("psglst") ? (
        <UixPsglstPrcessManual cookie={cookie} update={update} status={status} />
      ) : (
        <div className="flexctr h-10 w-full text-gray-600">Only accepted users</div>
      )}
    </>
  );
}
