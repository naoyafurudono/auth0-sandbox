import { cookies } from "next/headers";
import { type NextRequest, NextResponse } from "next/server";

export async function GET(_request: NextRequest) {
  const cookieStore = await cookies();
  cookieStore.delete("access_token");
  cookieStore.delete("id_token");

  const domain = process.env.AUTH0_ISSUER_BASE_URL;
  const clientId = process.env.AUTH0_CLIENT_ID;
  const returnTo = process.env.AUTH0_BASE_URL;

  if (!clientId || !returnTo) {
    return NextResponse.json(
      { error: "Missing AUTH0_CLIENT_ID or AUTH0_BASE_URL" },
      { status: 500 },
    );
  }

  const logoutUrl = new URL(`${domain}/v2/logout`);
  logoutUrl.searchParams.set("client_id", clientId);
  logoutUrl.searchParams.set("returnTo", returnTo);

  return NextResponse.redirect(logoutUrl.toString());
}
