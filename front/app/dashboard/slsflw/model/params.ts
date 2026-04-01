// Global
export interface MdlSlsflwGlobalSrcprm {
  update_global: string;
  mnthfl_psgdtl: string;
  datefl_psgdtl: string;
  airlfl_psgdtl: string;
  flnbfl_psgdtl: string;
  depart_psgdtl: string;
  routfl_psgdtl: string;
  pnrcde_psgdtl: string;
  tktnfl_psgdtl: string;
  nclear_psgdtl: string;
  format_psgdtl: string;
  keywrd_psgdtl: string;
  pagenw_psgdtl: number;
  limitp_psgdtl: number;
  pagenw_errlog: number;
  limitp_errlog: number;
  erdvsn_errlog: string;
  mnthfl_psgsmr: string;
  datefl_psgsmr: string;
  airlfl_psgsmr: string;
  flnbfl_psgsmr: string;
  depart_psgsmr: string;
  routfl_psgsmr: string;
  pagenw_psgsmr: number;
  limitp_psgsmr: number;
}

// Log action
export interface MdlSlsflwActlogDtbase {
  timeup: number;
  dateup: number;
  datefl: number;
  statdt: string;
}

// Search param
export interface MdlSlsflwAcpedtDtbase {
  params: string;
  length: number;
  dvsion: string;
}

// error Log
export interface MdlSlsflwErrlogDtbase {
  update_global: string;
  pagenw_errlog: number;
  limitp_errlog: number;
  erdvsn_errlog: string;
}

// Passangger list summary
export interface MdlSlsflwPsgsmrSrcprm {
  update_global: string;
  mnthfl_psgsmr: string;
  datefl_psgsmr: string;
  airlfl_psgsmr: string;
  flnbfl_psgsmr: string;
  depart_psgsmr: string;
  routfl_psgsmr: string;
  keywrd_psgsmr: string;
  pagenw_psgsmr: number;
  limitp_psgsmr: number;
}
export interface MdlSlsflwPsgsmrFrntnd {
  prmkey: string;
  airlfl: string;
  provnc: string;
  depart: string;
  flnbfl: string;
  routfl: string;
  ndayfl: string;
  datefl: number;
  mnthfl: number;
  flstat: string;
  seatcn: string;
  airtyp: string;
  flhour: number;
  totnta: number;
  tottyq: number;
  totpax: number;
  totfae: number;
  totqfr: number;
  totrph: number;
}

// Passangger list detail
export interface MdlSlsflwPsgdtlSrcprm {
  update_global: string;
  mnthfl_psgdtl: string;
  datefl_psgdtl: string;
  airlfl_psgdtl: string;
  flnbfl_psgdtl: string;
  depart_psgdtl: string;
  routfl_psgdtl: string;
  pnrcde_psgdtl: string;
  tktnfl_psgdtl: string;
  nclear_psgdtl: string;
  format_psgdtl: string;
  keywrd_psgdtl: string;
  pagenw_psgdtl: number;
  limitp_psgdtl: number;
}
export interface MdlSlsflwPsgdtlFrntnd {
  mnfest: string;
  slsrpt: string;
  noterr: string;
  source: string;
  tktnfl: string;
  tktnvc: string;
  pnrcde: string;
  pnritl: string;
  curncy: string;
  ntaffl: number;
  ntafvc: number;
  yqtxfl: number;
  yqtxvc: number;
  frrate: number;
  frbcde: string;
  qsrcrw: string;
  qsrcvc: number;
  frcalc: string;
  ndayfl: string;
  datefl: number;
  datevc: number;
  daterv: number;
  mnthfl: number;
  timefl: number;
  timerv: number;
  timeis: number;
  timecr: number;
  airlfl: string;
  airlvc: string;
  airtyp: string;
  seatcn: string;
  flhour: number;
  flnbfl: string;
  flnbvc: string;
  flgate: string;
  bookdc: number;
  bookdy: number;
  depart: string;
  arrivl: string;
  routfl: string;
  routvc: string;
  routvf: string;
  routac: string;
  routmx: string;
  routfr: string;
  routfx: string;
  routsg: string;
  linenb: number;
  ckinnb: number;
  gender: string;
  typepx: string;
  seatpx: string;
  groupc: string;
  totpax: number;
  segpnr: string;
  segtkt: string;
  psgrid: string;
  tourcd: string;
  staloc: string;
  stanbr: string;
  wrkloc: string;
  hmeloc: string;
  lniata: string;
  emplid: string;
  cpnbfl: number;
  cpnbvc: number;
  clssfl: string;
  clssvc: string;
  statvc: string;
  cbinfl: string;
  cbinvc: string;
  agtdie: string;
  agtdcr: string;
  codels: string;
  isitfl: string;
  isittx: string;
  isitir: string;
  isitct: string;
  isittf: string;
  isitnr: string;
  noteup: string;
  updtby: string;
  prmkey: string;

  // Ancillary
  gpcdae: string;
  sbcdae: string;
  descae: string;
  wgbgae: number;
  qtbgae: number;
  routae: string;
  fareae: number;
  currae: string;
  emdnae: string;
}
