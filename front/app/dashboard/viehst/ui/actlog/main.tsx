import { MdlGlobalActlogDtbase } from "../../../global/model/params";
import UixGlobalActlogTablex from "../../../global/ui/action/actlog";

export default async function UixViehstActlogMainpg({
  actlog,
}: {
  actlog: MdlGlobalActlogDtbase[];
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
