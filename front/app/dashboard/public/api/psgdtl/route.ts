import { NextRequest } from "next/server";

export async function POST(req: NextRequest) {
  const formData = await req.formData();
  const data = formData.get("data");

  const body = new URLSearchParams();
  body.append("data", String(data));

  const res = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVER}/psglst/psgdtl/getall/downld`, {
    method: "POST",
    body: body,
    headers: {
      "Content-Type": "application/x-www-form-urlencoded", // ✅ FIX
    },
  });

  return new Response(res.body, {
    headers: {
      "Content-Type": "text/csv",
      "Content-Disposition": "attachment; filename=data.csv",
    },
  });
}
