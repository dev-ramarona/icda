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
  fmtdef,
}: {
  prmPsgdtl: MdlSlsflwPsgdtlSrcprm;
  datefl: string[];
  cookie: mdlAllusrCookieObjson;
  fmtdef: boolean;
}) {
  const psgdtl = await ApiSlsflwPsgdtlGetall({
    ...prmPsgdtl,
    keywrd_psgdtl: JSON.stringify(cookie.keywrd.filter((item) => item.includes("REG "))),
  });
  const arrdta: MdlSlsflwPsgdtlFrntnd[] = psgdtl.arrdta;
  const totdta: number = psgdtl.totdta;
  const acpedt: MdlApndixAcpedtDtbase[] = await ApiApndixAcpedtDtbase("slsrpt");
  return (
    <>
      <UixSlsflwDetailSearch prmPsgdtl={prmPsgdtl} datefl={datefl} fmtdef={fmtdef} />
      {arrdta.length > 0 ? (
        <UixSlsflwDetailTablex arrdta={arrdta} acpedt={acpedt} cookie={cookie} fmtdef={fmtdef} />
      ) : (
        <div className="flexctr h-fit w-full text-base font-semibold text-sky-800">
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
