package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PropertyValuation represents third-party valuation data from Zillow, Redfin, etc.
type PropertyValuation struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PropertyID    uuid.UUID `json:"property_id" gorm:"type:uuid;not null;index"`
	Source        string    `json:"source" gorm:"not null;size:50;check:source IN ('Zillow', 'Redfin', 'Rentimate')" validate:"required,oneof=Zillow Redfin Rentimate"`
	ValuationType string    `json:"valuation_type" gorm:"not null;size:20;check:valuation_type IN ('market_value', 'rental_estimate')" validate:"required,oneof=market_value rental_estimate"`
	Value         float64   `json:"value" gorm:"type:decimal(12,2);not null;check:value > 0" validate:"required,gt=0"`
	ValuationDate time.Time `json:"valuation_date" gorm:"type:date;not null" validate:"required"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	Property Property `json:"property,omitempty" gorm:"foreignKey:PropertyID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID if not provided
func (pv *PropertyValuation) BeforeCreate(tx *gorm.DB) (err error) {
	if pv.ID == uuid.Nil {
		pv.ID = uuid.New()
	}
	return
}

// TableName specifies the table name for GORM
func (PropertyValuation) TableName() string {
	return "property_valuations"
}

// IsMarketValue returns true if this is a market value valuation
func (pv *PropertyValuation) IsMarketValue() bool {
	return pv.ValuationType == "market_value"
}

// IsRentalEstimate returns true if this is a rental estimate valuation
func (pv *PropertyValuation) IsRentalEstimate() bool {
	return pv.ValuationType == "rental_estimate"
}
