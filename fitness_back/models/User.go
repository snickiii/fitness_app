package models

type User struct {
	UserID              uint   `gorm:"primary_key;auto_increment"`
	Email               string `gorm:"unique;not null"`
	Name                string
	SurName             string
	Username            string                `gorm:"unique;not null"`
	Password            string                `gorm:"not null"`
	UserCharacteristics []UserCharacteristics `gorm:"foreignKey:UserID"`
	DailyRation         []DailyRation         `gorm:"foreignKey:UserID"`
}
