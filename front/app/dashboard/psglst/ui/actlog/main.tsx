import UixGlobalActlogTablex from "../../../global/ui/client/actlog";
import { MdlPsglstActlogDtbase } from "../../model/params";

export default async function UixPsglstActlogMainpg({ actlog }: { actlog: MdlPsglstActlogDtbase[] }) {

  return (
    <>
      {actlog.length > 0 ? (
        <UixGlobalActlogTablex actlog={actlog} />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
