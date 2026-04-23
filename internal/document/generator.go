package document

import (
	"fmt"
	"strings"

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
	err := m.OutputFileAndClose("output.pdf")
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
			m.Col(12, func() { m.Text("Furnishing Checklist:", props.Text{Style: consts.Bold, Size: 9}) })
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
			m.Text(fmt.Sprintf("Maintenance Fee: RM %.2f/psf", quote.LegalAndFees.MaintenanceFeePSF), valueProp)
		})
		m.Col(8, func() {
			m.Text(fmt.Sprintf("Total Estimated: RM %.2f/month", quote.LegalAndFees.MaintenanceFeeTotal), valueProp)
		})
	})

	// 1. INCLUDED SECTION
	// Header
	m.Row(10, func() {
		m.Col(12, func() { m.Text("Included:", labelProp) })
	})

	// Bulleted List
	for _, item := range quote.LegalAndFees.Included {
		m.Row(6, func() {
			m.Col(1, func() { m.Text("-", valueProp) })   // The bullet
			m.Col(11, func() { m.Text(item, valueProp) }) // The text
		})
	}

	// Add a little vertical spacing between the two sections
	m.Row(5, func() {})

	// 2. NOT INCLUDED SECTION
	// Header
	m.Row(10, func() {
		m.Col(12, func() { m.Text("Not Included:", labelProp) })
	})

	// Bulleted List
	for _, item := range quote.LegalAndFees.NotIncluded {
		m.Row(6, func() {
			m.Col(1, func() { m.Text("-", valueProp) })   // The bullet
			m.Col(11, func() { m.Text(item, valueProp) }) // The text
		})
	}

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
