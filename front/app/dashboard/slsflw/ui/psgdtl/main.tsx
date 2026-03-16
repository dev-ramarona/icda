import UixSlsflwDetailSearch from "./search";
import UixSlsflwDetailTablex from "./tablex";
import { MdlSlsflwPsgdtlFrntnd, MdlSlsflwPsgdtlSrcprm } from "../../model/params";
import { ApiSlsflwPsgdtlGetall } from "../../api/psgdtl";
import UixGlobalPagntnMainpg from "../../../global/ui/client/pagntn";
import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import { MdlApndixAcpedtDtbase } from "../../../apndix/model/parmas";
import { ApiApndixAcpedtDtbase } from "../../../apndix/api/dtbase";

export default async function UixSlsflwDetailMainpg({
  prmPsgdtl,
  datefl,
  cookie,
}: {
  prmPsgdtl: MdlSlsflwPsgdtlSrcprm;
  datefl: string[];
  cookie: mdlAllusrCookieObjson;
}) {
  const psgdtl = await ApiSlsflwPsgdtlGetall({
    ...prmPsgdtl, nclear_psgdtl:
      (prmPsgdtl.nclear_psgdtl == "") ? "SLSRPT" : prmPsgdtl.nclear_psgdtl
  });
  const arrdta: MdlSlsflwPsgdtlFrntnd[] = psgdtl.arrdta;
  const totdta: number = psgdtl.totdta;
  const acpedt: MdlApndixAcpedtDtbase[] = await ApiApndixAcpedtDtbase("slsrpt");
  return (
    <>
      <UixSlsflwDetailSearch prmPsgdtl={prmPsgdtl} datefl={datefl} />
      {arrdta.length > 0 ? (
        <UixSlsflwDetailTablex detail={arrdta} acpedt={acpedt} cookie={cookie} />
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
