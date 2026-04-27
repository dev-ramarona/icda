import { SetStateAction } from "react";
import { UixGlobalIconvcCancel, UixGlobalIconvcCeklis } from "../server/iconvc";
import { MdlGlobalConfrmAction } from "../../model/params";

export default function UixGlobalConfrmAction({
  confrm,
  confdt,
  action,
  goupdt,
  confrmSet,
}: {
  confrm: boolean;
  confdt: MdlGlobalConfrmAction[];
  action: string;
  goupdt: () => Promise<void>;
  confrmSet: (value: SetStateAction<boolean>) => void;
}) {
  return (
    <div
      className={`flexctr absolute z-30 h-full w-full ${
        confrm ? "backdrop-blur-[2px]" : "pointer-events-none"
      }`}
    >
      <div
        className={`flexctr flex-col gap-y-6 rounded-xl bg-white text-gray-700 ring-2 ring-gray-200 ${
          confrm ? "h-60 w-72 translate-y-0 p-3" : "h-0 w-0 -translate-y-10 opacity-0"
        } overflow-hidden duration-300`}
      >
        <div className="text-center text-lg font-semibold text-nowrap text-gray-600">
          <span>Confirm</span>
          <span className="font-bold text-red-600"> {action} </span>
          <span>this data?</span>
        </div>
        <div className="w-full overflow-y-auto">
          {confdt.map((val, idx) => (
            <div className="flexctr gap-1.5" key={idx}>
              <div className="w-20 font-bold whitespace-nowrap text-gray-600">{val.paramx}</div>
              <div className="flexctr w-3">:</div>
              <div className="w-full overflow-x-auto whitespace-nowrap">{val.valuex}</div>
            </div>
          ))}
        </div>
        <div className="flexctr gap-3">
          <div className="btnsbm flexctr h-8 w-10" onClick={() => goupdt()}>
            <UixGlobalIconvcCeklis bold={4} color="#53eafd" size={1.4} />
          </div>
          <div className="btnsbm flexctr h-8 w-10" onClick={() => confrmSet(false)}>
            <UixGlobalIconvcCancel bold={4} color="#fb2c36" size={1.4} />
          </div>
        </div>
      </div>
    </div>
  );
}
