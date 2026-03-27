"use client";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useRef, useState } from "react";
import {
  UixGlobalIconvcLogout,
  UixGlobalIconvcProfle,
  UixGlobalIconvcTolink,
  UixGlobalIconvcUsrdtl,
} from "../server/iconvc";
import { MdlAllusrApplstParams, mdlAllusrCookieObjson } from "../../../allusr/model/params";
import { ApiAllusrCookieLogout } from "../../../allusr/api/cookie";

export default function UixGlobalAppbarClient({
  cookie,
  applst,
}: {
  cookie: mdlAllusrCookieObjson;
  applst: MdlAllusrApplstParams[];
}) {
  const pthnme = usePathname();
  const divref = useRef<HTMLDivElement>(null);
  const [lstpth, lstpthSet] = useState("xxxxxx");
  const [onclik, onclikSet] = useState(false);
  const [onhide, onhideSet] = useState(false);
  const [scroll, scrollSet] = useState(true);
  const [random, randomSet] = useState("");
  useEffect(() => {
    const segment = pthnme.split("/").filter(Boolean).pop();
    lstpthSet(segment || "");
  }, [pthnme]);
  useEffect(() => {
    randomSet(String(Math.random()));
    const scrollFnc = () => {
      if (window.scrollY == 0) {
        scrollSet(true);
      } else scrollSet(false);
    };
    const mousecFnc = (event: MouseEvent) => {
      if (divref.current && !divref.current.contains(event.target as Node)) onclikSet(false);
    };
    window.addEventListener("scroll", scrollFnc);
    document.addEventListener("mousedown", mousecFnc);
    return () => {
      window.removeEventListener("scroll", scrollFnc);
      document.removeEventListener("mousedown", mousecFnc);
    };
  }, []);

  return (
    <div className="fixed top-0 z-30" ref={divref}>
      <div
        className={`${scroll || onclik ? "h-10 w-10" : "h-12 w-12"} flexctr absolute left-0 cursor-pointer p-1.5 duration-300`}
        onClick={() => {
          onclikSet(!onclik);
          onhideSet(false);
        }}
      >
        <div className="afull flexstr relative z-10 rounded-md">
          <div
            className={`absolute ${scroll || onclik ? "w-36 opacity-100" : "w-full opacity-0"} flexstr h-full rounded-2xl bg-cyan-600 pl-10 font-semibold whitespace-nowrap text-white duration-300 ease-in-out`}
          >
            <div>IC Data Analyst</div>
            <div
              className={`w-6 ${onclik ? "translate-x-8 opacity-100" : "translate-x-0 opacity-0"} flexctr absolute right-0 h-full flex-col gap-1 duration-300`}
            >
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onclik ? "-translate-x-2 -translate-y-2" : "-translate-y-2"} duration-300`}
              ></div>
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onclik ? "translate-x-2 -translate-y-2" : ""} duration-300`}
              ></div>
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onclik ? "" : ""}`}
              ></div>
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onclik ? "-translate-x-2 translate-y-2" : ""} duration-300`}
              ></div>
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onclik ? "translate-x-2 translate-y-2" : "translate-y-2"} duration-300`}
              ></div>
            </div>
          </div>
          <div
            className={`afull absolute rounded-md bg-cyan-600 p-1 ${scroll || onclik ? "opacity-100" : "opacity-50 hover:opacity-100"} rounded-xl duration-300`}
          >
            <Image className="invert" src="/lionairblack.png" width={1000} height={1000} alt="" />
          </div>
        </div>
      </div>
      <div
        className={`${onclik ? "w-screen opacity-100 md:w-64" : "w-0 opacity-0 md:w-0"} h-screen overflow-hidden bg-white py-12 shadow-lg shadow-gray-600 duration-300 ease-in-out`}
      >
        <div className="afull flexctr flex-col gap-y-1">
          {applst.map((val, idx) => (
            <Link
              className={`flexctr h-9 w-full gap-1.5 pl-6 text-gray-700 ${lstpth == val.prmkey ? "h-12 bg-cyan-600 font-semibold tracking-wider text-white hover:bg-cyan-700" : "hover:bg-cyan-100"} group duration-300`}
              href={`/dashboard/${val.prmkey}?update_global=${random}&update=${random}`}
              key={idx}
              onClick={() => onclikSet(!onclik)}
            >
              <div
                className={`${lstpth == val.prmkey ? "rotate-45 group-hover:-translate-x-1 group-hover:scale-110" : "group-hover:rotate-45"} duration-300`}
              >
                <UixGlobalIconvcTolink
                  bold={2}
                  color={`${lstpth == val.prmkey ? "#ffffff" : "#4a5565"}`}
                  size={1.2}
                />
              </div>
              <div className="afull flexstr text-sm whitespace-nowrap">{val.detail}</div>
            </Link>
          ))}
          <div
            className="flexbtw group absolute bottom-5 h-10 w-full cursor-pointer pr-6 pl-3"
            onClick={() => onhideSet(!onhide)}
          >
            <div className="flexctr gap-1.5">
              <div className="flexctr h-7 w-7 rounded-lg bg-cyan-600">
                <UixGlobalIconvcProfle bold={3} color="#fff" size={1.4} />
              </div>
              <div>
                <div className="text-sm">{cookie.usrnme}</div>
                <div className="text-xs text-slate-500">{cookie.stfeml}</div>
              </div>
            </div>
            <div className="flexctr flex-col gap-1">
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onhide ? "-translate-x-2 -translate-y-2" : "-translate-y-2"} duration-300`}
              ></div>
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onhide ? "translate-x-2 -translate-y-2" : ""} duration-300`}
              ></div>
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onhide ? "" : ""}`}
              ></div>
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onhide ? "-translate-x-2 translate-y-2" : ""} duration-300`}
              ></div>
              <div
                className={`absolute h-1 w-1 rounded-full bg-cyan-600 ${onhide ? "translate-x-2 translate-y-2" : "translate-y-2"} duration-300`}
              ></div>
            </div>
            <div
              className={`absolute right-0 h-48 max-h-fit w-32 bg-white ${onhide ? "-translate-x-1/3 -translate-y-1/2 opacity-100 md:translate-x-11/12" : "pointer-events-none opacity-0 select-none"} rounded-md ring-2 ring-gray-300 duration-300 ease-in-out`}
            >
              <div className="py-1.5">
                <Link
                  className="afull flexstr cursor-pointer px-1.5 py-1.5 duration-300 hover:bg-cyan-100"
                  href={"/dashboard/allusr"}
                  onClick={() => onclikSet(!onclik)}
                >
                  <div className="flexctr pr-1">
                    <UixGlobalIconvcUsrdtl bold={2} color="#6a7282" size={1.4} />
                  </div>
                  <div>Create user</div>
                </Link>
              </div>
              <hr className="text-gray-300" />
              <form className="py-1.5" action={ApiAllusrCookieLogout}>
                <button
                  className="afull flexstr cursor-pointer px-1.5 py-1.5 duration-300 hover:bg-cyan-100"
                  type="submit"
                >
                  <div className="flexctr pr-1">
                    <UixGlobalIconvcLogout bold={2} color="#6a7282" size={1.4} />
                  </div>
                  <div>Log out</div>
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
