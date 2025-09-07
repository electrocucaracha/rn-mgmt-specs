package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents an individual investor or team member
type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Email        string    `json:"email" gorm:"unique;not null;size:255" validate:"required,email"`
	PasswordHash string    `json:"-" gorm:"not null;size:255"`
	FirstName    string    `json:"first_name" gorm:"not null;size:50" validate:"required,max=50"`
	LastName     string    `json:"last_name" gorm:"not null;size:50" validate:"required,max=50"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Properties         []Property          `json:"properties,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments           []Comment           `json:"comments,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	BuyingBoxCriterias []BuyingBoxCriteria `json:"buying_criteria,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID if not provided
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}

// PublicUser returns user data safe for public consumption (no password hash)
func (u *User) PublicUser() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"email":      u.Email,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"is_active":  u.IsActive,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}
