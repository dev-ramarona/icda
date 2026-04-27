import { UixGlobalIconvcRfresh } from "./ui/server/iconvc";

export default function Page() {
  return (
    <div className="flexctr fixed top-0 h-screen w-screen bg-linear-to-br from-zinc-100 via-white to-cyan-100">
      <div className="flexctr text-3xl font-semibold text-slate-600">
        <div>Loading</div>
        <div className="animate-spin">
          <UixGlobalIconvcRfresh bold={3} color="#45556c" size={2} />
        </div>
      </div>
    </div>
  );
}
