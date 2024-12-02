package models

type Dish struct {
	DishID   uint   `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"not null"`
	Calories string `gorm:"not null"`
}
