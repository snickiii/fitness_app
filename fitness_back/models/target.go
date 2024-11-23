package models

type Target struct {
	TargetID    uint   `gorm:"primary_key;auto_increment"`
	TargetName  string `gorm:"not null"`
	AddressPull string
	DomainPull  string
	UserID      uint `gorm:"not null"`
}
