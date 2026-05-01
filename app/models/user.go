package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	Username  string         `json:"username"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"-"`
	Scans     []Scan         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"scans,omitempty"`
	EmailReviews []EmailReview `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"email_reviews,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
