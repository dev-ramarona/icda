import { mdlAllusrCookieObjson, MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import { ApiSlsflwPsgsmrGetall } from "../../api/psgsmr";
import { MdlSlsflwPsgsmrFrntnd, MdlSlsflwPsgsmrSrcprm } from "../../model/params";
import UixSlsflwPsgsmrSearch from "./search";
import UixSlsflwPsgsmrTablex from "./tablex";

export default async function UixSlsflwPsgsmrMainpg({
  prmPsgsmr,
  datefl,
  cookie,
  status,
  update,
}: {
  prmPsgsmr: MdlSlsflwPsgsmrSrcprm;
  datefl: string[];
  cookie: mdlAllusrCookieObjson;
  status: MdlAllusrStatusPrcess;
  update: string;
}) {
  prmPsgsmr.keywrd_psgsmr = JSON.stringify(cookie.keywrd);
  const psgsmr = await ApiSlsflwPsgsmrGetall(prmPsgsmr);
  const arrdta: MdlSlsflwPsgsmrFrntnd[] = psgsmr.arrdta;
  const totdta: number = psgsmr.totdta;
  const joinrd: boolean = psgsmr.joinrd;
  const joinar: string[] = joinrd ? ["Combined", "Separated"] : ["Separated"];
  return (
    <>
      <UixSlsflwPsgsmrSearch
        prmPsgsmr={prmPsgsmr}
        datefl={datefl}
        joinar={joinar}
        status={status}
        update={update}
      />
      {arrdta.length > 0 ? (
        <UixSlsflwPsgsmrTablex arrdta={arrdta} />
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
