package handlers

import (
	"auth-service/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
	db.AutoMigrate(&models.User{}, &models.Target{})
}
