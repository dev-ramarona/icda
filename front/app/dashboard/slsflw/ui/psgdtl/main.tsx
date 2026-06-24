import UixSlsflwDetailSearch from "./search";
import UixSlsflwDetailTablex from "./tablex";
import { MdlSlsflwPsgdtlFrntnd, MdlSlsflwPsgdtlSrcprm } from "../../model/params";
import { ApiSlsflwPsgdtlGetall } from "../../api/psgdtl";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { MdlApndixAcpedtDtbase } from "../../../apndix/model/parmas";
import { ApiApndixAcpedtDtbase } from "../../../apndix/api/dtbase";

export default async function UixSlsflwDetailMainpg({
  prmPsgdtl,
  datefl,
  cookie,
  fmtdef,
  status,
  update,
}: {
  prmPsgdtl: MdlSlsflwPsgdtlSrcprm;
  datefl: string[];
  cookie: mdlAllusrCookieObjson;
  fmtdef: boolean;
  status: MdlAllusrStatusPrcess;
  update: string;
}) {
  prmPsgdtl.keywrd_psgdtl = JSON.stringify(cookie.keywrd);
  const psgdtl = await ApiSlsflwPsgdtlGetall(prmPsgdtl);
  const arrdta: MdlSlsflwPsgdtlFrntnd[] = psgdtl.arrdta;
  const totdta: number = psgdtl.totdta;
  const acpedt: MdlApndixAcpedtDtbase[] = await ApiApndixAcpedtDtbase("slsrpt");
  const ftrtru: string[] = ["regall", "regthl"];
  const ftricl = ftrtru.some((item) => cookie.keywrd.includes(item));
  const ftrfmt: string[] = ftricl ? ["SLSERR", "FMTHAI"] : ["SLSERR"];
  return (
    <>
      <UixSlsflwDetailSearch
        prmPsgdtl={prmPsgdtl}
        datefl={datefl}
        fmtdef={fmtdef}
        status={status}
        update={update}
        ftrfmt={ftrfmt}
      />
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
