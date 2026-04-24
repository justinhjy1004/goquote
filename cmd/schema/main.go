package main

import (
	"fmt"
	"os"

	"github.com/hypersequent/zen"
	"github.com/justinhjy1004/goquote/internal/models"
)

func main() {
	// 1. Generate the Zod schema string
	// Note: zen.StructToZodSchema returns a string, not a struct/map
	schema := zen.StructToZodSchema(models.PropertyQuotation{})

	// 2. Format the output string for TypeScript
	// We add 'export const' so you can import it in your TS files
	content := fmt.Sprintf("import { z } from 'zod';\n\n%s;", schema)

	// 3. Write directly to a .ts file
	err := os.WriteFile("schema.ts", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing schema file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated property_quotation.gen.ts")
}
