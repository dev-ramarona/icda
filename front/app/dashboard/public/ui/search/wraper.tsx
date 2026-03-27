import { ChangeEvent } from "react";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import UixGlobalWaitngAction from "../action/waitng";

export default function UixGlobalWraperSearch({
  children,
  chnged,
  downld,
  resetx,
  updtfl,
  namefl,
}: {
  children: React.ReactNode;
  chnged: boolean;
  downld: { lnk: string; prm: any } | null;
  resetx: () => void | null;
  updtfl: (e: ChangeEvent<HTMLInputElement, Element>) => void | null;
  namefl: string;
}) {
  return (
    <div className="flexctr relative z-30 h-24 min-h-fit w-full py-3">
      <UixGlobalWaitngAction chnged={chnged} />
      <div
        className={`afull flexstr flex-wrap gap-y-3 ${
          chnged ? "animate-pulse select-none" : ""
        } duration-300`}
      >
        {/* Main data search */}
        {children}
      </div>
      <div
        className={`flexend w-1/3 flex-wrap gap-3 px-3 ${
          chnged ? "animate-pulse select-none" : ""
        } duration-300`}
      >
        {/* Download button */}
        {downld && (
          <form className="flexctr relative h-10 w-full md:w-28" method="POST" action={downld.lnk}>
            <input type="hidden" name="data" value={JSON.stringify(downld.prm)} />
            <button type="submit" className="afull btnsbm flexctr">
              Download
            </button>
          </form>
        )}

        {/* Reset button */}
        {resetx && (
          <div className="flexctr relative h-10 w-full md:w-28">
            <div className="afull btnwrn flexctr" onClick={() => resetx()}>
              Reset
            </div>
          </div>
        )}

        {/* Reset button */}
        {resetx && (
          <div className="flexctr relative h-20 w-full flex-col gap-1.5 rounded-md p-1.5 ring-2 ring-gray-200 md:w-28">
            <div className="btnstb flexctr h-1/2 w-full">
              <UixGlobalInputxFormdt
                labelx="hidden"
                length={13}
                params={namefl}
                plchdr="Select file"
                queryx="file"
                repprm={updtfl}
                typipt="file"
              />
            </div>
            <div className="btnsbm flexctr h-1/2 w-full">Update</div>
          </div>
        )}
      </div>
    </div>
  );
}
