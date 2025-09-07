package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Comment represents user collaboration and notes on properties
type Comment struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PropertyID uuid.UUID  `json:"property_id" gorm:"type:uuid;not null;index"`
	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;index"`
	Content    string     `json:"content" gorm:"type:text;not null;check:LENGTH(TRIM(content)) > 0 AND LENGTH(content) <= 2000" validate:"required,max=2000"`
	ParentID   *uuid.UUID `json:"parent_id" gorm:"type:uuid;index"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime;index"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Property Property  `json:"property,omitempty" gorm:"foreignKey:PropertyID;constraint:OnDelete:CASCADE"`
	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Parent   *Comment  `json:"parent,omitempty" gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"`
	Replies  []Comment `json:"replies,omitempty" gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID and validate content
func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	// Validate content
	c.Content = strings.TrimSpace(c.Content)
	if len(c.Content) == 0 {
		return gorm.ErrInvalidValue
	}

	return
}

// BeforeUpdate hook to validate content
func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	// Validate content
	c.Content = strings.TrimSpace(c.Content)
	if len(c.Content) == 0 {
		return gorm.ErrInvalidValue
	}

	return
}

// TableName specifies the table name for GORM
func (Comment) TableName() string {
	return "comments"
}

// IsReply returns true if this comment is a reply to another comment
func (c *Comment) IsReply() bool {
	return c.ParentID != nil
}

// IsTopLevel returns true if this comment is not a reply
func (c *Comment) IsTopLevel() bool {
	return c.ParentID == nil
}

// GetContentPreview returns a truncated version of the content for previews
func (c *Comment) GetContentPreview(maxLength int) string {
	if len(c.Content) <= maxLength {
		return c.Content
	}
	return c.Content[:maxLength] + "..."
}
