// Allusr
export interface mdlGlobalAllusrCookie {
  stfnme: string;
  usrnme: string;
  stfeml: string;
  access: string[];
  keywrd: string[];
}
export interface mdlGlobalAlluserFilter {
  keywrd: string;
  output: string;
}
export interface mdlGlobalAlluserStatus {
  keywrd: string;
  output: string;
}

// Status data
export interface MdlGlobalStatusPrcess {
  sbrapi: number;
  action: number;
}

// Model accepted edit
export interface MdlGlobalAcpedtDtbase {
  params: string;
  length: number;
  dvsion: string;
}

// Action log
export interface MdlGlobalActlogDtbase {
  timeup: number;
  dateup: number;
  datefl: number;
  statdt: string;
}
