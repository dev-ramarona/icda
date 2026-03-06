import { MdlGlobalStatusPrcess } from "../../../global/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/client/pagntn";
import { ApiPsglstErrlogDtbase } from "../../../psglst/api/errlog";
import { MdlPsglstErrlogDtbase, MdlPsglstErrlogSrcprm } from "../../../psglst/model/params";
import UixPsglstErrlogTablex from "./tablex";

export default async function UixPsglstErrlogMainpg({ prmErrlog, status }:
  { prmErrlog: MdlPsglstErrlogSrcprm; status: MdlGlobalStatusPrcess }) {
  const rslobj = await ApiPsglstErrlogDtbase({
    ...prmErrlog, erdvsn_errlog:
      (prmErrlog.erdvsn_errlog == "") ? "SLSRPT" : prmErrlog.erdvsn_errlog
  });
  const errlog: MdlPsglstErrlogDtbase[] = rslobj.arrdta
  const totdta: number = rslobj.totdta
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
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Error
        </div>
      )}

    </>
  );
}
