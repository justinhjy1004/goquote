package test

import (
	"testing"
	"time"

	"github.com/justinhjy1004/goquote/internal/document"
	"github.com/justinhjy1004/goquote/internal/models"
)

// Test for multiple options (Original Case)
func TestGenerateQuotationMultiOptionPDF(t *testing.T) {
	quote := models.PropertyQuotation{
		AppointmentDate:   time.Now(),
		QuotationValidity: time.Now().AddDate(0, 1, 0),
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
				DownPayment:       75000.00,
				NettPrice:         725000.00,
				LoanAmount:        675000.00,
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
			Logo:        "https://dummyimage.com/150x50/000/fff&text=AGENCY+LOGO",
			Signature:   "https://dummyimage.com/200x80/ffffff/000000&text=Sarah+Lim+Signature",
		},
	}

	outputPath := "sample_quotation_multi_option.pdf"
	err := document.GeneratePDFDocument(quote, outputPath)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	t.Logf("Generated Multi-Option PDF: %s", outputPath)
}

// Test for a single option (New Case)
func TestGenerateQuotationSingleOptionPDF(t *testing.T) {
	quote := models.PropertyQuotation{
		AppointmentDate:   time.Now(),
		QuotationValidity: time.Now().AddDate(0, 0, 14), // Valid for 14 days
		LeadInfo: models.Lead{
			Name:    "John Doe",
			Contact: "+60 11-222 3333",
		},
		ProjectDetails: models.Project{
			ProjectName: "Urban Suites",
			Developer:   "City Dev Group",
			Tenure:      "Leasehold",
			UnitNo:      "B-10-10",
			Facing:      "Pool View",
			LayoutType:  "Studio",
			AreaSqft:    550,
			SPAPrice:    450000.00,
		},
		Options: []models.Option{
			{
				OptionName:        "Essential Entry Pack",
				Rebate:            45000.00,
				Cashback:          0.00,
				DownPayment:       0.00, // 0 Downpayment promo
				NettPrice:         405000.00,
				LoanAmount:        405000.00,
				InterestRate:      4.10,
				MonthlyInstalment: 1850.20,
				Furnishing: models.Furnishing{
					KitchenCabinet: true,
					Airconds:       1,
					Heater:         true,
				},
			},
		},
		LegalAndFees: models.LegalFees{
			MaintenanceFeePSF:   0.40,
			MaintenanceFeeTotal: 220.00,
			Included:            []string{"SPA Legal Fee"},
			NotIncluded:         []string{"Loan Stamp Duty", "MOT"},
		},
		Agent: models.Agent{
			Name:        "Michael Tan",
			PhoneNumber: "+60 12-999 8888",
			Email:       "michael.tan@proptech.com",
			Logo:        "https://dummyimage.com/150x50/222/eee&text=PROPTECH",
			Signature:   "https://dummyimage.com/200x80/ffffff/000000&text=MT+Signature",
		},
	}

	outputPath := "sample_quotation_single_option.pdf"
	err := document.GeneratePDFDocument(quote, outputPath)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	t.Logf("Generated Single-Option PDF: %s", outputPath)
}
