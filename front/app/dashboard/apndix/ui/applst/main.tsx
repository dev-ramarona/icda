import { mdlAllusrCookieObjson } from "../../../allusr/model/params";
import { ApiApndixApplstDtbase } from "../../api/dtbase";
import UixApndixApplstSelect from "./select";

export default async function UixApndixApplstMainpg({
  pagedb,
  cookie,
}: {
  pagedb: string;
  cookie: mdlAllusrCookieObjson;
}) {
  const apndix = await ApiApndixApplstDtbase();
  return (
    <>
      <UixApndixApplstSelect apndix={apndix} pagedb={pagedb} cookie={cookie} />
    </>
  );
}
