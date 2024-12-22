package handlers

import (
	"encoding/json"
	"fitness_back/DTO"
	"fitness_back/models"
	"fitness_back/utils"
	"log"
	"net/http"
)

// CharsHistoryHandler возвращает информацию о характеристиках пользователя
// @Summary Get user profile
// @Description Retrieves the chars history of the authenticated user
// @Tags Ration
// @Accept json
// @Produce json
// @Success 200 {object} DTO.RationHistory "User chars history retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /user/char-history [get]
// @Security BearerAuth
func CharsHistoryHandler(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var user models.User
	if err := db.Preload("UserCharacteristics").First(&user, claims.UserID).Error; err != nil {
		log.Printf("User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	CharHistory := DTO.CharHistory{
		CharHistory: user.UserCharacteristics,
	}

	json.NewEncoder(w).Encode(CharHistory)
}

// CreateChars создает новый прием пищи и связывает его с текущим пользователем
// @Summary Create a new Chars and associate it with the current user
// @Description Create a new Chars using the provided data and associate it with the authenticated user based on JWT claims.
// @Tags Ration
// @Accept  json
// @Produce  json
// @Param ration body models.DailyRation true "Ration details"
// @Success 201 {object} models.DailyRation "Chars add successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /user/add-Chars [post]
// @Security BearerAuth
func CreateChars(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var chars models.UserCharacteristics
	if err := json.NewDecoder(r.Body).Decode(&chars); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	chars.UserID = claims.UserID

	if err := db.Create(&chars).Error; err != nil {
		log.Printf("Error add Chars: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(chars)
}
