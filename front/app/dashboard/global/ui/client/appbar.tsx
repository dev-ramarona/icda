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
  const divref = useRef<HTMLDivElement>(null)
  const [lstpth, lstpthSet] = useState("xxxxxx");
  const [onclik, onclikSet] = useState(false);
  const [onhide, onhideSet] = useState(false);
  const [scroll, scrollSet] = useState(true)
  useEffect(() => {
    const segment = pthnme.split("/").filter(Boolean).pop();
    lstpthSet(segment || "");
  }, [pthnme]);
  useEffect(() => {
    const scrollFnc = () => {
      if (window.scrollY == 0) {
        scrollSet(true)
      } else scrollSet(false)
    }
    const mousecFnc = (event: MouseEvent) => {
      if (divref.current && !divref.current.contains(event.target as Node))
        onclikSet(false)
    }
    window.addEventListener("scroll", scrollFnc)
    document.addEventListener("mousedown", mousecFnc)
    return () => {
      window.removeEventListener("scroll", scrollFnc)
      document.removeEventListener("mousedown", mousecFnc)
    }
  }, [])


  return (
    <div className="fixed top-0 z-20" ref={divref}>
      <div
        className={`${scroll || onclik ? "w-10 h-10" : "w-12 h-12"} p-1.5 absolute left-0 duration-300 flexctr cursor-pointer`}
        onClick={() => {
          onclikSet(!onclik);
          onhideSet(false);
        }}
      >
        <div className="afull rounded-md z-10 relative flexstr">
          <div className={`absolute ${scroll || onclik ? "w-36 opacity-100" : "w-full opacity-0"} h-full bg-cyan-600 flexstr whitespace-nowrap pl-10 text-white duration-300 ease-in-out font-semibold rounded-2xl`}>
            <div>IC Data Analyst</div>
            <div className={`w-6 ${onclik ? "translate-x-8 opacity-100" : "translate-x-0 opacity-0"} h-full absolute right-0 flexctr flex-col gap-1 duration-300`}>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onclik ? "-translate-x-2 -translate-y-2" : "-translate-y-2"} duration-300`}></div>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onclik ? "translate-x-2 -translate-y-2" : ""} duration-300`}></div>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onclik ? "" : ""}`}></div>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onclik ? "-translate-x-2 translate-y-2" : ""} duration-300`}></div>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onclik ? "translate-x-2 translate-y-2" : "translate-y-2"} duration-300`}></div>
            </div>
          </div>
          <div className={`absolute afull bg-cyan-600 p-1 rounded-md ${scroll || onclik ? "opacity-100" : "opacity-50 hover:opacity-100"} rounded-xl duration-300`}>
            <Image
              className="invert"
              src="/lionairblack.png"
              width={1000}
              height={1000}
              alt=""
            />
          </div>
        </div>
      </div>
      <div className={`${onclik ? "w-screen md:w-64 opacity-100" : "w-0 md:w-0 opacity-0"} h-screen bg-white shadow-lg shadow-gray-600 duration-300 ease-in-out py-12 overflow-hidden`}>
        <div className="afull flexctr flex-col gap-y-1">
          {applst.map((val, idx) => (
            <Link className={`w-full h-9 pl-6 flexctr gap-1.5 text-gray-700 ${lstpth == val.prmkey ? "bg-cyan-600 hover:bg-cyan-700 h-12 text-white font-semibold tracking-wider" : "hover:bg-cyan-100"} group duration-300`} href={`/dashboard/${val.prmkey}`} key={idx} onClick={() => onclikSet(!onclik)}>
              <div className={`${lstpth == val.prmkey ? "rotate-45 group-hover:scale-110 group-hover:-translate-x-1" : "group-hover:rotate-45"} duration-300`}>
                <UixGlobalIconvcTolink bold={2} color={`${lstpth == val.prmkey ? "#ffffff" : "#4a5565"}`} size={1.2} />
              </div>
              <div className="afull text-sm whitespace-nowrap flexstr">{val.detail}</div>
            </Link>
          ))}
          <div className="absolute bottom-5 w-full h-10 pl-3 pr-6 flexbtw group cursor-pointer" onClick={() => onhideSet(!onhide)}>
            <div className="flexctr gap-1.5">
              <div className="w-7 h-7 bg-cyan-600 rounded-lg flexctr"><UixGlobalIconvcProfle bold={3} color="#fff" size={1.4} /></div>
              <div>
                <div className="text-sm">{cookie.usrnme}</div>
                <div className="text-xs text-slate-500">{cookie.stfeml}</div>
              </div>
            </div>
            <div className="flexctr flex-col gap-1">
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onhide ? "-translate-x-2 -translate-y-2" : "-translate-y-2"} duration-300`}></div>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onhide ? "translate-x-2 -translate-y-2" : ""} duration-300`}></div>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onhide ? "" : ""}`}></div>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onhide ? "-translate-x-2 translate-y-2" : ""} duration-300`}></div>
              <div className={`absolute h-1 w-1 bg-cyan-600 rounded-full ${onhide ? "translate-x-2 translate-y-2" : "translate-y-2"} duration-300`}></div>
            </div>
            <div className={`absolute w-32 h-48 max-h-fit bg-white right-0 ${onhide ? "-translate-x-1/3 md:translate-x-11/12 -translate-y-1/2 opacity-100" : "opacity-0 select-none pointer-events-none"} ring-2 ring-gray-300 rounded-md duration-300 ease-in-out`}>
              <div className="py-1.5">
                <Link className="afull flexstr hover:bg-cyan-100 duration-300 px-1.5 py-1.5 cursor-pointer" href={"/dashboard/allusr"} onClick={() => onclikSet(!onclik)}>
                  <div className="pr-1 flexctr">
                    <UixGlobalIconvcUsrdtl bold={2} color="#6a7282" size={1.4} />
                  </div>
                  <div>Create user</div>
                </Link>
              </div>
              <hr className="text-gray-300" />
              <form className="py-1.5" action={ApiAllusrCookieLogout}>
                <button className="afull flexstr hover:bg-cyan-100 duration-300 px-1.5 py-1.5 cursor-pointer" type="submit">
                  <div className="pr-1 flexctr">
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
