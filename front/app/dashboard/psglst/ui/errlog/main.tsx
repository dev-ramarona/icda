import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import { ApiPsglstErrlogDtbase } from "../../api/errlog";
import { MdlPsglstErrlogDtbase, MdlPsglstErrlogSrcprm } from "../../model/params";
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
      <UixPsglstErrlogTablex
        errlog={errlog}
        update={prmErrlog.update_global}
        status={status}
        showdt={showdt}
      />
      <UixGlobalPagntnMainpg pgview={5} pgenbr={pagenb} pgestr={pagest} totdta={totdta} />
    </>
  );
}
