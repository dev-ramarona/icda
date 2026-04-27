import { ApiApndixGetallDtbase } from "../../api/dtbase";
import {
  MdlApndixAcpedtDtbase,
  MdlApndixFrtaxsFrntnd,
  MdlApndixSearchQueryx,
} from "../../model/parmas";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import UixPsglstFrtaxsTablex from "./table";
import UixApndixFrtaxsSearch from "./search";

export default async function UixApndixFrtaxsMainpg({
  qryprm,
  cookie,
  acpedt,
}: {
  qryprm: MdlApndixSearchQueryx;
  cookie: mdlAllusrCookieObjson;
  acpedt: MdlApndixAcpedtDtbase[];
}) {
  const rslobj = await ApiApndixGetallDtbase(qryprm);
  const arrdta: MdlApndixFrtaxsFrntnd[] = rslobj.arrdta;
  const totdta: number = rslobj.totdta;
  return (
    <>
      <UixApndixFrtaxsSearch qryprm={qryprm} />
      <UixPsglstFrtaxsTablex
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
