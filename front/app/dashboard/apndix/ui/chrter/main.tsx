import { ApiApndixGetallDtbase } from "../../api/dtbase";
import {
  MdlApndixAcpedtDtbase,
  MdlApndixChrterFrntnd,
  MdlApndixSearchQueryx,
} from "../../model/parmas";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import UixPsglstChrtercTablex from "./table";
import UixApndixChrtercSearch from "./search";

export default async function UixApndixChrtercMainpg({
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
  const arrdta: MdlApndixChrterFrntnd[] = rslobj.arrdta;
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
      <UixApndixChrtercSearch qryprm={qryprm} datefl={actdte} />
      <UixPsglstChrtercTablex
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
