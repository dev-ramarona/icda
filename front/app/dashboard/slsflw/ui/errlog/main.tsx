import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import { ApiPsglstErrlogDtbase } from "../../../psglst/api/errlog";
import { MdlPsglstErrlogDtbase, MdlPsglstErrlogSrcprm } from "../../../psglst/model/params";
import UixPsglstErrlogTablex from "./tablex";

export default async function UixPsglstErrlogMainpg({
  prmErrlog,
  status,
  cookie,
  viewdt,
  showdt,
  pagest,
  pagenb,
}: {
  prmErrlog: MdlPsglstErrlogSrcprm;
  status: MdlAllusrStatusPrcess;
  cookie: mdlAllusrCookieObjson;
  viewdt: string;
  showdt: boolean;
  pagest: string;
  pagenb: number;
}) {
  const rslobj = await ApiPsglstErrlogDtbase({ ...prmErrlog, erdvsn_errlog: viewdt }, pagenb);
  const errlog: MdlPsglstErrlogDtbase[] = rslobj.arrdta;
  const totdta: number = rslobj.totdta;
  return (
    <>
      {errlog.length > 0 ? (
        <>
          <UixPsglstErrlogTablex
            errlog={errlog}
            update={prmErrlog.update_global}
            status={status}
            showdt={showdt}
          />
          <UixGlobalPagntnMainpg pgview={5} pgenbr={pagenb} pgestr={pagest} totdta={totdta} />
        </>
      ) : (
        <div className="flexctr h-fit w-full text-base font-semibold text-sky-800">
          No database Log Error
        </div>
      )}
    </>
  );
}
