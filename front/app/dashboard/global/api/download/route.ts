import { NextRequest } from "next/server";

export async function POST(req: NextRequest) {
  const form = await req.formData();
  const link = form.get("link");
  const data = form.get("data");
  console.log("Adawdw");

  const res = await fetch(link as string, {
    method: "POST",
    body: data,
    headers: { "Content-Type": "application/json" },
  });

  return new Response(res.body, {
    headers: res.headers,
  });
}
