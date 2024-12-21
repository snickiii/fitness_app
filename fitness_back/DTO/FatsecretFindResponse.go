package DTO

type FatsecretMeal struct {
	FoodDescription string
	FoodID          uint
	FoodName        string
	Portion         string
	Calories        float64
	Fat             float64
	Carbs           float64
	Protein         float64
}

type FindResponse struct {
	FindResponse []FatsecretMeal
}
