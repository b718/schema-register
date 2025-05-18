export async function deleteSchema(schemaId: string) {
  const response = await fetch(
    `http://localhost:8080/delete-schema/${schemaId}`,
    {
      method: "DELETE",
    }
  );

  if (!response.ok) {
    throw new Error("Failed to delete schema");
  }

  return response.json();
}
