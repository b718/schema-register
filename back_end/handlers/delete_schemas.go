package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	databaseClient "registerschemas/database"

	_ "github.com/joho/godotenv/autoload"
)

var (
	sqlConnectionLink = "root:" + os.Getenv("SQL_CONNECTION_PASSWORD") + "@tcp(127.0.0.1:3306)/registered_schemas"
)

type DeleteSchemaResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func DeleteSchema(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract schemaId from URL path
	schemaId := r.URL.Path[len("/delete-schema/"):]
	if schemaId == "" {
		http.Error(w, "Schema ID is required", http.StatusBadRequest)
		return
	}

	databaseClient, err := databaseClient.NewDatabaseClient(sqlConnectionLink)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to database: %v", err), http.StatusInternalServerError)
		return
	}

	err = databaseClient.DeleteSchema(schemaId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete schema: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(DeleteSchemaResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Schema %s deleted successfully", schemaId),
	})
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
