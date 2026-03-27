import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import { ApiApndixAcpedtDtbase } from "../../../apndix/api/dtbase";
import { MdlApndixAcpedtDtbase } from "../../../apndix/model/parmas";
import UixGlobalPagntnMainpg from "../../../global/ui/client/pagntn";
import { ApiSlsflwPsgsmrGetall } from "../../api/psgsmr";
import { MdlSlsflwPsgsmrFrntnd, MdlSlsflwPsgsmrSrcprm } from "../../model/params";
import UixSlsflwPsgsmrSearch from "./search";
import UixSlsflwPsgsmrTablex from "./tablex";

export default async function UixSlsflwPsgsmrMainpg({
  prmPsgsmr,
  datefl,
  cookie,
}: {
  prmPsgsmr: MdlSlsflwPsgsmrSrcprm;
  datefl: string[];
  cookie: mdlAllusrCookieObjson;
}) {
  const psgsmr = await ApiSlsflwPsgsmrGetall(prmPsgsmr);
  const arrdta: MdlSlsflwPsgsmrFrntnd[] = psgsmr.arrdta;
  const totdta: number = psgsmr.totdta;
  return (
    <>
      <UixSlsflwPsgsmrSearch prmPsgsmr={prmPsgsmr} datefl={datefl} />
      {arrdta.length > 0 ? (
        <UixSlsflwPsgsmrTablex psgsmr={arrdta} />
      ) : (
        <div className="flexctr h-fit w-full text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgview={15}
        pgenbr={prmPsgsmr.pagenw_psgsmr}
        pgestr="pagenw_psgsmr"
        totdta={totdta}
      />
    </>
  );
}
