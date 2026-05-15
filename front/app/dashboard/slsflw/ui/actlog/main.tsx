import { MdlAllusrStatusPrcess } from "../../../allusr/model/params";
import UixGlobalActlogTablex from "../../../global/ui/action/actlog";
import { MdlSlsflwActlogDtbase } from "../../model/params";

export default async function UixSlsflwActlogMainpg({
  actlog,
  update,
  status,
}: {
  actlog: MdlSlsflwActlogDtbase[];
  update: string;
  status: MdlAllusrStatusPrcess;
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
