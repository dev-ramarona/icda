import { ChangeEvent } from "react";
import UixGlobalInputxFormdt from "../../../global/ui/client/inputx";
import UixGlobalWaitngAction from "../action/waitng";

export default function UixGlobalWraperSearch({
  children,
  chnged,
  fmtdef,
  downld,
  upload,
  resetx,
  updtfl,
  namefl,
}: {
  children: React.ReactNode;
  chnged: boolean;
  fmtdef: boolean;
  downld: { lnk: string; prm: any } | null;
  upload: { lnk: string; prm: FileList } | null;
  resetx: () => void | null;
  updtfl: (e: ChangeEvent<HTMLInputElement, Element>) => void | null;
  namefl: string;
}) {
  return (
    <div className="flexctr relative z-30 h-24 min-h-fit w-full gap-3 py-3">
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
        className={`flexend w-1/3 flex-wrap gap-3 ${
          chnged ? "animate-pulse select-none" : ""
        } duration-300`}
      >
        {/* Download button */}
        {downld && (
          <form
            className="flexctr relative h-10 w-full md:w-28"
            method="POST"
            action="/dashboard/public/api/download"
          >
            <input
              type="hidden"
              name="link"
              value={`${process.env.NEXT_PUBLIC_URL_SERVER}${downld.lnk}`}
            />
            <input type="hidden" name="data" value={JSON.stringify(downld.prm)} />
            <button className="afull btnsbm flexctr" type="submit">
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
        {fmtdef && (
          <form
            className="flexctr relative h-20 w-full flex-col gap-1.5 rounded-md ring-2 ring-gray-200 md:h-10 md:w-59 md:flex-row"
            method="POST"
            encType="multipart/form-data"
            action="/dashboard/public/api/upload"
          >
            <input
              type="hidden"
              name="link"
              value={`${process.env.NEXT_PUBLIC_URL_SERVER}${upload.lnk}`}
            />
            <input type="hidden" name="data" value={JSON.stringify(downld.prm)} />
            <div className="flexctr h-1/2 w-full px-1.5 md:h-full md:w-2/3">
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
            <div className="h-1/2 w-full p-1 md:h-full md:w-1/3">
              <button className="btnsbm flexctr afull" type="submit">
                Update
              </button>
            </div>
          </form>
        )}
      </div>
    </div>
  );
}
