package handlers

import (
	"fitness_back/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.DailyRation{})
	db.AutoMigrate(&models.UserCharacteristics{})
}
