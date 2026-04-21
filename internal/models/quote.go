package models

import "time"

// PropertyQuotation represents the root document
type PropertyQuotation struct {
	AppointmentDate   time.Time `json:"appointment_date"`
	QuotationValidity time.Time `json:"quotation_validity"`
	LeadInfo          Lead      `json:"lead_info"`
	ProjectDetails    Project   `json:"project_details"`
	Options           []Option  `json:"options"`
	LegalAndFees      LegalFees `json:"legal_and_fees"`
	Agent             Agent     `json:"agent"`
}

type Lead struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

type Project struct {
	ProjectName string  `json:"project_name"`
	Developer   string  `json:"developer"`
	Tenure      string  `json:"tenure"` // e.g., "Freehold" or "Leasehold"
	UnitNo      string  `json:"unit_no"`
	Facing      string  `json:"facing"`
	LayoutType  string  `json:"layout_type"`
	AreaSqft    int     `json:"area_sqft"`
	SPAPrice    float64 `json:"spa_price"`
}

type Option struct {
	OptionName        string     `json:"option_name"`
	Rebate            float64    `json:"rebate"`
	Discounts         []Discount `json:"other_discounts"` // For "add different type of discount"
	Cashback          float64    `json:"cashback"`
	DownPayment       float64    `json:"down_payment"`
	NettPrice         float64    `json:"nett_price"`
	LoanAmount        float64    `json:"loan_amount"`
	InterestRate      float64    `json:"interest_rate"` // e.g., 4.25
	MonthlyInstalment float64    `json:"monthly_instalment"`
	Furnishing        Furnishing `json:"furnishing"`
}

type Discount struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type Furnishing struct {
	KitchenCabinet bool     `json:"kitchen_cabinet"`
	HoodAndHob     bool     `json:"hood_and_hob"`
	Fridge         bool     `json:"fridge"`
	WashingMachine int      `json:"washing_machine_qty"`
	Airconds       int      `json:"airconds_qty"`
	Toilet         bool     `json:"toilet"`
	Heater         bool     `json:"heater"`
	ShowerScreen   bool     `json:"shower_screen"`
	WardrobeQty    int      `json:"wardrobe_qty"`
	BedSetQty      int      `json:"bed_set_qty"`
	Additional     []string `json:"additional_items"` // Free text for agent input
}

type LegalFees struct {
	MaintenanceFeePSF   float64  `json:"maintenance_fee_psf"`
	MaintenanceFeeTotal float64  `json:"maintenance_fee_total"`
	Included            []string `json:"included"` // e.g., ["SPA Legal Fee", "MOT"]
	NotIncluded         []string `json:"not_included"`
}

type Agent struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Signature   string `json:"signature_url"` // "https://storage.com/sig_123.png"
	Logo        string `json:"logo_url"`
}
