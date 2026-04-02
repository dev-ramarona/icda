import { MdlApndixSearchQueryx } from "../model/parmas";

// Treatment function params
export function FncApndixSearchQueryx(params: MdlApndixSearchQueryx, actdte: string[]) {
  return {
    pagedb_apndix: params.pagedb_apndix || "",
    datefl_apndix: params.datefl_apndix || actdte[actdte.length - 1],
    airlfl_apndix: params.airlfl_apndix || "",
    depart_apndix: params.depart_apndix || "",
    flnbfl_apndix: params.flnbfl_apndix || "",
    routfl_apndix: params.routfl_apndix || "",
    clssfl_apndix: params.clssfl_apndix || "",
    pagenw_apndix: Number(params.pagenw_apndix) || 1,
    limitp_apndix: Number(params.limitp_apndix) || 15,
  } as MdlApndixSearchQueryx;
}
