package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FinancialMetrics represents calculated investment metrics for each property
type FinancialMetrics struct {
	ID                     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PropertyID             uuid.UUID `json:"property_id" gorm:"type:uuid;unique;not null;index"`
	MonthlyMortgagePayment *float64  `json:"monthly_mortgage_payment" gorm:"type:decimal(10,2)"`
	NetOperatingIncome     *float64  `json:"net_operating_income" gorm:"type:decimal(10,2)"`
	CapRate                *float64  `json:"cap_rate" gorm:"type:decimal(5,2);index"`
	CashOnCashReturn       *float64  `json:"cash_on_cash_return" gorm:"type:decimal(5,2);index"`
	CashToClose            *float64  `json:"cash_to_close" gorm:"type:decimal(12,2)"`
	RentToValueRatio       *float64  `json:"rent_to_value_ratio" gorm:"type:decimal(5,2)"`
	GrossRentMultiplier    *float64  `json:"gross_rent_multiplier" gorm:"type:decimal(5,2)"`
	CalculatedAt           time.Time `json:"calculated_at" gorm:"autoCreateTime"`
	IsCurrent              bool      `json:"is_current" gorm:"default:true"`

	// Relationships
	Property Property `json:"property,omitempty" gorm:"foreignKey:PropertyID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID if not provided
func (fm *FinancialMetrics) BeforeCreate(tx *gorm.DB) (err error) {
	if fm.ID == uuid.Nil {
		fm.ID = uuid.New()
	}
	return
}

// TableName specifies the table name for GORM
func (FinancialMetrics) TableName() string {
	return "financial_metrics"
}

// SetFloat64 safely sets a float64 pointer value
func (fm *FinancialMetrics) SetFloat64(field **float64, value float64) {
	*field = &value
}

// GetFloat64 safely gets a float64 pointer value, returning 0 if nil
func (fm *FinancialMetrics) GetFloat64(field *float64) float64 {
	if field == nil {
		return 0
	}
	return *field
}

// IsComplete returns true if all essential metrics are calculated
func (fm *FinancialMetrics) IsComplete() bool {
	return fm.NetOperatingIncome != nil &&
		fm.CapRate != nil &&
		fm.CashOnCashReturn != nil &&
		fm.RentToValueRatio != nil &&
		fm.GrossRentMultiplier != nil
}

// MarkAsOutdated marks the metrics as no longer current
func (fm *FinancialMetrics) MarkAsOutdated() {
	fm.IsCurrent = false
}
