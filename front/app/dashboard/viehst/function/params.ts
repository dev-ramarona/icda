import { MdlViehstGlobalSrcprm } from "../model/params";

// Treatment function params
export function FncViehstPsgdtlSrcprm(params: MdlViehstGlobalSrcprm) {
  return {
    update_global: params.update_global,
    datefl_prcess: params.datefl_prcess,
    airlfl_prcess: params.airlfl_prcess,
    flnbfl_prcess: params.flnbfl_prcess,
    depart_prcess: params.depart_prcess,
    worker_prcess: params.worker_prcess,
    pagenw_viehst: Number(params.pagenw_viehst) || 1,
    limitp_viehst: Number(params.limitp_viehst) || 15,
  } as MdlViehstGlobalSrcprm;
}
