import { revalidatePath } from "next/cache";
import { NextRequest } from "next/server";

export async function POST(req: NextRequest) {
  const form = await req.formData();
  const link = form.get("link");
  const res = await fetch(link as string, {
    method: "POST",
    body: form,
  });
  revalidatePath("/", "layout");
  return new Response(res.body, {
    headers: res.headers,
  });
}
