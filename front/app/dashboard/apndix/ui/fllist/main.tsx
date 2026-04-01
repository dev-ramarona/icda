import { ApiApndixGetallDtbase } from "../../api/dtbase";
import {
  MdlApndixAcpedtDtbase,
  MdlApndixFllistFrntnd,
  MdlApndixSearchQueryx,
} from "../../model/parmas";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/client/pagntn";
import UixPsglstFllistTablex from "./table";
import UixApndixFllistSearch from "./search";
import { ApiGlobalActlogDtbase } from "../../../global/api/dtbase";

export default async function UixApndixFllistMainpg({
  qryprm,
  cookie,
  acpedt,
}: {
  qryprm: MdlApndixSearchQueryx;
  cookie: mdlAllusrCookieObjson;
  acpedt: MdlApndixAcpedtDtbase[];
}) {
  const rslobj = await ApiApndixGetallDtbase(qryprm);
  const arrdta: MdlApndixFllistFrntnd[] = rslobj.arrdta;
  const totdta: number = rslobj.totdta;
  const actobj = await ApiGlobalActlogDtbase("psglst");
  const actdte: string[] = actobj.datefl;
  return (
    <>
      <UixApndixFllistSearch qryprm={qryprm} datefl={actdte} />
      <UixPsglstFllistTablex
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
