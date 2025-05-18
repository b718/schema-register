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

type SubmitSchemasRequest struct {
	Schemas []string `json:"schemas"`
}

type ValidateSchemasRequest struct {
	SchemaId string   `json:"schemaId"`
	Schemas  []string `json:"schemas"`
}

type InsertSchemasRequest struct {
	Name   string          `json:"name"`
	Schema json.RawMessage `json:"schema"`
}

type InsertSchemasResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type FetchSchemasResponse struct {
	StatusCode int                     `json:"statusCode"`
	Schemas    []databaseClient.Schema `json:"schemas"`
}

type DeleteSchemaResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type UpdateSchemaRequest struct {
	Schema json.RawMessage `json:"schema"`
}

type UpdateSchemaResponse struct {
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

	var updateRequest struct {
		Schema json.RawMessage `json:"schema"`
	}

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

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
