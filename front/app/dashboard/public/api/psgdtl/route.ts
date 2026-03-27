import { NextRequest } from "next/server";

export async function POST(req: NextRequest) {
  const formData = await req.formData();
  const data = formData.get("data");

  const res = await fetch(`${process.env.URL_SERVER}/psglst/psgdtl/getall/downld`, {
    method: "POST",
    body: JSON.stringify({ data }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  const blob = await res.blob();

  return new Response(blob, {
    headers: {
      "Content-Type": "text/csv",
      "Content-Disposition": "attachment; filename=data.csv",
    },
  });
}
