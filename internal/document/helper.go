package document

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func buildSectionTitle(m pdf.Maroto, title string, prop props.Text) {
	m.Row(1, func() {}) // spacer

	m.Row(8, func() {
		m.Col(12, func() {
			m.Text(title, prop)
		})
	})

}

func getDarkGray() color.Color {
	return color.Color{Red: 50, Green: 50, Blue: 50}
}

// urlToBase64 fetches an image URL and converts it to a base64 string
func urlToBase64(imageURL string) (string, error) {

	if imageURL == "" {
		return "", nil
	}

	resp, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	// Determine mime type if necessary, but returning pure base64 for Maroto rendering
	b64 := base64.StdEncoding.EncodeToString(bytes)
	return b64, nil
}

// checkBox returns a marked or empty box for boolean values
func checkBox(b bool) string {
	if b {
		return "[X]"
	}
	return "[  ]"
}

// formatIntegerPart is a helper to add commas to the integer part of a number string.
func formatIntegerPart(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var result strings.Builder
	result.Grow(n + (n-1)/3) // Pre-allocate memory for efficiency

	// Calculate the position of the first comma
	firstComma := n % 3
	if firstComma == 0 {
		firstComma = 3
	}

	result.WriteString(s[:firstComma])

	for i := firstComma; i < n; i += 3 {
		result.WriteByte(',')
		result.WriteString(s[i : i+3])
	}
	return result.String()
}

// formatCurrency formats a float64 as a currency string with thousands commas and two decimal places.
func formatCurrency(amount float64) string {
	// Format to two decimal places
	s := fmt.Sprintf("%.2f", amount)

	// Split into integer and decimal parts
	parts := strings.Split(s, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	// Handle negative sign
	sign := ""
	if strings.HasPrefix(integerPart, "-") {
		sign = "-"
		integerPart = integerPart[1:]
	}

	return sign + formatIntegerPart(integerPart) + decimalPart
}

// formatInteger formats an int as a string with thousands commas.
func formatInteger(num int) string {
	s := fmt.Sprintf("%d", num)

	sign := ""
	if strings.HasPrefix(s, "-") {
		sign = "-"
		s = s[1:]
	}

	return sign + formatIntegerPart(s)
}
