package models

import (
	"time"
	"gorm.io/gorm"
)

type Scan struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `json:"user_id"`
	URL       string    `json:"url"`
	Score     int       `json:"score"`
	RiskLevel string    `json:"risk_level"` // Safe, Suspicious, High Risk, Invalid URL
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type EmailReview struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `json:"user_id"`
	InputText string    `gorm:"type:text" json:"input_text"`
	Score     int       `json:"score"`
	RiskLevel string    `json:"risk_level"` // Safe, Suspicious, High Risk
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
