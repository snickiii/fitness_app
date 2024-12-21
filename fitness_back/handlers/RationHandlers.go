package handlers

import (
	"encoding/json"
	"fitness_back/DTO"
	"fitness_back/models"
	"fitness_back/utils"
	"log"
	"net/http"
)

// RationHistoryHandler возвращает информацию о рационе пользователя
// @Summary Get user profile
// @Description Retrieves the ration history of the authenticated user
// @Tags Ration
// @Accept json
// @Produce json
// @Success 200 {object} DTO.RationHistory "User ration history retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /user/ration-history [get]
// @Security BearerAuth
func RationHistoryHandler(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var user models.User
	if err := db.Preload("DailyRation").First(&user, claims.UserID).Error; err != nil {
		log.Printf("User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	RationHistory := DTO.RationHistory{
		RationHistory: user.DailyRation,
	}

	json.NewEncoder(w).Encode(RationHistory)
}

// CreateMeal создает новый прием пищи и связывает его с текущим пользователем
// @Summary Create a new meal and associate it with the current user
// @Description Create a new meal using the provided data and associate it with the authenticated user based on JWT claims.
// @Tags Ration
// @Accept  json
// @Produce  json
// @Param ration body models.DailyRation true "Ration details"
// @Success 201 {object} models.DailyRation "Meal add successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /user/add-meal [post]
// @Security BearerAuth
func CreateMeal(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var meal models.DailyRation
	if err := json.NewDecoder(r.Body).Decode(&meal); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	meal.UserID = claims.UserID

	if err := db.Create(&meal).Error; err != nil {
		log.Printf("Error add meal: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(meal)
}

// FoodDataHandler
// @Summary Получить информацию о пище
// @Description Запрашивает данные о пище с API и возвращает информацию в формате JSON
// @Tags Ration
// @Accept json
// @Produce json
// @Success 200 {object} FindResponse
// @Failure 500 {object} map[string]string "error message"
// @Router /search-food [get]
func FoodDataHandler(w http.ResponseWriter, r *http.Request) {

	searchExp := r.URL.Query().Get("searchExp")

	if searchExp == "" {
		http.Error(w, "Missing searchExp parameter", http.StatusBadRequest)
		return
	}

	data, err := utils.FetchFoodDataFromAPI(searchExp)
	if err != nil {
		log.Printf("Error fetching food data: %v", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
