package main

import (
	"fmt"
	"syscall/js"

	"github.com/justinhjy1004/goquote/internal/document"
	"github.com/justinhjy1004/goquote/internal/models"
)

// generatePDFWrapper acts as the bridge between JS and Go
func generatePDFWrapper(this js.Value, args []js.Value) any {

	quote, err := models.ParseJSON(this, args)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	buffer, err := document.GenerateQuotationPDF(quote)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	uint8Array := js.Global().Get("Uint8Array").New(len(buffer))
	js.CopyBytesToJS(uint8Array, buffer)

	return uint8Array
}

func main() {
	// Prevent the Go program from exiting immediately
	c := make(chan struct{}, 0)

	// Register the function so it's accessible in JS as 'generatePDF'
	js.Global().Set("generatePDF", js.FuncOf(generatePDFWrapper))

	fmt.Println("Go Wasm Quote Module Loaded")
	<-c
}
