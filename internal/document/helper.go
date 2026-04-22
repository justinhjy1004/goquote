package document

import (
	"encoding/base64"
	"io"
	"net/http"

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
