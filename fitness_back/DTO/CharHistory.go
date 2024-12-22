package DTO

import "fitness_back/models"

type CharHistory struct {
	CharHistory []models.UserCharacteristics `json:"objects"`
}
