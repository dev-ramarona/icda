import { MdlApndixSearchQueryx } from "../model/parmas";

// Treatment function params
export function FncApndixSearchQueryx(params: MdlApndixSearchQueryx) {
  return {
    pagedb: params.pagedb || "",
    datefl: params.datefl || "",
    airlfl: params.airlfl || "",
    depart: params.depart || "",
    flnbfl: params.flnbfl || "",
    routfl: params.routfl || "",
    clssfl: params.clssfl || "",
    pagenw: Number(params.pagenw) || 1,
    limitp: Number(params.limitp) || 15,
  } as MdlApndixSearchQueryx;
}
