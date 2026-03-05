import { MdlAllusrSearchParams } from "../model/params";

// Treatment function params
export function FncAllusrSearchParams(params: MdlAllusrSearchParams) {
  return {
    usredt: params.usredt || "",
    usrnme: params.usrnme || "",
    stfnme: params.stfnme || "",
    stfeml: params.stfeml || "",
    pagenw: Number(params.pagenw) || 1,
    limitp: Number(params.limitp) || 5,
    update: params.update || "",
  } as MdlAllusrSearchParams;
}
