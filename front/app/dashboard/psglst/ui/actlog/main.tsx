import UixGlobalActlogTablex from "../../../global/ui/action/actlog";
import { MdlPsglstActlogDtbase } from "../../model/params";

export default async function UixPsglstActlogMainpg({
  actlog,
}: {
  actlog: MdlPsglstActlogDtbase[];
}) {
  return (
    <>
      {actlog.length > 0 ? (
        <UixGlobalActlogTablex actlog={actlog} />
      ) : (
        <div className="flexctr h-fit w-full text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
