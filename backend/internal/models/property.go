package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// JSONB is a custom type for PostgreSQL JSONB fields
type JSONB map[string]interface{}

// Scan implements the Scanner interface for database/sql
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONB)
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, j)
	case string:
		return json.Unmarshal([]byte(v), j)
	default:
		*j = make(JSONB)
		return nil
	}
}

// Value implements the Valuer interface for database/sql
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	return json.Marshal(j)
}

// Property represents a rental property with all investment-related data
type Property struct {
	ID                   uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID               uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Address              string    `json:"address" gorm:"not null;size:255" validate:"required,max=255"`
	YearBuilt            *int      `json:"year_built" gorm:"check:year_built >= 1800 AND year_built <= EXTRACT(YEAR FROM NOW()) + 1"`
	LandAreaSqft         *int      `json:"land_area_sqft" gorm:"check:land_area_sqft > 0"`
	BuildingAreaSqft     *int      `json:"building_area_sqft" gorm:"check:building_area_sqft > 0"`
	PurchasePrice        float64   `json:"purchase_price" gorm:"type:decimal(12,2);not null" validate:"required,gt=0"`
	IntendedRent         *float64  `json:"intended_rent" gorm:"type:decimal(10,2)"`
	OperatingExpenses    JSONB     `json:"operating_expenses" gorm:"type:jsonb;default:'{}'"`
	FinancingTerms       JSONB     `json:"financing_terms" gorm:"type:jsonb;default:'{}'"`
	OperatingAssumptions JSONB     `json:"operating_assumptions" gorm:"type:jsonb;default:'{}'"`
	LocalContext         JSONB     `json:"local_context" gorm:"type:jsonb;default:'{}'"`
	CreatedAt            time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User             User                `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments         []Comment           `json:"comments,omitempty" gorm:"foreignKey:PropertyID;constraint:OnDelete:CASCADE"`
	FinancialMetrics *FinancialMetrics   `json:"financial_metrics,omitempty" gorm:"foreignKey:PropertyID;constraint:OnDelete:CASCADE"`
	Valuations       []PropertyValuation `json:"valuations,omitempty" gorm:"foreignKey:PropertyID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID if not provided
func (p *Property) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

// TableName specifies the table name for GORM
func (Property) TableName() string {
	return "properties"
}

// HasRequiredFieldsForMetrics checks if property has all required fields for financial calculations
func (p *Property) HasRequiredFieldsForMetrics() bool {
	return p.PurchasePrice > 0 &&
		p.IntendedRent != nil && *p.IntendedRent > 0 &&
		p.OperatingExpenses != nil &&
		p.FinancingTerms != nil &&
		p.OperatingAssumptions != nil
}

// GetOperatingExpense safely gets an operating expense value
func (p *Property) GetOperatingExpense(key string) float64 {
	if p.OperatingExpenses == nil {
		return 0
	}
	if val, ok := p.OperatingExpenses[key]; ok {
		if fval, ok := val.(float64); ok {
			return fval
		}
	}
	return 0
}

// GetFinancingTerm safely gets a financing term value
func (p *Property) GetFinancingTerm(key string) float64 {
	if p.FinancingTerms == nil {
		return 0
	}
	if val, ok := p.FinancingTerms[key]; ok {
		if fval, ok := val.(float64); ok {
			return fval
		}
	}
	return 0
}

// GetOperatingAssumption safely gets an operating assumption value
func (p *Property) GetOperatingAssumption(key string) float64 {
	if p.OperatingAssumptions == nil {
		return 0
	}
	if val, ok := p.OperatingAssumptions[key]; ok {
		if fval, ok := val.(float64); ok {
			return fval
		}
	}
	return 0
}
