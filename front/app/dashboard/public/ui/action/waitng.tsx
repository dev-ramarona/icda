import { UixGlobalIconvcRfresh } from "../../../global/ui/server/iconvc";

export default function UixGlobalWaitngAction({ chnged }: { chnged: boolean }) {
    return (
        <div className={`flexctr absolute pointer-events-none z-20 duration-300 
                ${chnged ? "w-full h-full translate-y-0" : "opacity-0 -translate-y-20"}`}>
            <div className="w-20 h-10 bg-white ring-2 ring-cyan-600 px-5 py-2 rounded-xl flexctr drop-shadow-sm drop-shadow-gray-600">
                <div>Wait</div>
                <div className="animate-spin">
                    <UixGlobalIconvcRfresh bold={3} color="#0092b8" size={1.2} />
                </div>
            </div>
        </div>
    );
}