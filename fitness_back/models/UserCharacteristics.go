package models

import "time"

type UserCharacteristics struct {
	UserCharacteristicsID uint      `gorm:"primary_key;auto_increment"`
	UserID                uint      `gorm:"not null"`
	Ration                string    `gorm:"not null"`
	LowerStr              string    `gorm:"not null"`
	Flexibility           string    `gorm:"not null"`
	Endurance             string    `gorm:"not null"`
	Height                string    `gorm:"not null"`
	Weight                string    `gorm:"not null"`
	IMT                   string    `gorm:"not null"`
	CreatedAt             time.Time `gorm:"autoCreateTime"`
}
