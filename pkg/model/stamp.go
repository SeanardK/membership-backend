package model

import (
	"time"
)

type Stamp struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	EarnedDate time.Time `json:"earned_date"`
	Redeemed   bool      `json:"redeemed" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type StampRequest struct {
	UserID     uint   `json:"user_id" form:"user_id" binding:"required"`
	EarnedDate string `json:"earned_date" form:"earned_date" binding:"required"`
}

func (Stamp) TableName() string {
	return "Stamps"
}
