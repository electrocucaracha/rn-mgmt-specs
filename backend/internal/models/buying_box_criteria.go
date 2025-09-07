package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BuyingBoxCriteria represents user-defined investment criteria for property evaluation
type BuyingBoxCriteria struct {
	ID                      uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID                  uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Name                    string    `json:"name" gorm:"not null;size:100" validate:"required,max=100"`
	MinCapRate              *float64  `json:"min_cap_rate" gorm:"type:decimal(5,2)"`
	MinCashOnCash           *float64  `json:"min_cash_on_cash" gorm:"type:decimal(5,2)"`
	MaxPurchasePrice        *float64  `json:"max_purchase_price" gorm:"type:decimal(12,2)"`
	MinRentToValue          *float64  `json:"min_rent_to_value" gorm:"type:decimal(5,2)"`
	MaxYearBuilt            *int      `json:"max_year_built"`
	MinYearBuilt            *int      `json:"min_year_built" gorm:"check:min_year_built IS NULL OR max_year_built IS NULL OR min_year_built <= max_year_built"`
	LocationPreferences     JSONB     `json:"location_preferences" gorm:"type:jsonb;default:'{}'"`
	PropertyTypePreferences JSONB     `json:"property_type_preferences" gorm:"type:jsonb;default:'{}'"`
	IsActive                bool      `json:"is_active" gorm:"default:true;index"`
	CreatedAt               time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt               time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID if not provided
func (bbc *BuyingBoxCriteria) BeforeCreate(tx *gorm.DB) (err error) {
	if bbc.ID == uuid.Nil {
		bbc.ID = uuid.New()
	}
	return
}

// TableName specifies the table name for GORM
func (BuyingBoxCriteria) TableName() string {
	return "buying_box_criteria"
}

// PropertyComparison represents how a property compares against buying criteria
type PropertyComparison struct {
	Property       *Property          `json:"property"`
	Criteria       *BuyingBoxCriteria `json:"criteria"`
	Matches        map[string]bool    `json:"matches"`
	Score          float64            `json:"score"`
	FailureReasons []string           `json:"failure_reasons,omitempty"`
}

// CompareProperty evaluates a property against this buying criteria
func (bbc *BuyingBoxCriteria) CompareProperty(property *Property, metrics *FinancialMetrics) *PropertyComparison {
	comparison := &PropertyComparison{
		Property:       property,
		Criteria:       bbc,
		Matches:        make(map[string]bool),
		FailureReasons: []string{},
	}

	totalCriteria := 0
	metCriteria := 0

	// Check cap rate
	if bbc.MinCapRate != nil && metrics != nil && metrics.CapRate != nil {
		totalCriteria++
		if *metrics.CapRate >= *bbc.MinCapRate {
			comparison.Matches["cap_rate"] = true
			metCriteria++
		} else {
			comparison.Matches["cap_rate"] = false
			comparison.FailureReasons = append(comparison.FailureReasons, "Cap rate below minimum")
		}
	}

	// Check cash-on-cash return
	if bbc.MinCashOnCash != nil && metrics != nil && metrics.CashOnCashReturn != nil {
		totalCriteria++
		if *metrics.CashOnCashReturn >= *bbc.MinCashOnCash {
			comparison.Matches["cash_on_cash"] = true
			metCriteria++
		} else {
			comparison.Matches["cash_on_cash"] = false
			comparison.FailureReasons = append(comparison.FailureReasons, "Cash-on-cash return below minimum")
		}
	}

	// Check maximum purchase price
	if bbc.MaxPurchasePrice != nil {
		totalCriteria++
		if property.PurchasePrice <= *bbc.MaxPurchasePrice {
			comparison.Matches["purchase_price"] = true
			metCriteria++
		} else {
			comparison.Matches["purchase_price"] = false
			comparison.FailureReasons = append(comparison.FailureReasons, "Purchase price above maximum")
		}
	}

	// Check rent-to-value ratio
	if bbc.MinRentToValue != nil && metrics != nil && metrics.RentToValueRatio != nil {
		totalCriteria++
		if *metrics.RentToValueRatio >= *bbc.MinRentToValue {
			comparison.Matches["rent_to_value"] = true
			metCriteria++
		} else {
			comparison.Matches["rent_to_value"] = false
			comparison.FailureReasons = append(comparison.FailureReasons, "Rent-to-value ratio below minimum")
		}
	}

	// Check year built range
	if property.YearBuilt != nil {
		if bbc.MinYearBuilt != nil {
			totalCriteria++
			if *property.YearBuilt >= *bbc.MinYearBuilt {
				comparison.Matches["min_year_built"] = true
				metCriteria++
			} else {
				comparison.Matches["min_year_built"] = false
				comparison.FailureReasons = append(comparison.FailureReasons, "Property too old")
			}
		}

		if bbc.MaxYearBuilt != nil {
			totalCriteria++
			if *property.YearBuilt <= *bbc.MaxYearBuilt {
				comparison.Matches["max_year_built"] = true
				metCriteria++
			} else {
				comparison.Matches["max_year_built"] = false
				comparison.FailureReasons = append(comparison.FailureReasons, "Property too new")
			}
		}
	}

	// Calculate score
	if totalCriteria > 0 {
		comparison.Score = float64(metCriteria) / float64(totalCriteria) * 100
	} else {
		comparison.Score = 100 // No criteria defined, so property "passes"
	}

	return comparison
}

// Deactivate marks the criteria as inactive
func (bbc *BuyingBoxCriteria) Deactivate() {
	bbc.IsActive = false
}

// Activate marks the criteria as active
func (bbc *BuyingBoxCriteria) Activate() {
	bbc.IsActive = true
}
