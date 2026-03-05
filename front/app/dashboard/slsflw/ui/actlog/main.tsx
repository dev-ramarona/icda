import UixGlobalActlogTablex from "../../../global/ui/client/actlog";
import { MdlSlsflwActlogDtbase } from "../../model/params";

export default async function UixSlsflwActlogMainpg({ actlog }: { actlog: MdlSlsflwActlogDtbase[] }) {

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
