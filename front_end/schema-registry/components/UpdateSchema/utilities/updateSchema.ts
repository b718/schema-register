export async function updateSchema(schemaId: string, schemaDefinition: string) {
  const response = await fetch(
    `http://localhost:8080/update-schema/${schemaId}`,
    {
      method: "PUT",
      body: JSON.stringify({ schema: schemaDefinition }),
    }
  );

  if (!response.ok) {
    throw new Error("Failed to update schema");
  }

  return response.json();
}
