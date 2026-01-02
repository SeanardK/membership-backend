package model

import (
	"time"
)

type Portfolio struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Image       string    `json:"image" gorm:"type:varchar(512)"`
	Description string    `json:"description" gorm:"type:text"`
	Detail      string    `json:"detail" gorm:"type:text"`
	Framework   string    `json:"framework" gorm:"type:varchar(255)"`
	Libraries   string    `json:"libraries" gorm:"type:varchar(255)"`
	Repository  string    `json:"repository" gorm:"type:varchar(512)"`
	URL         string    `json:"url" gorm:"type:varchar(512)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Portfolio) TableName() string {
	return "portfolios"
}
