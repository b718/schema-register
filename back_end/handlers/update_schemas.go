package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	databaseClient "registerschemas/database"
)

type UpdateSchemaRequest struct {
	Schema json.RawMessage `json:"schema"`
}

type UpdateSchemaResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func UpdateSchema(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract schemaId from URL path
	schemaId := r.URL.Path[len("/update-schema/"):]
	if schemaId == "" {
		http.Error(w, "Schema ID is required", http.StatusBadRequest)
		return
	}

	var updateRequest UpdateSchemaRequest

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	databaseClient, err := databaseClient.NewDatabaseClient(sqlConnectionLink)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to database: %v", err), http.StatusInternalServerError)
		return
	}

	err = databaseClient.UpdateSchema(schemaId, updateRequest.Schema)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update schema: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UpdateSchemaResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Schema %s updated successfully", schemaId),
	})
}
