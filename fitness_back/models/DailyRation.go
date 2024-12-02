package models

import "time"

type DailyRation struct {
	DailyRationID uint      `gorm:"primary_key;auto_increment"`
	UserID        uint      `gorm:"not null"`
	Ration        []Dish    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}
