import Link from "next/link";
import { UixGlobalIconvcTolink } from "./ui/server/iconvc";
import { ApiAllusrCookieGetdta } from "../allusr/api/cookie";

export default async function Page() {
  const cookie = await ApiAllusrCookieGetdta();
  const random = String(Math.random());
  return (
    <div className="afull flexctr fixed top-0 flex-col text-sky-900">
      <div className="text-3xl">
        Wellcome <span className="font-semibold">{cookie.stfnme}</span>
      </div>
      <div>You only accepted on Page</div>
      <div className="flexctr w-2/3 flex-wrap md:w-1/3">
        {cookie.access.map((item, index) => (
          <Link
            href={"/dashboard/" + item + `?update_global=${random}&update=${random}`}
            className="flexctr group text-base"
            key={index}
          >
            <div className="duration-300 group-hover:scale-110 group-hover:rotate-45">
              <UixGlobalIconvcTolink color="#024a70" size={1.1} bold={2} />
            </div>
            <div className="font-semibold">{item.toUpperCase()}</div>
            <div className="pr-2 pl-0">,</div>
          </Link>
        ))}
      </div>
      <div className="flexctr w-full flex-col py-5 text-center">
        <div>for request Access or new User please confirm to email :</div>
        <div className="font-semibold">rama.rona@lionair.com</div>
      </div>
    </div>
  );
}
