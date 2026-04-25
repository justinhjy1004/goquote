package document

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"

	"github.com/justinhjy1004/goquote/internal/models"
)

func generateHeader(m pdf.Maroto, quote models.PropertyQuotation, styles standardTextProperty) {

	m.Row(20, func() {
		m.Col(3, func() {
			// Fetch logo from URL and embed as Base64
			if logoB64, err := urlToBase64(quote.Agent.Logo); err == nil && logoB64 != "" {
				m.Base64Image(logoB64, consts.Png, props.Rect{Center: true, Percent: 100})
			} else {
				m.Text(quote.Agent.Name, styles.headerProp)
			}
		})

		m.ColSpace(1)

		m.Col(8, func() {
			m.Text("OFFICIAL QUOTATION", props.Text{Top: 5, Style: consts.Bold, Size: 16, Align: consts.Right})

			m.Text(fmt.Sprintf("Date: %s | Valid until: %s", quote.AppointmentDate.Format("02 Jan 2006"), quote.QuotationValidity.Format("02 Jan 2006")), props.Text{Top: 12, Style: consts.Normal, Size: 9, Align: consts.Right})
		})
	})

	m.Row(1, func() {})
}

func generateClientDetails(m pdf.Maroto, quote models.PropertyQuotation, styles standardTextProperty) {

	buildSectionTitle(m, "CLIENT DETAILS", styles.sectionProp)

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Client Name:", styles.labelProp) })
		m.Col(3, func() { m.Text(quote.LeadInfo.Name, styles.valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Client Contact:", styles.labelProp) })
		m.Col(3, func() { m.Text(quote.LeadInfo.Contact, styles.valueProp) })
	})

	m.Row(1, func() {})

}

func generatePropertyDetails(m pdf.Maroto, quote models.PropertyQuotation, styles standardTextProperty) {

	buildSectionTitle(m, "PROPERTY DETAILS", styles.sectionProp)

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Project Name:", styles.labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.ProjectName, styles.valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Developer:", styles.labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.Developer, styles.valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Unit No:", styles.labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.UnitNo, styles.valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Layout Type:", styles.labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.LayoutType, styles.valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Tenure:", styles.labelProp) })
		m.Col(3, func() { m.Text(quote.ProjectDetails.Tenure, styles.valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("SPA Price:", styles.labelProp) })
		m.Col(3, func() { m.Text(fmt.Sprintf("RM %.2f", quote.ProjectDetails.SPAPrice), styles.valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Facing:", styles.labelProp) })
		m.Col(3, func() { m.Text(fmt.Sprintf("RM %s", formatCurrency(quote.ProjectDetails.SPAPrice)), styles.valueProp) })
	})

	m.Row(5, func() {
		m.Col(3, func() { m.Text("Area Sqft:", styles.labelProp) })
		m.Col(3, func() { m.Text(fmt.Sprintf("%s sqft", formatInteger(quote.ProjectDetails.AreaSqft)), styles.valueProp) })
	})

	m.Row(1, func() {})
}
