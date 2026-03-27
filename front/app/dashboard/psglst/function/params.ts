import {
  MdlPsglstErrlogSrcprm,
  MdlPsglstPsgdtlFrntnd,
  MdlPsglstPsgdtlSrcprm,
} from "../model/params";

// Treatment function params
export function FncPsglstPsgdtlSrcprm(params: MdlPsglstPsgdtlSrcprm, actdte: string[]) {
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
    isitfl_psgdtl: params.isitfl_psgdtl || "",
    isittx_psgdtl: params.isittx_psgdtl || "",
    isitir_psgdtl: params.isitir_psgdtl || "",
    nclear_psgdtl: params.nclear_psgdtl || "",
    format_psgdtl: params.format_psgdtl || "",
    pagenw_psgdtl: Number(params.pagenw_psgdtl) || 1,
    limitp_psgdtl: Number(params.limitp_psgdtl) || 15,
  } as MdlPsglstPsgdtlSrcprm;
}

// Treatment function params
export function FncPsglstErrlogSrcprm(params: MdlPsglstErrlogSrcprm) {
  return {
    update_global: params.update_global || "",
    erdvsn_errlog: params.erdvsn_errlog || "",
    pagenw_errlog: Number(params.pagenw_errlog) || 1,
    limitp_errlog: Number(params.limitp_errlog) || 5,
  } as MdlPsglstErrlogSrcprm;
}

export function FncPsglstRawdtaParams() {
  const fnlobj: MdlPsglstPsgdtlFrntnd = {
    mnfest: "",
    noterr: "",
    source: "",
    tktnfl: "",
    tktnvc: "",
    pnrcde: "",
    pnritl: "",
    ndayfl: "",
    datefl: 0,
    datevc: 0,
    daterv: 0,
    mnthfl: 0,
    timefl: 0,
    timerv: 0,
    timeis: 0,
    timecr: 0,
    airlfl: "",
    airlvc: "",
    airtyp: "",
    seatcn: "",
    flnbfl: "",
    flnbvc: "",
    flgate: "",
    depart: "",
    arrivl: "",
    routfl: "",
    routvc: "",
    routvf: "",
    routac: "",
    routmx: "",
    routsg: "",
    linenb: 0,
    ckinnb: 0,
    gender: "",
    typepx: "",
    seatpx: "",
    groupc: "",
    totpax: 0,
    segpnr: "",
    segtkt: "",
    psgrid: "",
    tourcd: "",
    staloc: "",
    stanbr: "",
    wrkloc: "",
    hmeloc: "",
    lniata: "",
    emplid: "",
    nmefst: "",
    nmelst: "",
    cpnbfl: 0,
    cpnbvc: 0,
    clssfl: "",
    clssvc: "",
    statvc: "",
    cbinfl: "",
    cbinvc: "",
    agtdie: "",
    agtdcr: "",
    codels: "",
    isitfl: "",
    isittx: "",
    isitir: "",
    isitct: "",
    isittf: "",
    isitnr: "",
    noteup: "",
    updtby: "",
    prmkey: "",

    // Ancillary
    gpcdae: "",
    sbcdae: "",
    descae: "",
    wgbgae: 0,
    qtbgae: 0,
    routae: "",
    currae: "",
    emdnae: "",

    // Bagtag
    nmbrbt: "",
    qntybt: 0,
    wghtbt: 0,
    paidbt: 0,
    fbavbt: 0,
    hfbabt: 0,
    qtotbt: 0,
    wtotbt: 0,
    ptotbt: 0,
    ftotbt: 0,
    excsbt: 0,
    typebt: "",
    coment: "",

    // Outbound
    airlob: "",
    flnbob: "",
    clssob: "",
    routob: "",
    dateob: 0,
    timeob: 0,

    // Inbound
    airlib: "",
    flnbib: "",
    clssib: "",
    dstrib: "",
    dateib: 0,
    timeib: 0,

    // Ireg
    codeir: "",
    airlir: "",
    flnbir: "",
    dateir: 0,

    // Infant
    tktnif: "",
    cpnbif: 0,
    dateif: 0,
    clssif: "",
    routif: "",
    statif: "",
    paxsif: "",

    // Cancel bagtag
    airlxt: "",
    dstrxt: "",
    nmbrxt: "",
  };

  return fnlobj;
}
