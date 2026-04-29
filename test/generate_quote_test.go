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
			Name:        "Ahmad Albab",
			Contact:     "+60 12-345 6789",
			Citizenship: "Malaysian", // Added citizenship
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
			CarParkLot:  "2 (Side by Side)", // Added car park lot
		},
		Options: []models.Option{
			{
				OptionName:        "Standard Rebate",
				Rebate:            20000.00,
				RebatePercentage:  2.66, // Added rebate percentage
				Cashback:          5000.00,
				CashbackType:      "Cash Out", // Added cashback type
				DownPayment:       75000.00,
				NettPrice:         725000.00,
				LoanAmount:        675000.00,
				InterestRate:      4.25,
				MonthlyInstalment: 3320.50,
				LoanTenureYear:    35, // Added loan tenure year
				Discounts: []models.Discount{
					{Type: "Early Bird Discount", Percentage: 5.00},
				},
				Furnishing: models.Furnishing{
					KitchenCabinet:      true,
					HoodAndHob:          true,
					Airconds:            2,
					Heater:              true,
					BathroomAccessories: true,  // Added bathroom accessories
					LightFixtures:       false, // Added light fixtures
				},
			},
			{
				OptionName:        "Fully Furnished Package",
				Rebate:            10000.00,
				RebatePercentage:  1.33,
				Cashback:          0.00,
				CashbackType:      "N/A",
				DownPayment:       75000.00,
				NettPrice:         740000.00,
				LoanAmount:        675000.00,
				InterestRate:      4.25,
				MonthlyInstalment: 3320.50,
				LoanTenureYear:    35,
				Discounts: []models.Discount{
					{Type: "Loyalty Discount", Percentage: 2.00},
				},
				Furnishing: models.Furnishing{
					KitchenCabinet:      true,
					HoodAndHob:          true,
					Fridge:              true,
					WashingMachine:      1,
					Airconds:            4,
					Toilet:              true,
					Heater:              true,
					ShowerScreen:        true,
					WardrobeQty:         3,
					BedSetQty:           3,
					Additional:          []string{"Smart Home System", "Digital Lockset", "Curtains"},
					BathroomAccessories: true,
					LightFixtures:       true,
				},
			},
		},
		LegalAndFees: models.LegalFees{
			MaintenanceFeePSF:    0.35,
			MaintenanceFeeTotal:  367.50,
			Included:             []string{"SPA Legal Fee", "Loan Legal Fee"},
			NotIncluded:          []string{"MOT", "Valuation Fee"},
			MOT:                  15000.00, // Added MOT
			SPALegalFree:         "Yes",    // Added SPA Legal Free (String)
			SPADisbursementFree:  "Yes",    // Added SPA Disbursement Free (String)
			LoanAgreementFree:    "Yes",    // Added Loan Agreement Free (String)
			LoanDisbursementFree: "No",     // Added Loan Disbursement Free (String)
			LoanStampDutyFree:    "No",     // Added Loan Stamp Duty Free (String)
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
			Name:        "John Doe",
			Contact:     "+60 11-222 3333",
			Citizenship: "Non-Malaysian",
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
			CarParkLot:  "1 (Basement 2)",
		},
		Options: []models.Option{
			{
				OptionName:        "Essential Entry Pack",
				Rebate:            45000.00,
				RebatePercentage:  10.00,
				Cashback:          0.00,
				CashbackType:      "None",
				DownPayment:       0.00, // 0 Downpayment promo
				NettPrice:         405000.00,
				LoanAmount:        405000.00,
				InterestRate:      4.10,
				MonthlyInstalment: 1850.20,
				LoanTenureYear:    30,
				Furnishing: models.Furnishing{
					KitchenCabinet:      true,
					Airconds:            1,
					Heater:              true,
					BathroomAccessories: false,
					LightFixtures:       true,
				},
			},
		},
		LegalAndFees: models.LegalFees{
			MaintenanceFeePSF:    0.40,
			MaintenanceFeeTotal:  220.00,
			Included:             []string{"SPA Legal Fee"},
			NotIncluded:          []string{"Loan Stamp Duty", "MOT"},
			MOT:                  0.00,
			SPALegalFree:         "Yes",
			SPADisbursementFree:  "No",
			LoanAgreementFree:    "No",
			LoanDisbursementFree: "No",
			LoanStampDutyFree:    "No",
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

// Test for minimal data (Edge Case: Missing optional fields, zero values)
func TestGenerateQuotationMinimumDataPDF(t *testing.T) {
	quote := models.PropertyQuotation{
		AppointmentDate:   time.Now(),
		QuotationValidity: time.Now().AddDate(0, 0, 7), // Valid for 7 days
		LeadInfo: models.Lead{
			Name:        "Ali Bin Abu",
			Contact:     "010-1111111",
			Citizenship: "Malaysian",
		},
		ProjectDetails: models.Project{
			ProjectName: "Basic Apartment",
			Developer:   "Standard Dev",
			Tenure:      "Leasehold",
			UnitNo:      "1-1",
			Facing:      "Unknown",
			LayoutType:  "Type A",
			AreaSqft:    800,
			SPAPrice:    300000.00,
			CarParkLot:  "1",
		},
		Options: []models.Option{
			{
				OptionName:        "No Frills Option",
				Rebate:            0.00,
				RebatePercentage:  0.00,
				Cashback:          0.00,
				CashbackType:      "None",
				DownPayment:       30000.00, // Standard 10%
				NettPrice:         300000.00,
				LoanAmount:        270000.00, // 90% loan
				InterestRate:      4.00,
				MonthlyInstalment: 1289.00,
				LoanTenureYear:    35,
				Discounts:         []models.Discount{}, // Empty discounts
				Furnishing: models.Furnishing{
					// All false / 0 / empty
					KitchenCabinet:      false,
					HoodAndHob:          false,
					Fridge:              false,
					WashingMachine:      0,
					Airconds:            0,
					Toilet:              false,
					Heater:              false,
					ShowerScreen:        false,
					WardrobeQty:         0,
					BedSetQty:           0,
					Additional:          []string{},
					BathroomAccessories: false,
					LightFixtures:       false,
				},
			},
		},
		LegalAndFees: models.LegalFees{
			MaintenanceFeePSF:    0.20,
			MaintenanceFeeTotal:  160.00,
			Included:             []string{}, // Nothing included
			NotIncluded:          []string{"All Legal Fees", "MOT", "Valuation"},
			MOT:                  5000.00,
			SPALegalFree:         "No",
			SPADisbursementFree:  "No",
			LoanAgreementFree:    "No",
			LoanDisbursementFree: "No",
			LoanStampDutyFree:    "No",
		},
		Agent: models.Agent{
			Name:        "New Agent",
			PhoneNumber: "012-0000000",
			Email:       "agent@agency.com",
			Logo:        "", // Testing empty image URLs
			Signature:   "",
		},
	}

	outputPath := "sample_quotation_minimum_data.pdf"
	err := document.GeneratePDFDocument(quote, outputPath)
	if err != nil {
		t.Fatalf("Expected no error for minimal data, but got: %v", err)
	}
	t.Logf("Generated Minimum Data PDF: %s", outputPath)
}

// Test for high value and complex data (Edge Case: Luxury property, foreign buyer, maxed out fields)
func TestGenerateQuotationHighValueComplexPDF(t *testing.T) {
	quote := models.PropertyQuotation{
		AppointmentDate:   time.Now(),
		QuotationValidity: time.Now().AddDate(0, 3, 0), // Valid for 3 months
		LeadInfo: models.Lead{
			Name:        "Alexander Sterling",
			Contact:     "+44 7911 123456",
			Citizenship: "British", // Foreign buyer scenario
		},
		ProjectDetails: models.Project{
			ProjectName: "The Pinnacle Penthouse",
			Developer:   "Luxury Horizons Group",
			Tenure:      "Freehold",
			UnitNo:      "PH-99-01",
			Facing:      "Panoramic City Skyline",
			LayoutType:  "Penthouse - 5B5R4CP",
			AreaSqft:    6500,
			SPAPrice:    8500000.00, // 8.5 Million
			CarParkLot:  "4 (Private Garage)",
		},
		Options: []models.Option{
			{
				OptionName:        "VVIP Signature Package",
				Rebate:            850000.00,
				RebatePercentage:  10.00,
				Cashback:          150000.00,
				CashbackType:      "Offset to Loan", // Testing different cashback type
				DownPayment:       850000.00,
				NettPrice:         7650000.00,
				LoanAmount:        6000000.00, // Large loan amount
				InterestRate:      3.85,
				MonthlyInstalment: 28150.75,
				LoanTenureYear:    25, // Shorter tenure for foreigners usually
				Discounts: []models.Discount{
					{Type: "VVIP Direct Discount", Percentage: 2.00},
					{Type: "Foreign Exchange Subsidy", Percentage: 5.00},
					{Type: "Referral Bonus", Percentage: 1.00},
				},
				Furnishing: models.Furnishing{
					KitchenCabinet:      true,
					HoodAndHob:          true,
					Fridge:              true,
					WashingMachine:      2,
					Airconds:            8, // High quantities
					Toilet:              true,
					Heater:              true,
					ShowerScreen:        true,
					WardrobeQty:         5,
					BedSetQty:           5,
					Additional:          []string{"Imported Marble Flooring", "Private Elevator Lobby", "Smart Home Integration", "Wine Cellar", "Jacuzzi"},
					BathroomAccessories: true,
					LightFixtures:       true,
				},
			},
		},
		LegalAndFees: models.LegalFees{
			MaintenanceFeePSF:    0.85,
			MaintenanceFeeTotal:  5525.00,
			Included:             []string{"SPA Legal Fee", "Loan Legal Fee", "Foreign State Consent Fee"},
			NotIncluded:          []string{"Valuation Fee", "MOT"},
			MOT:                  250000.00, // Massive MOT
			SPALegalFree:         "Yes",
			SPADisbursementFree:  "Yes",
			LoanAgreementFree:    "Yes",
			LoanDisbursementFree: "Yes",
			LoanStampDutyFree:    "No", // Foreigners usually pay stamp duty
		},
		Agent: models.Agent{
			Name:        "James Bond",
			PhoneNumber: "+60 12-007 0007",
			Email:       "james.bond@luxuryprop.com",
			Logo:        "https://dummyimage.com/150x50/gold/black&text=LUXURY+ESTATES",
			Signature:   "https://dummyimage.com/200x80/ffffff/000000&text=JB+Signature",
		},
	}

	outputPath := "sample_quotation_high_value.pdf"
	err := document.GeneratePDFDocument(quote, outputPath)
	if err != nil {
		t.Fatalf("Expected no error for high value data, but got: %v", err)
	}
	t.Logf("Generated High Value PDF: %s", outputPath)
}
