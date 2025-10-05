import { type NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest) {
  const domain = process.env.AUTH0_ISSUER_BASE_URL;
  const clientId = process.env.AUTH0_CLIENT_ID;
  const redirectUri = process.env.AUTH0_BASE_URL + "/api/auth/callback";
  const audience = process.env.AUTH0_AUDIENCE;

  const authUrl = new URL(`${domain}/authorize`);
  authUrl.searchParams.set("response_type", "code");
  authUrl.searchParams.set("client_id", clientId!);
  authUrl.searchParams.set("redirect_uri", redirectUri);
  authUrl.searchParams.set("scope", "openid profile email");
  authUrl.searchParams.set("audience", audience!);

  return NextResponse.redirect(authUrl.toString());
}
