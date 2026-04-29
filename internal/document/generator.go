package document

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/justinhjy1004/goquote/internal/models"
)

func GeneratePDFByteString(quote models.PropertyQuotation) ([]byte, error) {

	m := generatePDFMaroto(quote)

	buffer, err := m.Output()

	if err != nil {
		return nil, fmt.Errorf("could not generate PDF buffer: %v", err)
	}

	// 3. Return the bytes from the buffer
	return buffer.Bytes(), nil

}

func GeneratePDFDocument(quote models.PropertyQuotation, output string) error {

	m := generatePDFMaroto(quote)

	// This replaces the need for os.Create or os.Write
	err := m.OutputFileAndClose(output)
	if err != nil {
		return err
	}

	return nil

}

type standardTextProperty struct {
	headerProp  props.Text
	sectionProp props.Text
	labelProp   props.Text
	valueProp   props.Text
}

// GenerateQuotationPDF takes a PropertyQuotation and outputs a PDF file
func generatePDFMaroto(quote models.PropertyQuotation) pdf.Maroto {

	// Page Outline and Dimension
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	// Text Styles
	headerProp := props.Text{Style: consts.Bold, Size: 14, Align: consts.Center}
	sectionProp := props.Text{Style: consts.Bold, Size: 11, Color: getDarkGray()}
	labelProp := props.Text{Style: consts.Bold, Size: 9}
	valueProp := props.Text{Style: consts.Normal, Size: 9}

	textStyles := standardTextProperty{
		headerProp:  headerProp,
		sectionProp: sectionProp,
		labelProp:   labelProp,
		valueProp:   valueProp,
	}

	// Generate Header of Quote
	generateHeader(m, quote, textStyles)

	m.Line(2)

	// Client Information
	generateClientDetails(m, quote, textStyles)

	// Property Information
	generatePropertyDetails(m, quote, textStyles)

	m.Line(2)

	// --- OPTIONS ---
	for i, opt := range quote.Options {
		m.Row(5, func() {}) // spacer

		if len(quote.Options) > 1 {
			buildSectionTitle(m, fmt.Sprintf("OPTION %d: %s", i+1, strings.ToUpper(opt.OptionName)), sectionProp)
		}

		m.Row(5, func() {
			m.Col(3, func() { m.Text("Rebate (%):", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("%.2f%%", opt.RebatePercentage), valueProp) })
			m.Col(3, func() { m.Text("Rebate Amount:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %s", formatCurrency(opt.Rebate)), valueProp) })
		})

		m.Row(5, func() {
			m.Col(3, func() { m.Text("Cashback:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %s", formatCurrency(opt.Cashback)), valueProp) })

			m.Col(3, func() { m.Text("Cashback Type:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("%s", opt.CashbackType), valueProp) })
		})

		if len(opt.Discounts) > 0 {
			m.Row(5, func() {
				m.Col(12, func() { m.Text("Additional Discounts:", props.Text{Style: consts.Bold, Size: 9}) })
			})

			// Add custom discounts
			for _, disc := range opt.Discounts {
				m.Row(5, func() {
					m.Col(3, func() { m.Text("- "+disc.Type, labelProp) })
					m.Col(9, func() { m.Text(fmt.Sprintf("%.2f%%", disc.Percentage), valueProp) })
				})
			}
		}

		m.Row(5, func() {
			m.Col(3, func() {
				m.Text("Nett Price:", props.Text{Style: consts.Bold, Size: 9.5, Color: color.Color{Red: 204, Green: 85, Blue: 0}})
			})
			m.Col(3, func() {
				m.Text(fmt.Sprintf("RM %s", formatCurrency(opt.NettPrice)), props.Text{Style: consts.Bold, Size: 9.5, Color: color.Color{Red: 204, Green: 85, Blue: 0}})
			})
		})

		m.Row(5, func() {
			m.Col(3, func() { m.Text("Down Payment:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %s", formatCurrency(opt.DownPayment)), valueProp) })
			m.Col(3, func() { m.Text("Loan Amount:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("RM %s", formatCurrency(opt.LoanAmount)), valueProp) })
		})

		m.Row(5, func() {
			m.Col(3, func() { m.Text("Interest Rate (%):", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("%.2f%%", opt.InterestRate), valueProp) })
			m.Col(3, func() { m.Text("Loan Tenure:", labelProp) })
			m.Col(3, func() { m.Text(fmt.Sprintf("%d", opt.LoanTenureYear), valueProp) })
		})

		m.Row(5, func() {
			m.Col(3, func() {
				m.Text("Est. Monthly Instalment:", props.Text{Style: consts.Bold, Size: 9.5, Color: color.Color{Red: 204, Green: 85, Blue: 0}})
			})
			m.Col(9, func() {
				m.Text(fmt.Sprintf("RM %s", formatCurrency(opt.MonthlyInstalment)), props.Text{Style: consts.Bold, Size: 9.5, Color: color.Color{Red: 204, Green: 85, Blue: 0}})
			})
		})

		type BooleanItems struct {
			Label string
			Show  bool
		}

		type QuantItems struct {
			Label string
			Count int
		}

		// Map your data to a slice of items
		boolItems := []BooleanItems{
			{Label: "Kitchen Cabinet", Show: opt.Furnishing.KitchenCabinet},
			{Label: "Hood & Hob", Show: opt.Furnishing.HoodAndHob},
			{Label: "Fridge", Show: opt.Furnishing.Fridge},
			{Label: "Toilet Fittings", Show: opt.Furnishing.Toilet},
			{Label: "Water Heater", Show: opt.Furnishing.Heater},
			{Label: "Shower Screen", Show: opt.Furnishing.ShowerScreen},
			{Label: "Bathroom Accessories", Show: opt.Furnishing.BathroomAccessories},
			{Label: "Light Fixtures", Show: opt.Furnishing.LightFixtures},
		}

		quantItems := []QuantItems{
			{Label: "Washing Machine", Count: opt.Furnishing.WashingMachine},
			{Label: "Aircond", Count: opt.Furnishing.Airconds},
			{Label: "Wardrobe", Count: opt.Furnishing.WardrobeQty},
			{Label: "Bed Set", Count: opt.Furnishing.BedSetQty},
		}

		// Filter out the items that are false or 0
		var activeBoolItems []BooleanItems
		for _, i := range boolItems {
			if i.Show {
				activeBoolItems = append(activeBoolItems, i)
			}
		}

		var activeQuantItems []QuantItems
		for _, i := range quantItems {
			if i.Count > 0 {
				activeQuantItems = append(activeQuantItems, i)
			}
		}

		if len(activeBoolItems) != 0 && len(activeQuantItems) != 0 {
			// --- FURNISHING CHECKLIST (Grid Layout) ---
			m.Row(6, func() {
				m.Col(12, func() { m.Text("Furnishing Checklist:", props.Text{Style: consts.Bold, Size: 9}) })
			})
		}

		for i := 0; i < len(activeBoolItems); i += 3 {
			m.Row(5, func() {
				// Render up to 3 columns per row
				for j := 0; j < 3 && (i+j) < len(activeBoolItems); j++ {
					currentItem := activeBoolItems[i+j]
					m.Col(4, func() {
						// Since we only show 'true' items, we can hardcode the checkmark
						// or keep the checkBox() helper if you prefer the visual style.
						m.Text(fmt.Sprintf("%s %s", checkBox(true), currentItem.Label), valueProp)
					})
				}
			})
		}

		for i := 0; i < len(activeQuantItems); i += 3 {
			m.Row(5, func() {
				// Render up to 3 columns per row
				for j := 0; j < 3 && (i+j) < len(activeQuantItems); j++ {
					currentItem := activeQuantItems[i+j]
					m.Col(4, func() {
						// Since we only show 'true' items, we can hardcode the checkmark
						// or keep the checkBox() helper if you prefer the visual style.
						m.Text(fmt.Sprintf("[%d] %s", currentItem.Count, currentItem.Label), valueProp)
					})
				}
			})
		}

		// Row 5+: Iterate over the free-text "Additional" items dynamically
		if len(opt.Furnishing.Additional) > 0 {
			m.Row(5, func() { m.Text("Additional Furnishing:", props.Text{Style: consts.Bold, Size: 9}) }) // spacer

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
			m.Text("Maintenance Fee (per square feet):", labelProp)
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("RM %s/psf", formatCurrency(quote.LegalAndFees.MaintenanceFeePSF)), valueProp)
		})
	})

	m.Row(5, func() {
		m.Col(4, func() {
			m.Col(4, func() {
				m.Text("Total Maintenance Fee (per month):", labelProp)
			})
			m.Text(fmt.Sprintf("RM %s/month", formatCurrency(quote.LegalAndFees.MaintenanceFeeTotal)), valueProp)
		})
	})

	m.Row(5, func() {
		m.Col(4, func() {
			m.Text("MOT Duty Stamp:", labelProp)
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("RM %s", formatCurrency(quote.LegalAndFees.MOT)), valueProp)
		})
	})

	m.Row(5, func() {
		m.Col(4, func() {
			m.Text("SPA Legal Fee: ", labelProp)
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("%s", formatKey(quote.LegalAndFees.SPALegalFree)), valueProp)
		})
	})

	m.Row(5, func() {
		m.Col(4, func() {
			m.Text("SPA Disbursement Fee: ", labelProp)
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("%s", formatKey(quote.LegalAndFees.SPADisbursementFree)), valueProp)
		})
	})

	m.Row(5, func() {
		m.Col(4, func() {
			m.Text("Loan Agreement Fee: ", labelProp)
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("%s", formatKey(quote.LegalAndFees.LoanAgreementFree)), valueProp)
		})
	})

	m.Row(5, func() {
		m.Col(4, func() {
			m.Text("Loan Disbursement Fee: ", labelProp)
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("%s", formatKey(quote.LegalAndFees.LoanDisbursementFree)), valueProp)
		})
	})

	m.Row(5, func() {
		m.Col(4, func() {
			m.Text("Loan Stamp Duty Fee: ", labelProp)
		})
		m.Col(4, func() {
			m.Text(fmt.Sprintf("%s", formatKey(quote.LegalAndFees.LoanStampDutyFree)), valueProp)
		})
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

	return m
}
