export interface MdlAllusrSearchParams {
  usredt: string;
  stfnme: string;
  usrnme: string;
  stfeml: string;
  limitp: number;
  pagenw: number;
  update: string;
}
export interface MdlAllusrFormipParams {
  prmkey: string;
  stfnme: string;
  usrnme: string;
  stfeml: string;
  psswrd: string;
  access: string[];
  keywrd: string[];
}
export interface MdlAllusrFrntndParams {
  prmkey: string;
  stfnme: string;
  usrnme: string;
  stfeml: string;
  psswrd: string;
  access: string[];
  keywrd: string[];
}
export interface MdlAllusrApplstParams {
  pagenb: number;
  prmkey: string;
  detail: string;
}
export interface mdlAllusrCookieObjson {
  stfnme: string;
  usrnme: string;
  stfeml: string;
  access: string[];
  keywrd: string[];
}
export interface MdlAllusrStatusPrcess {
  sbrapi: number;
  action: string;
}
