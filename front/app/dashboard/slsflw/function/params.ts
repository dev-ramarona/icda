import {
  MdlSlsflwErrlogDtbase,
  MdlSlsflwPsgdtlSrcprm,
  MdlSlsflwPsgsmrSrcprm,
} from "../model/params";

// Treatment function params
export function FncSlsflwPsgdtlSrcprm(
  params: MdlSlsflwPsgdtlSrcprm,
  actdte: string[],
) {
  return {
    update_global: params.update_global || "",
    mnthfl_psgdtl: params.mnthfl_psgdtl || "",
    datefl_psgdtl: params.datefl_psgdtl || actdte[actdte.length - 1],
    airlfl_psgdtl: params.airlfl_psgdtl || "",
    flnbfl_psgdtl: params.flnbfl_psgdtl || "",
    depart_psgdtl: params.depart_psgdtl || "",
    routfl_psgdtl: params.routfl_psgdtl || "",
    pnrcde_psgdtl: params.pnrcde_psgdtl || "",
    tktnfl_psgdtl: params.tktnfl_psgdtl || "",
    nclear_psgdtl: params.nclear_psgdtl || "",
    pagenw_psgdtl: Number(params.pagenw_psgdtl) || 1,
    limitp_psgdtl: Number(params.limitp_psgdtl) || 15,
  } as MdlSlsflwPsgdtlSrcprm;
}

// Treatment function params
export function FncSlsflwPsgsmrSrcprm(
  params: MdlSlsflwPsgsmrSrcprm,
  actdte: string[],
) {
  return {
    update_global: params.update_global || "",
    mnthfl_psgsmr: params.update_global || "",
    datefl_psgsmr: params.datefl_psgsmr || actdte[actdte.length - 1],
    airlfl_psgsmr: params.airlfl_psgsmr || "",
    flnbfl_psgsmr: params.flnbfl_psgsmr || "",
    depart_psgsmr: params.depart_psgsmr || "",
    routfl_psgsmr: params.routfl_psgsmr || "",
    pagenw_psgsmr: Number(params.pagenw_psgsmr) || 1,
    limitp_psgsmr: Number(params.limitp_psgsmr) || 15,
  } as MdlSlsflwPsgsmrSrcprm;
}

// Treatment function params
export function FncSlsflwErrlogSrcprm(params: MdlSlsflwErrlogDtbase) {
  return {
    update_global: params.update_global || "",
    erdvsn_errlog: params.erdvsn_errlog || "",
    pagenw_errlog: Number(params.pagenw_errlog) || 1,
    limitp_errlog: Number(params.limitp_errlog) || 5,
  } as MdlSlsflwErrlogDtbase;
}
