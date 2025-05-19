package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	databaseClient "registerschemas/database"
)

type FetchSchemasResponse struct {
	StatusCode int                     `json:"statusCode"`
	Schemas    []databaseClient.Schema `json:"schemas"`
}

func FetchSchemas(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	databaseClient, err := databaseClient.NewDatabaseClient(sqlConnectionLink)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to database: %v", err), http.StatusInternalServerError)
		return
	}

	schemas, err := databaseClient.GetSchemas()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get schemas: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(FetchSchemasResponse{
		StatusCode: http.StatusOK,
		Schemas:    schemas,
	})
}
