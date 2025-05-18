import React, { useState } from "react";
import Modal from "../Modal/Modal";
import "../../styles/components.css";
import { submitSchema } from "./utilities/submitSchema";

interface FormData {
  name: string;
  version: string;
  schema: string;
}

interface AddNewSchemaProps {
  isOpen: boolean;
  onClose: () => void;
}

const initialFormData: FormData = {
  name: "",
  version: "",
  schema: "",
};

const AddNewSchema: React.FC<AddNewSchemaProps> = ({ isOpen, onClose }) => {
  const [formData, setFormData] = useState<FormData>(initialFormData);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isValidJSON, setIsValidJSON] = useState(true);

  const validateJSON = (jsonString: string): boolean => {
    try {
      JSON.parse(jsonString);
      return true;
    } catch (e) {
      return false;
    }
  };

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));

    // Validate JSON when schema field changes
    if (name === "schema") {
      setIsValidJSON(validateJSON(value));
      setError(null);
    }
  };

  const formatJSON = () => {
    try {
      const parsed = JSON.parse(formData.schema);
      const formatted = JSON.stringify(parsed, null, 2);
      setFormData((prev) => ({ ...prev, schema: formatted }));
      setIsValidJSON(true);
      setError(null);
    } catch (e) {
      setError("Invalid JSON: Please check your input");
      setIsValidJSON(false);
    }
  };

  const minifyJSON = () => {
    try {
      const parsed = JSON.parse(formData.schema);
      const minified = JSON.stringify(parsed);
      setFormData((prev) => ({ ...prev, schema: minified }));
      setIsValidJSON(true);
      setError(null);
    } catch (e) {
      setError("Invalid JSON: Please check your input");
      setIsValidJSON(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!validateJSON(formData.schema)) {
      setError(
        "Invalid JSON: Please format your JSON correctly before submitting"
      );
      setIsValidJSON(false);
      return;
    }

    try {
      const parsedJSON = JSON.parse(formData.schema);
      const formattedData = {
        ...formData,
        schema: parsedJSON,
      };

      await submitSchema(formattedData);
      setIsModalOpen(false);
      setFormData(initialFormData);
      setError(null);
    } catch (error) {
      setError(error instanceof Error ? error.message : "An error occurred");
    }
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Add New Schema">
      <form onSubmit={handleSubmit} className="form-container">
        <div className="form-group">
          <label htmlFor="name" className="form-label">
            Schema Name:
          </label>
          <input
            type="text"
            id="name"
            name="name"
            value={formData.name}
            onChange={handleChange}
            className="form-input"
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="version" className="form-label">
            Version:
          </label>
          <input
            type="text"
            id="version"
            name="version"
            value={formData.version}
            onChange={handleChange}
            className="form-input"
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="schema" className="form-label">
            Schema Definition (JSON):
          </label>
          <div className="json-controls">
            <button
              type="button"
              onClick={formatJSON}
              className="button button-secondary"
            >
              Format JSON
            </button>
            <button
              type="button"
              onClick={minifyJSON}
              className="button button-secondary"
            >
              Minify JSON
            </button>
          </div>
          <textarea
            id="schema"
            name="schema"
            value={formData.schema}
            onChange={handleChange}
            className={`form-textarea ${!isValidJSON ? "invalid-json" : ""}`}
            placeholder="Enter valid JSON here..."
            required
          />
          {!isValidJSON && (
            <div className="json-error">Please enter valid JSON</div>
          )}
          {error && <div className="error-message">{error}</div>}
        </div>

        <div className="button-group">
          <button
            type="button"
            onClick={onClose}
            className="button button-secondary"
          >
            Cancel
          </button>
          <button type="submit" className="button button-primary">
            Add Schema
          </button>
        </div>
      </form>
    </Modal>
  );
};

export default AddNewSchema;
