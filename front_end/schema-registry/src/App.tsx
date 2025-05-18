import { useState } from "react";
import DisplayAllSchemas from "../components/DisplayAllSchemas/DisplayAllSchemas";
import AddNewSchema from "../components/AddNewSchema/AddNewSchema";
import UpdateSchema from "../components/UpdateSchema/UpdateSchema";
import DeleteSchema from "../components/DeleteSchema/DeleteSchema";

import "../styles/components.css";

function App() {
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [isUpdateModalOpen, setIsUpdateModalOpen] = useState(false);
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState(false);

  return (
    <div className="app-container">
      <h1>Ember Schema Registry</h1>
      <div className="button-container">
        <button
          onClick={() => setIsAddModalOpen(true)}
          className="add-schema-button"
        >
          Add New Schema
        </button>
        <AddNewSchema
          isOpen={isAddModalOpen}
          onClose={() => setIsAddModalOpen(false)}
        />
        <button
          onClick={() => setIsUpdateModalOpen(true)}
          className="add-schema-button"
        >
          Update Schema
        </button>
        <UpdateSchema
          isOpen={isUpdateModalOpen}
          onClose={() => setIsUpdateModalOpen(false)}
        />
        <button
          onClick={() => setIsDeleteModalOpen(true)}
          className="add-schema-button"
        >
          Delete Schema
        </button>
        <DeleteSchema
          isOpen={isDeleteModalOpen}
          onClose={() => setIsDeleteModalOpen(false)}
        />
      </div>
      <DisplayAllSchemas />
    </div>
  );
}

export default App;
