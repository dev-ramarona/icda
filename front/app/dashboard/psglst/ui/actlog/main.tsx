import { MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import { MdlGlobalActlogDtbase } from "../../../global/model/params";
import UixGlobalActlogTablex from "../../../global/ui/action/actlog";

export default async function UixPsglstActlogMainpg({
  actlog,
  status,
  update,
}: {
  actlog: MdlGlobalActlogDtbase[];
  status: MdlAllusrStatusPrcess;
  update: string;
}) {
  return (
    <>
      {actlog.length > 0 ? (
        <UixGlobalActlogTablex actlog={actlog} status={status} update={update} />
      ) : (
        <div className="flexctr h-fit w-full text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
