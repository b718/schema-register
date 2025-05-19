package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	databaseClient "registerschemas/database"
)

type InsertSchemasRequest struct {
	Name   string          `json:"name"`
	Schema json.RawMessage `json:"schema"`
}

type InsertSchemasResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func InsertSchemas(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//validate the json
	var insertSchemasRequest InsertSchemasRequest
	err := json.NewDecoder(r.Body).Decode(&insertSchemasRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body in InsertSchemas: %v", err), http.StatusBadRequest)
		return
	}

	if !json.Valid([]byte(insertSchemasRequest.Schema)) {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//save the schemas to the database
	dbClient, err := databaseClient.NewDatabaseClient(sqlConnectionLink)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to database in InsertSchemas: %v", err), http.StatusInternalServerError)
		return
	}

	schemaDefinition, err := json.Marshal(insertSchemasRequest.Schema)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal schema in InsertSchemas: %v", err), http.StatusInternalServerError)
		return
	}

	err = dbClient.InsertSchema(insertSchemasRequest.Name, schemaDefinition)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert schema in InsertSchemas: %v", err), http.StatusInternalServerError)
		return
	}

	// return the result
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(InsertSchemasResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Schemas inserted successfully: %v", insertSchemasRequest.Name),
	})
}
