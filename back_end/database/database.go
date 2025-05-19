package databaseClient

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseClient struct {
	dataBaseClient *sql.DB
}

type Schema struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

func NewDatabaseClient(connectionString string) (*DatabaseClient, error) {
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection within NewDatabaseClient: %w", err)
	}

	if err := connection.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping connection within NewDatabaseClient: %w", err)
	}

	return &DatabaseClient{dataBaseClient: connection}, nil
}

func (databaseClient *DatabaseClient) GetSchemas() ([]Schema, error) {
	err := databaseClient.createTable()
	if err != nil {
		return nil, fmt.Errorf("failed to create table within getSchemas: %w", err)
	}

	rows, err := databaseClient.dataBaseClient.Query("SELECT * FROM valid_schemas")
	if err != nil {
		return nil, fmt.Errorf("failed to query for schemas within getSchemas: %w", err)
	}

	defer rows.Close()

	schemas := []Schema{}

	for rows.Next() {
		var schema Schema
		if err := rows.Scan(&schema.ID, &schema.Name, &schema.Schema); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		schemas = append(schemas, schema)
	}

	return schemas, nil
}

func (databaseClient *DatabaseClient) InsertSchema(schemaName string, schemaDefinition json.RawMessage) error {
	err := databaseClient.createTable()
	if err != nil {
		return fmt.Errorf("failed to create table within InsertSchema: %w", err)
	}

	query := `
	INSERT INTO valid_schemas (name, schema_definition) VALUES (?, ?)
	`
	_, err = databaseClient.dataBaseClient.Exec(query, schemaName, schemaDefinition)
	if err != nil {
		return fmt.Errorf("failed to insert schema within InsertSchema: %w", err)
	}

	return nil
}

func (databaseClient *DatabaseClient) DeleteSchema(schemaId string) error {
	err := databaseClient.createTable()
	if err != nil {
		return fmt.Errorf("failed to create table within DeleteSchema: %w", err)
	}

	query := `
	DELETE FROM valid_schemas WHERE id = ?
	`
	_, err = databaseClient.dataBaseClient.Exec(query, schemaId)
	if err != nil {
		return fmt.Errorf("failed to delete schema within DeleteSchema: %w", err)
	}

	return nil
}

func (databaseClient *DatabaseClient) UpdateSchema(schemaId string, schemaDefinition json.RawMessage) error {
	err := databaseClient.createTable()
	if err != nil {
		return fmt.Errorf("failed to create table within UpdateSchema: %w", err)
	}

	query := `
	UPDATE valid_schemas SET schema_definition = ? WHERE id = ?
	`
	_, err = databaseClient.dataBaseClient.Exec(query, schemaDefinition, schemaId)
	if err != nil {
		return fmt.Errorf("failed to update schema within UpdateSchema: %w", err)
	}

	return nil
}

func (databaseClient *DatabaseClient) createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS valid_schemas(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		schema_definition JSON NOT NULL
	);
	`

	_, err := databaseClient.dataBaseClient.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
