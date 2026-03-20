package model

import (
	"time"
)

type Stamp struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	EarnedDate time.Time `json:"earned_date" gorm:"default:CURRENT_TIMESTAMP"`
	Redeemed   bool      `json:"redeemed" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type StampRequest struct {
	UserID     uint   `json:"user_id" form:"user_id" binding:"required"`
}

func (Stamp) TableName() string {
	return "Stamps"
}
