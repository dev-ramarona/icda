// Model accepted edit
export interface MdlApndixAcpedtDtbase {
  params: string;
  length: number;
  dvsion: string;
}

export interface MdlApndixSearchQueryx {
  update: string;
  pagedb: string;
  datefl: string;
  airlfl: string;
  depart: string;
  flnbfl: string;
  routfl: string;
  clssfl: string;
  pagenw: number;
  limitp: number;
}

// Flhour
export interface MdlApndixFlhourDtbase {
  prmkey: string;
  airlfl: string;
  routfl: string;
  flnbfl: string;
  flhour: number;
  timefl: number;
  timerv: number;
  timeup: number;
  dateup: number;
  datend: number;
  airtyp: string;
  airmls: number;
  hstory: string;
  updtby: string;
}

// Provnc
export interface MdlApndixProvncDtbase {
  prmkey: string;
  routfl: string;
  provnc: string;
  updtby: string;
}
