import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";

export default function UixGlobalWaitngAction({ chnged }: { chnged: boolean }) {
  return (
    <div
      className={`flexctr pointer-events-none absolute z-20 duration-300 ${chnged ? "h-full w-full translate-y-0" : "-translate-y-20 opacity-0"}`}
    >
      <div className="flexctr h-10 w-20 rounded-xl bg-white px-5 py-2 ring-2 ring-cyan-600 drop-shadow-sm drop-shadow-gray-600">
        <div>Wait</div>
        <div className="animate-spin">
          <UixGlobalIconvcRfresh bold={3} color="#0092b8" size={1.2} />
        </div>
      </div>
    </div>
  );
}
