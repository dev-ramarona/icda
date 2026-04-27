import UixGlobalPagntnMainpg from "../../../global/ui/action/pagntn";
import { ApiAllusrUsrlstGetall } from "../../api/user";
import { MdlAllusrFrntndParams, MdlAllusrSearchParams } from "../../model/params";
import UixCrtusrAllusrSearch from "./search";
import UixCrtusrAllusrTablex from "./table";

export default async function UixAllusrUsrlstMainpg({
  prmAllusr,
}: {
  prmAllusr: MdlAllusrSearchParams;
}) {
  const allusr = await ApiAllusrUsrlstGetall(prmAllusr);
  const arrdta: MdlAllusrFrntndParams[] = allusr.arrdta;
  const totdta: number = allusr.totdta;
  return (
    <>
      <UixCrtusrAllusrSearch prmAllusr={prmAllusr} />
      {arrdta.length > 0 ? (
        <UixCrtusrAllusrTablex allusr={arrdta} />
      ) : (
        <div className="flexctr h-fit w-full text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgview={prmAllusr.limitp}
        pgenbr={prmAllusr.pagenw}
        pgestr="pagenw"
        totdta={totdta}
      />
    </>
  );
}
