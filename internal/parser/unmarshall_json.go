package parser

import (
	"encoding/json"
	"errors"
	"syscall/js"

	"github.com/justinhjy1004/goquote/internal/models"
)

// generatePDFWrapper acts as the bridge between JS and Go
func ParseJSON(this js.Value, args []js.Value) (models.PropertyQuotation, error) {

	if len(args) < 1 {
		return models.PropertyQuotation{}, errors.New("No JSON Provided")
	}

	// 1. Get the JSON string from JS argument
	jsonInput := args[0].String()

	// 2. Unmarshal into your struct
	var quotation models.PropertyQuotation
	err := json.Unmarshal([]byte(jsonInput), &quotation)
	if err != nil {
		return models.PropertyQuotation{}, errors.New("JSON cannot be parsed")
	}

	return quotation, nil
}
