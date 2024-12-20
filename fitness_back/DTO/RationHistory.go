package DTO

import "fitness_back/models"

type RationHistory struct {
	RationHistory []models.DailyRation `json:"objects"`
}
