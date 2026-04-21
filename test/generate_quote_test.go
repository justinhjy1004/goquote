package test

import (
	"testing"
	"time"

	"github.com/justinhjy1004/goquote/internal/document"
	"github.com/justinhjy1004/goquote/internal/models"
)

func TestGenerateQuotationPDF(t *testing.T) {
	// 1. Create realistic dummy data
	quote := models.PropertyQuotation{
		AppointmentDate:   time.Now(),
		QuotationValidity: time.Now().AddDate(0, 1, 0), // Valid for 1 month
		LeadInfo: models.Lead{
			Name:    "Ahmad Albab",
			Contact: "+60 12-345 6789",
		},
		ProjectDetails: models.Project{
			ProjectName: "Skyline Residency",
			Developer:   "MegaBina Sdn Bhd",
			Tenure:      "Freehold",
			UnitNo:      "A-15-03",
			Facing:      "KLCC View",
			LayoutType:  "Type B - 3B2R2CP",
			AreaSqft:    1050,
			SPAPrice:    750000.00,
		},
		Options: []models.Option{
			{
				OptionName:        "Standard Rebate",
				Rebate:            20000.00,
				Cashback:          5000.00,
				DownPayment:       75000.00, // 10%
				NettPrice:         725000.00,
				LoanAmount:        675000.00, // 90%
				InterestRate:      4.25,
				MonthlyInstalment: 3320.50,
				Discounts: []models.Discount{
					{Type: "Early Bird Discount", Amount: 5000.00},
				},
				Furnishing: models.Furnishing{
					KitchenCabinet: true,
					HoodAndHob:     true,
					Airconds:       2,
					Heater:         true,
				},
			},
			{
				OptionName:        "Fully Furnished Package",
				Rebate:            10000.00,
				Cashback:          0.00,
				DownPayment:       75000.00,
				NettPrice:         740000.00,
				LoanAmount:        675000.00,
				InterestRate:      4.25,
				MonthlyInstalment: 3320.50,
				Discounts: []models.Discount{
					{Type: "Loyalty Discount", Amount: 2000.00},
				},
				Furnishing: models.Furnishing{
					KitchenCabinet: true,
					HoodAndHob:     true,
					Fridge:         true,
					WashingMachine: 1,
					Airconds:       4,
					Toilet:         true,
					Heater:         true,
					ShowerScreen:   true,
					WardrobeQty:    3,
					BedSetQty:      3,
					Additional:     []string{"Smart Home System", "Digital Lockset", "Curtains"},
				},
			},
		},
		LegalAndFees: models.LegalFees{
			MaintenanceFeePSF:   0.35,
			MaintenanceFeeTotal: 367.50,
			Included:            []string{"SPA Legal Fee", "Loan Legal Fee"},
			NotIncluded:         []string{"MOT", "Valuation Fee"},
		},
		Agent: models.Agent{
			Name:        "Sarah Lim",
			PhoneNumber: "+60 19-876 5432",
			Email:       "sarah.lim@agency.com",
			// Using dummy image URLs that return safe placeholder images so the test doesn't crash
			Logo:      "https://dummyimage.com/150x50/000/fff&text=AGENCY+LOGO",
			Signature: "https://dummyimage.com/200x80/ffffff/000000&text=Sarah+Lim+Signature",
		},
	}

	// 2. Define output path
	outputPath := "sample_quotation_test.pdf"

	// 3. Call the generation function
	quote, err := document.GenerateQuotationPDF(quote)

	// 4. Check for errors
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	t.Logf("Successfully generated test PDF at: %s", outputPath)
}
