import { useState } from "react";
import DisplayAllSchemas from "../components/DisplayAllSchemas/DisplayAllSchemas";
import AddNewSchema from "../components/AddNewSchema/AddNewSchema";
import "../styles/components.css";

function App() {
  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <div className="app-container">
      <h1>Ember Schema Registry</h1>
      <button
        onClick={() => setIsModalOpen(true)}
        className="add-schema-button"
      >
        Add New Schema
      </button>
      <AddNewSchema
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
      />
      <DisplayAllSchemas />
    </div>
  );
}

export default App;
