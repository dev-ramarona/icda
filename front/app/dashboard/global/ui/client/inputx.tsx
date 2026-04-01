"use client";
import { useEffect, useRef, useState } from "react";
export default function UixGlobalInputxFormdt({
  typipt,
  labelx,
  length,
  queryx,
  params,
  plchdr,
  repprm,
}: {
  typipt: "text" | "date" | "month" | "number" | "file" | "datetime-local" | "select";
  labelx: "hidden" | "";
  length: number | null | string[];
  queryx: string;
  params: string;
  plchdr: string;
  repprm: null | ((e: React.ChangeEvent<HTMLInputElement>) => void);
}) {
  const refdte = useRef<HTMLInputElement>(null);
  const divref = useRef<HTMLDivElement>(null);
  const pckrdt = () => refdte.current?.showPicker();
  const [onclik, onclikSet] = useState(false);
  const [zindex, zindexSet] = useState(false);
  useEffect(() => {
    const mousecFnc = (event: MouseEvent) =>
      divref.current && !divref.current.contains(event.target as Node) && onclikSet(false);
    document.addEventListener("mousedown", mousecFnc);
    return () => document.removeEventListener("mousedown", mousecFnc);
  }, []);

  const refidx = useRef<NodeJS.Timeout | null>(null);
  useEffect(() => {
    if (refidx.current) clearTimeout(refidx.current);
    if (onclik) zindexSet(true);
    else refidx.current = setTimeout(() => zindexSet(false), 300);
  }, [onclik]);
  return (
    <div
      className={`afull flexstr relative px-1 py-1.5 text-[0.6rem] md:text-[0.66rem] ${zindex && "z-20"}`}
    >
      {typipt == "file" ? (
        // Input type file
        <>
          <input
            className="peer hidden"
            type={typipt}
            id={queryx}
            name={queryx}
            accept=".csv"
            multiple
            hidden={labelx == "hidden" ? true : false}
            onChange={(e) => (repprm ? repprm(e) : "")}
          />
          <label
            className={`afull flex cursor-pointer items-center rounded-md bg-white p-3 ring-2 ring-gray-200 ${
              params != ""
                ? "overflow-hidden text-left text-[0.5rem] whitespace-nowrap text-slate-700"
                : "text-white peer-focus:text-slate-500"
            } duration-300`}
            htmlFor={queryx}
          >
            {params}
          </label>
          <label
            className={`absolute cursor-pointer p-3 whitespace-nowrap select-none ${
              params != ""
                ? `mb-1 h-1/2 -translate-y-full text-[0.65rem] font-semibold text-slate-600 ${
                    labelx == "hidden" && "opacity-0"
                  }`
                : `text-slate-400 peer-focus:mb-1 peer-focus:h-1/2 peer-focus:-translate-y-full ${
                    labelx == "hidden" && "peer-focus:opacity-0"
                  }`
            } group/fst duration-300`}
            htmlFor={queryx}
          >
            <div className="flexctr">
              <div className="cursor-pointer rounded-md bg-white">{plchdr}</div>
            </div>
          </label>
        </>
      ) : typipt == "select" ? (
        // Input type file
        <>
          <div
            className={`afull flexctr relative cursor-pointer rounded-md bg-white shadow-md ring-gray-200 ${
              onclik ? "ring-0" : "ring-2"
            } duration-300`}
            ref={divref}
          >
            <div className="afull flexstr p-1.5" onClick={() => !onclik && onclikSet(!onclik)}>
              {params}
            </div>
            <div
              className={`absolute left-3 cursor-text whitespace-nowrap select-none ${
                params != ""
                  ? `mb-1 h-1/2 -translate-y-[150%] text-[0.65rem] font-semibold text-slate-600`
                  : `text-slate-400 ${onclik && "mb-1 h-1/2 -translate-y-[150%]"}`
              } group/fst duration-300`}
              onClick={() => !onclik && onclikSet(!onclik)}
            >
              <div className="flexctr">
                <div className="cursor-pointer rounded-md bg-white">{plchdr}</div>
              </div>
            </div>
            <div
              className={`absolute top-0 w-full rounded-md bg-white ring-2 ring-gray-200 ${
                onclik ? "h-[400%] opacity-100" : "h-0 opacity-0"
              } overflow-auto duration-300`}
            >
              {Array.isArray(length) &&
                length.map((val, key) => (
                  <div className="h-1/3 p-1.5" key={key}>
                    <input
                      value={val}
                      name={queryx}
                      id={queryx + key}
                      type="radio"
                      onChange={(e) => (repprm ? repprm(e) : "")}
                      hidden
                    />
                    <label
                      className={`afull flexctr rounded-md ${params == val ? "btnsbm" : "btnstb"}`}
                      htmlFor={queryx + key}
                      onClick={() => onclikSet(false)}
                    >
                      {val}
                    </label>
                  </div>
                ))}
              <div className="h-1/3 p-1.5">
                <input
                  value={""}
                  name={queryx}
                  id={"reset" + queryx}
                  type="radio"
                  onChange={(e) => (repprm ? repprm(e) : "")}
                  hidden
                />
                <label
                  className="afull flexctr btncxl rounded-md"
                  htmlFor={"reset" + queryx}
                  onClick={() => onclikSet(false)}
                >
                  Reset
                </label>
              </div>
            </div>
          </div>
        </>
      ) : (
        // Input type All
        <>
          <input
            className={`afull peer rounded-md bg-white p-1.5 shadow-md ${
              params != "" ? "text-slate-700" : "text-white focus:text-slate-500"
            } ring-2 ring-gray-200 duration-300`}
            value={params}
            maxLength={length && typeof length == "number" ? length : undefined}
            min={Array.isArray(length) ? length[0] : length}
            max={Array.isArray(length) ? length[length.length - 1] : length}
            type={typipt}
            id={queryx}
            name={queryx}
            onChange={(e) => (repprm ? repprm(e) : "")}
            onClick={() =>
              typipt == "date" || typipt == "month" || typipt == "datetime-local" ? pckrdt() : null
            }
            ref={
              typipt == "date" || typipt == "month" || typipt == "datetime-local" ? refdte : null
            }
          />
          <label
            className={`absolute left-3 cursor-text whitespace-nowrap select-none ${
              params != ""
                ? `mb-1 h-1/2 -translate-y-full text-[0.65rem] font-semibold text-slate-600 ${
                    labelx == "hidden" && "opacity-0"
                  }`
                : `text-slate-400 peer-focus:mb-1 peer-focus:h-1/2 peer-focus:-translate-y-full ${
                    labelx == "hidden" && "peer-focus:opacity-0"
                  }`
            } group/fst duration-300`}
            htmlFor={queryx}
          >
            <div className="flexctr">
              <div className="rounded-md bg-white px-1">{plchdr}</div>
            </div>
          </label>
        </>
      )}
    </div>
  );
}
