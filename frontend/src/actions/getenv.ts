// "use server"
import { connection } from "next/server";

export default async function Getenv() {
  await connection();
  const environment = process.env.ENVIRONMENT;
  console.log(environment)
  const baseUrl =
    environment === "production"
      ? "https://jeagereats-production.up.railway.app/api/v1"
      : "http://localhost:8080/api/v1";
  return baseUrl;
}

