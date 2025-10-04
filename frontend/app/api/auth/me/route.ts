import { NextResponse } from "next/server";
import { cookies } from "next/headers";

export async function GET() {
  const cookieStore = await cookies();
  const idToken = cookieStore.get("id_token")?.value;

  if (!idToken) {
    return new NextResponse(null, { status: 204 });
  }

  // IDトークンをデコードしてユーザー情報を取得
  const payload = JSON.parse(
    Buffer.from(idToken.split(".")[1], "base64").toString()
  );

  return NextResponse.json({
    sub: payload.sub,
    name: payload.name,
    email: payload.email,
    picture: payload.picture,
  });
}
