import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import { ApiJeddahPnrsmrGetall } from "../../api/prcess";
import { MdlJeddahGlobalSrcprm, MdlJeddahPrcessDtbase } from "../../model/params";
import UixJeddahPnrsmrSearch from "./search";
import UixJeddahPnrsmrTablex from "./tablex";

export default async function UixJeddahPnrsmrMainpg({
  prmPnrsmr,
  cookie,
  status,
  update,
}: {
  prmPnrsmr: MdlJeddahGlobalSrcprm;
  cookie: mdlAllusrCookieObjson;
  status: MdlAllusrStatusPrcess;
  update: string;
}) {
  const pnrsmr = await ApiJeddahPnrsmrGetall(prmPnrsmr);
  const arrdta: MdlJeddahPrcessDtbase[] = pnrsmr.arrdta;
  const totdta: number = pnrsmr.totdta;
  return (
    <>
      <UixJeddahPnrsmrSearch prmPnrsmr={prmPnrsmr} status={status} update={update} />
      {arrdta.length > 0 ? (
        <UixJeddahPnrsmrTablex arrdta={arrdta} />
      ) : (
        <div className="flexctr h-fit w-full text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgview={15}
        pgenbr={prmPnrsmr.pagenw_jeddah}
        pgestr="pagenw_jeddah"
        totdta={totdta}
      />
    </>
  );
}
