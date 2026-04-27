// Global
export interface MdlViehstGlobalSrcprm {
  update_global: string;
  datefl_prcess: number;
  airlfl_prcess: string;
  flnbfl_prcess: string;
  depart_prcess: string;
  worker_prcess: number;
  pagenw_viehst: number;
  limitp_viehst: number;
}

// Process
export interface MdlViehstPrcessSrcprm {
  datefl: number;
  airlfl: string;
  flnbfl: string;
  depart: string;
  worker: number;
}
