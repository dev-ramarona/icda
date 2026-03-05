import { ApiGlobalAcpedtDtbase } from "../../../global/api/dtbase";
import { MdlGlobalAcpedtDtbase, mdlGlobalAllusrCookie } from "../../../global/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/client/pagntn";
import { ApiPsglstPsgdtlGetall } from "../../api/psgdtl";
import {
  MdlPsglstPsgdtlFrntnd,
  MdlPsglstPsgdtlSrcprm,
} from "../../model/params";
import UixPsglstDetailSearch from "./search";
import UixPsglstDetailTablex from "./tablex";

export default async function UixPsglstDetailMainpg({
  prmPsgdtl,
  datefl,
  cookie,
}: {
  prmPsgdtl: MdlPsglstPsgdtlSrcprm;
  datefl: string[];
  cookie: mdlGlobalAllusrCookie;
}) {
  const psgdtl = await ApiPsglstPsgdtlGetall({
    ...prmPsgdtl, nclear_psgdtl:
      (prmPsgdtl.nclear_psgdtl == "") ? "MNFEST" : prmPsgdtl.nclear_psgdtl
  });
  const arrdta: MdlPsglstPsgdtlFrntnd[] = psgdtl.arrdta;
  const totdta: number = psgdtl.totdta;
  const acpedt: MdlGlobalAcpedtDtbase[] = await ApiGlobalAcpedtDtbase("mnfest");
  return (
    <>
      <UixPsglstDetailSearch prmPsgdtl={prmPsgdtl} datefl={datefl} />
      {arrdta.length > 0 ? (
        <UixPsglstDetailTablex detail={arrdta} acpedt={acpedt} cookie={cookie} />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgview={15}
        pgenbr={prmPsgdtl.pagenw_psgdtl}
        pgestr="pagenw_psgdtl"
        totdta={totdta}
      />
    </>
  );
}
