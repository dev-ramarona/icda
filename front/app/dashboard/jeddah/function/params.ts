import { MdlJeddahGlobalSrcprm } from "../model/params";

// Treatment function params
export function FncJeddahPnrsmrSrcprm(params: MdlJeddahGlobalSrcprm) {
  return {
    update_global: params.update_global || "",
    airlfl_jeddah: params.airlfl_jeddah || "",
    flnbfl_jeddah: params.flnbfl_jeddah || "",
    depart_jeddah: params.depart_jeddah || "",
    routfl_jeddah: params.routfl_jeddah || "",
    pnrcde_jeddah: params.pnrcde_jeddah || "",
    pagenw_jeddah: Number(params.pagenw_jeddah) || 1,
    limitp_jeddah: Number(params.limitp_jeddah) || 25,
  } as MdlJeddahGlobalSrcprm;
}
