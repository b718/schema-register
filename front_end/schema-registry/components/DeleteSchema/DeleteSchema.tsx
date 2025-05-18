import React, { useState } from "react";
import Modal from "../Modal/Modal";
import { deleteSchema } from "./utilities/deleteSchema";

interface FormData {
  id: string;
}

interface DeleteSchemaProps {
  isOpen: boolean;
  onClose: () => void;
}

const initialFormData: FormData = {
  id: "",
};

const DeleteSchema: React.FC<DeleteSchemaProps> = ({ isOpen, onClose }) => {
  const [formData, setFormData] = useState<FormData>(initialFormData);
  const [error, setError] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      await deleteSchema(formData.id);
      onClose();
      setFormData(initialFormData);
      setError(null);
    } catch (error) {
      setError("Failed to delete schema");
    }
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Delete Schema">
      <form onSubmit={handleSubmit} className="form-container">
        <div className="form-group">
          <label htmlFor="id" className="form-label">
            Schema ID:
          </label>
          <input
            type="text"
            id="id"
            name="id"
            value={formData.id}
            onChange={handleChange}
            className="form-input"
            required
          />
        </div>

        {error && <div className="error-message">{error}</div>}

        <div className="button-group">
          <button
            type="button"
            onClick={onClose}
            className="button button-secondary"
          >
            Cancel
          </button>
          <button type="submit" className="button button-primary">
            Delete Schema
          </button>
        </div>
      </form>
    </Modal>
  );
};

export default DeleteSchema;
