import axios from "axios";
import { NextRequest, NextResponse } from "next/server";
import { Roles } from "./types";

export async function middleware(request: NextRequest) {
  let userId = request.cookies.get("jeagereats_user_id");
  let accessToken = request.cookies.get("jeagereats_token");
  const baseurl =
    process.env.ENVIRONMENT == "production"
      ? "https://jeagereats-production.up.railway.app/api/v1"
      : "http://localhost:8080/api/v1";

  if (!userId || !accessToken) {
    return NextResponse.redirect(new URL("/auth/login", request.url));
  }

  try {
    let res = await axios.get(`${baseurl}/users`, {
      headers: {
        Authorization: `Bearer ${accessToken.value}`,
      },
    });
    const role: Roles = res?.data?.role;
    if (role && !request.url.includes(role)) {
      return NextResponse.redirect(new URL("/auth/login", request.url));
    }
  } catch (error) {
    return NextResponse.redirect(new URL("/auth/login", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: "/dashboard/:path*",
};
