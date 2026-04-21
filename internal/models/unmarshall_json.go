package models

import (
	"encoding/json"
	"errors"
	"syscall/js"
)

// generatePDFWrapper acts as the bridge between JS and Go
func ParseJSON(this js.Value, args []js.Value) (PropertyQuotation, error) {

	if len(args) < 1 {
		return PropertyQuotation{}, errors.New("No JSON Provided")
	}

	// 1. Get the JSON string from JS argument
	jsonInput := args[0].String()

	// 2. Unmarshal into your struct
	var quotation PropertyQuotation
	err := json.Unmarshal([]byte(jsonInput), &quotation)
	if err != nil {
		return PropertyQuotation{}, errors.New("JSON cannot be parsed")
	}

	return quotation, nil
}
