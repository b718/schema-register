type SubmitSchemaRequest = {
  name: string;
  version: string;
  schema: string;
};

export async function submitSchema(schema: SubmitSchemaRequest) {
  const response = await fetch("http://localhost:8080/insert-schemas", {
    method: "POST",
    body: JSON.stringify(schema),
  });

  if (!response.ok) {
    throw new Error("Failed to insert schema");
  }

  const responseData = await response.json();
  return responseData;
}
