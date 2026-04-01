// Model accepted edit
export interface MdlApndixAcpedtDtbase {
  params: string;
  length: number;
  dvsion: string;
}

export interface MdlApndixSearchQueryx {
  update_apndix: string;
  pagedb_apndix: string;
  datefl_apndix: string;
  airlfl_apndix: string;
  depart_apndix: string;
  flnbfl_apndix: string;
  routfl_apndix: string;
  clssfl_apndix: string;
  pagenw_apndix: number;
  limitp_apndix: number;
}

// Flhour
export interface MdlApndixFlhourFrntnd {
  prmkey: string;
  airlfl: string;
  routfl: string;
  flnbfl: string;
  flhour: string;
  timefl: string;
  timerv: string;
  timeup: string;
  dateup: string;
  datend: string;
  airtyp: string;
  airmls: string;
  hstory: string;
  updtby: string;
}

// Provnc
export interface MdlApndixProvncFrntnd {
  prmkey: string;
  routfl: string;
  provnc: string;
  updtby: string;
}

// Provnc
export interface MdlApndixFrbaseFrntnd {
  prmkey: string;
  scdkey: string;
  airlfl: string;
  clssfl: string;
  routfl: string;
  frbcde: string;
  frbnta: string;
  frbsbr: string;
  datend: string;
  hstory: string;
  updtby: string;
}

// Frtaxes
export interface MdlApndixFrtaxsFrntnd {
  prmkey: string;
  airlfl: string;
  cbinfl: string;
  routfl: string;
  ftppnx: string;
  ftaptx: string;
  ftfuel: string;
  ftiwjr: string;
  ftaxyr: string;
  datend: string;
  ftothr: string;
  hstory: string;
  updtby: string;
}

// Fllist
export interface MdlApndixFllistFrntnd {
  prmkey: string;
  airlfl: string;
  flnbfl: string;
  timeup: number;
  timefl: number;
  timerv: number;
  datefl: number;
  mnthfl: number;
  ndayfl: string;
  flstat: string;
  routfl: string;
  routac: string;
  flsarr: string;
  routmx: string;
  flhour: number;
  flrpdc: number;
  flgate: string;
  depart: string;
  arrivl: string;
  airtyp: string;
  aircnf: string;
  seatcn: string;
  autrzc: number;
  autrzy: number;
  bookdc: number;
  bookdy: number;
}
