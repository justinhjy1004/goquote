package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/invopop/jsonschema"
	"github.com/justinhjy1004/goquote/internal/models"
)

func main() {
	// Reflect the struct to generate the OpenAPI-like JSON Schema
	schema := jsonschema.Reflect(&models.PropertyQuotation{})
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println("Error generating schema:", err)
		os.Exit(1)
	}
	err = os.WriteFile("schema.json", schemaJSON, 0644)
	if err != nil {
		fmt.Println("Error writing schema file:", err)
		os.Exit(1)
	}
}
