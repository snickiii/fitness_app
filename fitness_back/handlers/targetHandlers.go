package handlers

import (
	"auth-service/models"
	"auth-service/utils"
	"encoding/json"
	"log"
	"net/http"
)

// CreateTarget создает новый целевой объект и связывает его с текущим пользователем
// @Summary Create a new target and associate it with the current user
// @Description Create a new target using the provided data and associate it with the authenticated user based on JWT claims.
// @Tags Target
// @Accept  json
// @Produce  json
// @Param target body models.Target true "Target details"
// @Success 201 {object} models.Target "Target created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /user/targets [post]
// @Security BearerAuth
func CreateTarget(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var target models.Target
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	target.UserID = claims.UserID

	if err := db.Create(&target).Error; err != nil {
		log.Printf("Error creating target: %v", err)
		http.Error(w, "Error creating target", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(target)
}

// UpdateTarget обновляет существующий целевой объект, связанный с текущим пользователем
// @Summary Update an existing target associated with the current user
// @Description Update the target details if the target is associated with the authenticated user
// @Tags Target
// @Accept  json
// @Produce  json
// @Param id query string true "Target ID"
// @Param target body models.Target true "Updated target details"
// @Success 200 {object} models.Target "Target updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Target not found"
// @Failure 500 {string} string "Internal server error"
// @Router /user/targets [put]
// @Security BearerAuth
func UpdateTarget(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	idStr := r.URL.Query().Get("id")

	var target models.Target
	if err := db.First(&target, idStr).Error; err != nil {
		log.Printf("Target not found: %v", err)
		http.Error(w, "Target not found", http.StatusNotFound)
		return
	}

	if target.UserID != claims.UserID {
		log.Println("Unauthorized access to target")
		http.Error(w, "Unauthorized access", http.StatusForbidden)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := db.Save(&target).Error; err != nil {
		log.Printf("Error updating target: %v", err)
		http.Error(w, "Error updating target", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(target)
}

// DeleteTarget удаляет целевой объект, связанный с текущим пользователем
// @Summary Delete a target associated with the current user
// @Description Delete the target if it is associated with the authenticated user based on JWT claims.
// @Tags Target
// @Accept  json
// @Produce  json
// @Param id query string true "Target ID"
// @Success 200 {string} string "Target deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Target not found"
// @Failure 500 {string} string "Internal server error"
// @Router /user/targets [delete]
// @Security BearerAuth
func DeleteTarget(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	idStr := r.URL.Query().Get("id")

	var target models.Target
	if err := db.First(&target, idStr).Error; err != nil {
		log.Printf("Target not found: %v", err)
		http.Error(w, "Target not found", http.StatusNotFound)
		return
	}

	if target.UserID != claims.UserID {
		log.Println("Unauthorized access to target")
		http.Error(w, "Unauthorized access", http.StatusForbidden)
		return
	}

	if err := db.Delete(&target).Error; err != nil {
		log.Printf("Error deleting target: %v", err)
		http.Error(w, "Error deleting target", http.StatusInternalServerError)
		return
	}

	log.Printf("Target %d deleted successfully", target.TargetID)
	json.NewEncoder(w).Encode(map[string]string{"message": "Target deleted successfully"})
}
