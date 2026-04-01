import { mdlGlobalAlluserFilter } from "../model/params";

export function FncGlobalFormatDfault(key: string, val: string): string | number {
  let now: string | number = "";
  const actkey = key.substring(0, 6);
  if (["routvc", "routfl"].includes(actkey)) {
    now = FncGlobalFormatRoutfl(val).substring(0, 7);
  } else if (
    ["airlfl", "airlvc", "depart", "curncy", "clssfl", "frbcde", "statvc"].includes(actkey)
  ) {
    const tmp = val.toUpperCase().replace(/[^A-Z]/g, "");
    if (["clssfl"].includes(actkey)) now = tmp.substring(0, 1);
    else if (["airlfl", "airlvc"].includes(actkey)) now = tmp.substring(0, 2);
    else if (["depart", "curncy"].includes(actkey)) now = tmp.substring(0, 3);
    else now = tmp;
  } else if (["flnbfl"].includes(actkey)) {
    now = val.replace(/[^0-9]/g, "").substring(0, 4);
  } else if (["tktnfl", "tktnvc"].includes(actkey)) {
    now = val.replace(/[^0-9]/g, "").substring(0, 13);
  } else if (["airmls", "ntafvc", "ntaffl", "cpnbvc", "frbnta", "frbsbr"].includes(actkey)) {
    now = val.replace(/[^0-9]/g, "");
  } else now = val;
  return now;
}

// Fucntion change format data yymmdd/hhmm to dd-MMM-yyyy hh:mm
export function FncGlobalFormatDatefm(inputd: string): string {
  if (inputd.length !== 6 && inputd.length !== 10 && inputd.length !== 4)
    return inputd == "0" ? "-" : inputd;
  const yearnw = inputd.slice(0, 2);
  const monthn = inputd.slice(2, 4);
  const daynow = inputd.length == 4 ? "01" : inputd.slice(4, 6);
  const yearfl = parseInt(yearnw) < 50 ? `20${yearnw}` : `19${yearnw}`;
  const hournw = inputd.length === 10 ? inputd.slice(6, 8) : null;
  const minute = inputd.length === 10 ? inputd.slice(8, 10) : null;

  // Buat Date object
  const datenw = new Date(`${yearfl}-${monthn}-${daynow}T${hournw ?? "00"}:${minute ?? "00"}`);
  let optons: Intl.DateTimeFormatOptions = {
    day: "2-digit",
    month: "short",
    year: "2-digit",
  };
  if (inputd.length === 4) {
    optons = {
      month: "short",
      year: "2-digit",
    };
  }
  const datetx = datenw.toLocaleDateString("en-GB", optons).replace(/ /g, "-");
  if (hournw && minute) return `${datetx} ${hournw}:${minute}`;
  return datetx;
}

// Function change format input date time to yymmddhhmm
export function FncGlobalFormatInptdt(v: string) {
  if (typeof v !== "string") return v;

  // datetime-local: "YYYY-MM-DDTHH:mm"
  if (v.includes("T")) {
    const [divide, timept] = v.split("T");
    const [yearnw, monthw, daynow] = divide.split("-");
    const [hournw, minute] = timept.split(":");
    return `${yearnw.slice(-2)}${monthw}${daynow}${hournw}${minute}`;
  }

  // date: "YYYY-MM-DD"
  else if (v.includes("-")) {
    const [yearnw, monthw, daynow] = v.split("-");
    return `${yearnw.slice(-2)}${monthw}${daynow}`;
  }

  // time: "HH:mm"
  else if (v.includes(":")) {
    const [hournw, minute] = v.split(":");
    return `${hournw}${minute}`;
  }

  return v; // fallback kalau format tidak dikenali
}

// Function get initial variable blank
export function FncGlobalIntialObject<o extends Record<string, any>>(obj: o): o {
  const result: any = {};
  if (obj) {
    Object.entries(obj).forEach(([key, val]) => {
      if (typeof val === "string") result[key] = "";
      else if (typeof val === "number") result[key] = 0;
      else if (typeof val === "boolean") result[key] = false;
      else result[key] = null;
    });
    return result;
  } else {
    const tmpobj: any = { empty: "" };
    return tmpobj;
  }
}

// // Fucntion change format data yymmdd/hhmm to dd-MMM-yyyy hh:mm
// export function FncGlobalFormatDateip(inputd: string): string {
//   if (inputd.length !== 6) return "Format harus YYMMDD";
//   const year = parseInt(inputd.slice(0, 2), 10) + 2000; // "25" → 2025
//   const month = inputd.slice(2, 4); // "07"
//   const day = inputd.slice(4, 6); // "30"
//   return `${year}-${month}-${day}`;
// }

// Function change format routef to 3-3 characters
export function FncGlobalFormatRoutfl(routef: string) {
  let raw = routef.toUpperCase().replace(/[^A-Z]/g, "");
  let rsl = "";
  for (let i = 0; i < raw.length; i += 3) {
    if (i > 0) rsl += "-";
    rsl += raw.slice(i, i + 3);
  }
  return rsl;
}

// Function change format routef to 3-3 characters
export function FncGlobalFormatPercnt(percent: string, prvprc: string) {
  if (!percent.includes("%")) {
    if (!prvprc.includes("%")) return percent + "%";
    else if (percent.length == 1) return "";
    return percent.substring(0, percent.length - 1) + "%";
  }
  const raw = percent.replace(/[^0-9]/g, "");
  return raw + "%";
}

// Function change format routef to 3-3 characters
export function FncGlobalFormatCpnfmt(cpnnbr: string) {
  if (cpnnbr === "") return cpnnbr;
  const raw = cpnnbr.toUpperCase().replace(/[^A-Z]/g, "");
  const nbr = parseInt(raw);
  if (isNaN(nbr)) return "";
  if (nbr < 10) {
    return `C0${nbr}`;
  }
  return `C${nbr}`;
}

// // Function change format routef to 3-3 characters
// export function FncGlobalFormatSorthl(params: string) {
//   const raw = params.trim().toUpperCase();
//   if (raw === "") return "";
//   if (raw.length < 3 && /^[LOW]/i.test(raw)) return "Lowest";
//   if (raw.length < 3 && /^[HIG]/i.test(raw)) return "Highest";
//   else if (raw.length === 1) return "Highest";
//   return ""; // default aman
// }

// // Function change format routef to 3-3 characters
// export function FncGlobalFormatFilter(params: string, arrays: mdlGlobalAlluserFilter[]) {
//   const raw = params.trim().toUpperCase();
//   if (raw === "") return "";
//   for (let i = 0; i < arrays.length; i++) {
//     const arr = arrays[i];
//     const reg = new RegExp(`^[${arr.keywrd}]`, "i");
//     if (raw.length <= arr.keywrd.length && reg.test(raw)) return arr.output;
//   }
//   if (raw.length === 1) return arrays[0].output;
//   return ""; // default aman
// }

// Function format arr split and cancel jeddah
export function FncGlobalFormatArrcpn(str: string) {
  if (!str.includes(":") && !str.includes("-")) return str;
  const val = str.split("|");
  const arr = [];
  for (let i = 0; i < val.length; i++) {
    const sep = val[i].includes(":") ? ":" : "-";
    const tmp = val[i].split(sep);
    const arx = [];
    for (let i = 0; i < tmp.length; i++) {
      const elm = tmp[i];
      if (Number.isInteger(Number(elm)) && elm.length >= 6) arx.push(FncGlobalFormatDatefm(elm));
      else arx.push(elm);
    }
    arr.push(arx.join("-"));
  }
  return arr.join(" | ");
}
