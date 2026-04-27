import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import { ApiApndixAcpedtDtbase } from "../../../apndix/api/dtbase";
import { MdlApndixAcpedtDtbase } from "../../../apndix/model/parmas";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import { ApiPsglstPsgdtlGetall } from "../../api/psgdtl";
import { MdlPsglstPsgdtlFrntnd, MdlPsglstPsgdtlSrcprm } from "../../model/params";
import UixPsglstDetailSearch from "./search";
import UixPsglstDetailTablex from "./tablex";

export default async function UixPsglstDetailMainpg({
  prmPsgdtl,
  datefl,
  cookie,
  fmtdef,
}: {
  prmPsgdtl: MdlPsglstPsgdtlSrcprm;
  datefl: string[];
  cookie: mdlAllusrCookieObjson;
  fmtdef: boolean;
}) {
  const psgdtl = await ApiPsglstPsgdtlGetall({
    ...prmPsgdtl,
    nclear_psgdtl: prmPsgdtl.nclear_psgdtl == "" ? "MNFEST" : prmPsgdtl.nclear_psgdtl,
  });
  const arrdta: MdlPsglstPsgdtlFrntnd[] = psgdtl.arrdta;
  const totdta: number = psgdtl.totdta;
  const acpedt: MdlApndixAcpedtDtbase[] = await ApiApndixAcpedtDtbase("mnfest");
  return (
    <>
      <UixPsglstDetailSearch prmPsgdtl={prmPsgdtl} datefl={datefl} fmtdef={fmtdef} />
      <UixPsglstDetailTablex arrdta={arrdta} acpedt={acpedt} cookie={cookie} fmtdef={fmtdef} />
      <UixGlobalPagntnMainpg
        pgview={15}
        pgenbr={prmPsgdtl.pagenw_psgdtl}
        pgestr="pagenw_psgdtl"
        totdta={totdta}
      />
    </>
  );
}
