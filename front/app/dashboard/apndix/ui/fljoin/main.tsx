import { ApiApndixGetallDtbase } from "../../api/dtbase";
import {
  MdlApndixAcpedtDtbase,
  MdlApndixFljoinFrntnd,
  MdlApndixSearchQueryx,
} from "../../model/parmas";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/client/pagntn";
import UixPsglstFljoinTablex from "./table";
import UixApndixFljoinSearch from "./search";

export default async function UixApndixFljoinMainpg({
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
  const arrdta: MdlApndixFljoinFrntnd[] = rslobj.arrdta;
  const totdta: number = rslobj.totdta;
  return (
    <>
      <UixApndixFljoinSearch qryprm={qryprm} datefl={actdte} />
      <UixPsglstFljoinTablex
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
