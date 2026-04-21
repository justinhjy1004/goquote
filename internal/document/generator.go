package document

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/justinhjy1004/goquote/internal/models"
)

// GenerateQuotationPDF takes a PropertyQuotation and outputs a PDF file
func GenerateQuotationPDF(quote models.PropertyQuotation) ([]byte, error) {
	// Initialize Maroto (Portrait, A4)
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	// Styles
	headerProp := props.Text{Style: consts.Bold, Size: 14, Align: consts.Center}
	sectionProp := props.Text{Style: consts.Bold, Size: 11, Color: getDarkGray()}
	labelProp := props.Text{Style: consts.Bold, Size: 9}
	valueProp := props.Text{Style: consts.Normal, Size: 9}

	// --- HEADER (Logo & Title) ---
	m.Row(20, func() {
		m.Col(3, func() {
			// Fetch logo from URL and embed as Base64
			if logoB64, err := urlToBase64(quote.Agent.Logo); err == nil && logoB64 != "" {
				m.Base64Image(logoB64, consts.Png, props.Rect{Center: true, Percent: 100})
			} else {
				m.Text("LOGO", headerProp)
			}
		})
		m.ColSpace(1)
		m.Col(8, func() {
			m.Text("PROPERTY QUOTATION", props.Text{Top: 5, Style: consts.Bold, Size: 18, Align: consts.Right})
			m.Text(fmt.Sprintf("Date: %s | Valid until: %s", quote.AppointmentDate.Format("02 Jan 2006"), quote.QuotationValidity.Format("02 Jan 2006")), props.Text{Top: 12, Style: consts.Normal, Size: 9, Align: consts.Right})
		})
	})
	m.Line(2)

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Client Name:", labelProp) })
		m.Col(3, func() { m.Text(quote.LeadInfo.Name, valueProp) })
		m.Col(3, func() { m.Text("Contact:", labelProp) })
		m.Col(3, func() { m.Text(quote.LeadInfo.Contact, valueProp) })
	})

	// --- CLIENT & PROJECT DETAILS ---
	buildSectionTitle(m, "PROPERTY DETAILS", sectionProp)

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Project Name:", labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.ProjectName, valueProp) })
		m.Col(3, func() { m.Text("Developer:", labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.Developer, valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Unit No:", labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.UnitNo, valueProp) })
		m.Col(3, func() { m.Text("Tenure:", labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.Tenure, valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Layout & Area:", labelProp) })
		m.Col(3, func() {
			m.Text(fmt.Sprintf("%s (%d sqft)", quote.ProjectDetails.LayoutType, quote.ProjectDetails.AreaSqft), valueProp)
		})
		m.Col(3, func() { m.Text("SPA Price:", labelProp) })
		m.Col(3, func() { m.Text(fmt.Sprintf("RM %.2f", quote.ProjectDetails.SPAPrice), valueProp) })
	})

	// --- OPTIONS ---
	for i, opt := range quote.Options {
		m.Row(5, func() {}) // spacer
		buildSectionTitle(m, fmt.Sprintf("OPTION %d: %s", i+1, strings.ToUpper(opt.OptionName)), sectionProp)

		m.Row(5, func() {
			m.Col(3, func() { m.Text("Rebate:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %.2f", opt.Rebate), valueProp) })
			m.Col(3, func() { m.Text("Cashback:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %.2f", opt.Cashback), valueProp) })
		})

		// Add custom discounts
		for _, disc := range opt.Discounts {
			m.Row(5, func() {
				m.Col(3, func() { m.Text(disc.Type+":", labelProp) })
				m.Col(9, func() { m.Text(fmt.Sprintf("RM %.2f", disc.Amount), valueProp) })
			})
		}

		m.Row(5, func() {
			m.Col(3, func() { m.Text("Nett Price:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %.2f", opt.NettPrice), props.Text{Style: consts.Bold, Size: 9}) })
			m.Col(3, func() { m.Text("Down Payment:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %.2f", opt.DownPayment), valueProp) })
		})

		m.Row(5, func() {
			m.Col(3, func() { m.Text("Loan Amount:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %.2f", opt.LoanAmount), valueProp) })
			m.Col(3, func() { m.Text("Interest Rate:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("%.2f%%", opt.InterestRate), valueProp) })
		})

		m.Row(5, func() {
			m.Col(3, func() { m.Text("Est. Monthly Instalment:", props.Text{Style: consts.BoldItalic, Size: 9}) })
			m.Col(9, func() {
				m.Text(fmt.Sprintf("RM %.2f", opt.MonthlyInstalment), props.Text{Style: consts.Bold, Size: 10})
			})
		})

		// --- FURNISHING CHECKLIST (Grid Layout) ---
		m.Row(6, func() {
			m.Col(12, func() { m.Text("Furnishing Checklist:", props.Text{Style: consts.BoldItalic, Size: 9}) })
		})

		// Row 1: Booleans
		m.Row(5, func() {
			m.Col(4, func() { m.Text(fmt.Sprintf("%s Kitchen Cabinet", checkBox(opt.Furnishing.KitchenCabinet)), valueProp) })
			m.Col(4, func() { m.Text(fmt.Sprintf("%s Hood & Hob", checkBox(opt.Furnishing.HoodAndHob)), valueProp) })
			m.Col(4, func() { m.Text(fmt.Sprintf("%s Fridge", checkBox(opt.Furnishing.Fridge)), valueProp) })
		})

		// Row 2: Booleans
		m.Row(5, func() {
			m.Col(4, func() { m.Text(fmt.Sprintf("%s Toilet Fittings", checkBox(opt.Furnishing.Toilet)), valueProp) })
			m.Col(4, func() { m.Text(fmt.Sprintf("%s Water Heater", checkBox(opt.Furnishing.Heater)), valueProp) })
			m.Col(4, func() { m.Text(fmt.Sprintf("%s Shower Screen", checkBox(opt.Furnishing.ShowerScreen)), valueProp) })
		})

		// Row 3: Numeric Quantities
		m.Row(5, func() {
			m.Col(4, func() { m.Text(fmt.Sprintf("[%d]  Washing Machine", opt.Furnishing.WashingMachine), valueProp) })
			m.Col(4, func() { m.Text(fmt.Sprintf("[%d]  Airconds", opt.Furnishing.Airconds), valueProp) })
			m.Col(4, func() { m.Text(fmt.Sprintf("[%d]  Built-in Wardrobe", opt.Furnishing.WardrobeQty), valueProp) })
		})

		// Row 4: Remaining Numeric
		m.Row(5, func() {
			m.Col(4, func() { m.Text(fmt.Sprintf("[%d]  Queen-size Bed", opt.Furnishing.BedSetQty), valueProp) })
		})

		// Row 5+: Iterate over the free-text "Additional" items dynamically
		if len(opt.Furnishing.Additional) > 0 {
			m.Row(3, func() {}) // spacer

			// Chunk the additional items into groups of 3 for the columns
			for i := 0; i < len(opt.Furnishing.Additional); i += 3 {
				m.Row(5, func() {
					// Column 1
					m.Col(4, func() { m.Text("[X] "+opt.Furnishing.Additional[i], valueProp) })

					// Column 2
					if i+1 < len(opt.Furnishing.Additional) {
						m.Col(4, func() { m.Text("[X] "+opt.Furnishing.Additional[i+1], valueProp) })
					}

					// Column 3
					if i+2 < len(opt.Furnishing.Additional) {
						m.Col(4, func() { m.Text("[X] "+opt.Furnishing.Additional[i+2], valueProp) })
					}
				})
			}
		}

		m.Row(3, func() {}) // Add a little breathing room before the line

		m.Line(1)
	}

	// --- LEGAL & FEES ---
	m.Row(5, func() {}) // spacer
	buildSectionTitle(m, "LEGAL & MAINTENANCE FEES", sectionProp)

	m.Row(5, func() {
		m.Col(4, func() {
			m.Text(fmt.Sprintf("Maintenance Fee: RM %.2f/psf", quote.LegalAndFees.MaintenanceFeePSF), valueProp)
		})
		m.Col(8, func() {
			m.Text(fmt.Sprintf("Total Estimated: RM %.2f/month", quote.LegalAndFees.MaintenanceFeeTotal), valueProp)
		})
	})

	m.Row(10, func() {
		m.Col(2, func() { m.Text("Included:", labelProp) })
		m.Col(4, func() { m.Text(strings.Join(quote.LegalAndFees.Included, ", "), valueProp) })
		m.Col(2, func() { m.Text("Not Included:", labelProp) })
		m.Col(4, func() { m.Text(strings.Join(quote.LegalAndFees.NotIncluded, ", "), valueProp) })
	})

	// --- AGENT SIGNATURE (FOOTER) ---
	m.Row(30, func() {
		m.Col(8, func() {}) // empty space on left
		m.Col(4, func() {
			if sigB64, err := urlToBase64(quote.Agent.Signature); err == nil && sigB64 != "" {
				// Extension doesn't matter too much for maroto's base64 reader as long as it's valid image data
				m.Base64Image(sigB64, consts.Png, props.Rect{Center: true, Percent: 80})
			}
			m.Text("___________________________", props.Text{Top: 20, Align: consts.Center})
			m.Text(quote.Agent.Name, props.Text{Top: 25, Align: consts.Center, Style: consts.Bold, Size: 9})
			m.Text(quote.Agent.PhoneNumber, props.Text{Top: 30, Align: consts.Center, Size: 9})
		})
	})

	buffer, err := m.Output()
	if err != nil {
		return nil, fmt.Errorf("could not generate PDF buffer: %v", err)
	}

	// 3. Return the bytes from the buffer
	return buffer.Bytes(), nil
}

// --- HELPER FUNCTIONS ---

func buildSectionTitle(m pdf.Maroto, title string, prop props.Text) {
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text(title, prop)
		})
	})
	m.Row(2, func() {}) // spacer
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

	// Determine mime type if necessary, but returning pure base64
	// (Maroto usually requires just the base64 characters, not the "data:image/png;base64," prefix)
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
