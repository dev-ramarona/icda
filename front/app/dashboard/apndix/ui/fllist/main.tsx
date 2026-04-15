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
  actdte,
}: {
  qryprm: MdlApndixSearchQueryx;
  cookie: mdlAllusrCookieObjson;
  acpedt: MdlApndixAcpedtDtbase[];
  actdte: string[];
}) {
  const rslobj = await ApiApndixGetallDtbase(qryprm);
  const arrdta: MdlApndixFllistFrntnd[] = rslobj.arrdta;
  const totdta: number = rslobj.totdta;

  // Get h+1 day
  const nowdte = new Date();
  nowdte.setDate(nowdte.getDate() + 1);
  const yearxx = nowdte.getFullYear();
  const monthx = String(nowdte.getMonth() + 1).padStart(2, "0");
  const dayxxx = String(nowdte.getDate()).padStart(2, "0");
  const formtd = `${yearxx}-${monthx}-${dayxxx}`;
  actdte.push(formtd);
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
