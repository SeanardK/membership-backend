package model

import (
	"time"
)

type Redemption struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	RedemptionDate time.Time `json:"redemption_date" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type RedemptionRequest struct {
	UserID     uint   `json:"user_id" form:"user_id" binding:"required"`
}

func (Redemption) TableName() string {
	return "Redemptions"
}
