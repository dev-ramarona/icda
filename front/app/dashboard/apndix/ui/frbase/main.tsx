import { ApiApndixGetallDtbase } from "../../api/dtbase";
import {
  MdlApndixAcpedtDtbase,
  MdlApndixFrbaseFrntnd,
  MdlApndixSearchQueryx,
} from "../../model/parmas";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/client/pagntn";
import UixPsglstFrbaseTablex from "./table";
import UixApndixFrbaseSearch from "./search";

export default async function UixApndixFrbaseMainpg({
  qryprm,
  cookie,
  acpedt,
}: {
  qryprm: MdlApndixSearchQueryx;
  cookie: mdlAllusrCookieObjson;
  acpedt: MdlApndixAcpedtDtbase[];
}) {
  const rslobj = await ApiApndixGetallDtbase(qryprm);
  const arrdta: MdlApndixFrbaseFrntnd[] = rslobj.arrdta;
  const totdta: number = rslobj.totdta;
  return (
    <>
      <UixApndixFrbaseSearch qryprm={qryprm} />
      <UixPsglstFrbaseTablex
        acpedt={acpedt}
        arrdta={arrdta}
        pagedb={qryprm.pagedb_apndix}
        cookie={cookie}
      />
      <UixGlobalPagntnMainpg
        pgview={5}
        pgenbr={qryprm.pagenw_apndix}
        pgestr="pagenw_apndix"
        totdta={totdta}
      />
    </>
  );
}
