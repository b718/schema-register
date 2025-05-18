import type { FetchSchemasResponse } from "../../../types/schema-types";

export async function fetchAllSchemas(): Promise<FetchSchemasResponse> {
  const response = await fetch("http://localhost:8080/fetch-schemas");
  const data = await response.json();
  return data;
}
