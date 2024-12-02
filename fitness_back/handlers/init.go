package handlers

import (
	"fitness_back/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
	db.AutoMigrate(&models.User{})
}
