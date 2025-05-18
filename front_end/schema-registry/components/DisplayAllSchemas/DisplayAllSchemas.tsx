import { useState, useEffect } from "react";
import type { Schema, FetchSchemasResponse } from "../../types/schema-types";
import "../../styles/components.css";
import { fetchAllSchemas } from "./utilities/fetchAllSchemas";

const DisplayAllSchemas = () => {
  const [schemas, setSchemas] = useState<Schema[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    fetchAllSchemas()
      .then((data: FetchSchemasResponse) => setSchemas(data.schemas))
      .catch((error: Error) => setError(error.message))
      .finally(() => setLoading(false));
  }, []);

  const formatJSON = (jsonString: string) => {
    try {
      const parsed = JSON.parse(jsonString);
      return JSON.stringify(parsed, null, 2);
    } catch (e) {
      return jsonString;
    }
  };

  if (loading) return <div className="loading">Loading...</div>;
  if (error) return <div className="error">{error}</div>;

  return (
    <div className="schemas-grid">
      {schemas.length > 0 ? (
        schemas.map((schema) => (
          <div key={schema.id} className="schema-card">
            <div className="schema-header">
              <h3 className="schema-title">{schema.name}</h3>
              <span className="schema-id">ID: {schema.id}</span>
            </div>
            <div className="schema-content">
              <pre className="schema-json">
                <code>{formatJSON(schema.schema)}</code>
              </pre>
            </div>
          </div>
        ))
      ) : (
        <div className="no-schemas">No schemas found</div>
      )}
    </div>
  );
};

export default DisplayAllSchemas;
