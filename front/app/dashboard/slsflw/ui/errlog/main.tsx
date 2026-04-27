import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import { ApiPsglstErrlogDtbase } from "../../../psglst/api/errlog";
import { MdlPsglstErrlogDtbase, MdlPsglstErrlogSrcprm } from "../../../psglst/model/params";
import UixPsglstErrlogTablex from "./tablex";

export default async function UixPsglstErrlogMainpg({
  prmErrlog,
  status,
  cookie,
}: {
  prmErrlog: MdlPsglstErrlogSrcprm;
  status: MdlAllusrStatusPrcess;
  cookie: mdlAllusrCookieObjson;
}) {
  const rslobj = await ApiPsglstErrlogDtbase({
    ...prmErrlog,
    erdvsn_errlog: prmErrlog.erdvsn_errlog == "" ? "SLSRPT" : prmErrlog.erdvsn_errlog,
  });
  const errlog: MdlPsglstErrlogDtbase[] = rslobj.arrdta;
  const totdta: number = rslobj.totdta;
  return (
    <>
      {errlog.length > 0 ? (
        <>
          <UixPsglstErrlogTablex errlog={errlog} update={prmErrlog.update_global} status={status} />
          <UixGlobalPagntnMainpg
            pgview={5}
            pgenbr={prmErrlog.pagenw_errlog}
            pgestr="pagenw_errlog"
            totdta={totdta}
          />
        </>
      ) : (
        <div className="flexctr h-fit w-full text-base font-semibold text-sky-800">
          No database Log Error
        </div>
      )}
    </>
  );
}
