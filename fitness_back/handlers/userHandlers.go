package handlers

import (
	"auth-service/models"
	"auth-service/utils"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// ProfileHandler возвращает информацию о профиле пользователя и связанных с ним целевых объектах
// @Summary Get user profile and associated targets
// @Description Retrieves the profile of the authenticated user along with the targets associated with their account
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.User "User profile and targets retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /user/profile [get]
// @Security BearerAuth
func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var user models.User
	if err := db.Preload("Targets").First(&user, claims.UserID).Error; err != nil {
		log.Printf("User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	profile := models.ProfileResponse{
		UserID:   user.UserID,
		Email:    user.Email,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
		Targets:  user.Targets,
	}

	json.NewEncoder(w).Encode(profile)
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// UpdatePassword обновляет пароль пользователя
// @Summary Update user password
// @Description Allows an authenticated user to update their password by providing the old password and a new password
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body UpdatePasswordRequest true "Update password request"
// @Success 200 {string} string "Password updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Invalid old password"
// @Failure 403 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /user/update-password [put]
// @Security BearerAuth
func UpdatePassword(w http.ResponseWriter, r *http.Request) { //Добавить парольные политики

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var req UpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.First(&user, claims.UserID).Error; err != nil {
		log.Printf("User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		log.Println("Invalid old password")
		http.Error(w, "Invalid old password", http.StatusUnauthorized)
		return
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing new password: %v", err)
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	user.Password = string(newHashedPassword)
	if err := db.Save(&user).Error; err != nil {
		log.Printf("Error updating password: %v", err)
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password updated successfully"})
}

type UpdateEmailRequest struct {
	NewEmail string `json:"new_email"`
}

// UpdateEmail обновляет адрес электронной почты пользователя
// @Summary Update user email
// @Description Allows an authenticated user to update their email address
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body UpdateEmailRequest true "Update email request"
// @Success 200 {string} string "Email updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal server error"
// @Router /user/update-email [put]
// @Security BearerAuth
func UpdateEmail(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var req UpdateEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.First(&user, claims.UserID).Error; err != nil {
		log.Printf("User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Email = req.NewEmail
	if err := db.Save(&user).Error; err != nil {
		log.Printf("Error updating email: %v", err)
		http.Error(w, "Error updating email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email updated successfully"})
}

type UpdateNameRequest struct {
	NewName string `json:"new_name"`
}

// UpdateName обновляет имя пользователя
// @Summary Update name
// @Description Allows an authenticated user to update their name
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body UpdateNameRequest true "Update name request"
// @Success 200 {string} string "Name updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal server error"
// @Router /user/update-name [put]
// @Security BearerAuth
func UpdateName(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var req UpdateNameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.First(&user, claims.UserID).Error; err != nil {
		log.Printf("User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Name = req.NewName
	if err := db.Save(&user).Error; err != nil {
		log.Printf("Error updating name: %v", err)
		http.Error(w, "Error updating name", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Name updated successfully"})
}

type UpdateUsernameRequest struct {
	NewUsername string `json:"new_username"`
}

// UpdateUsername обновляет имя пользователя
// @Summary Update user username
// @Description Allows an authenticated user to update their username. It also checks if the new username is already taken.
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body UpdateUsernameRequest true "Update username request"
// @Success 200 {string} string "Username updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 409 {string} string "Username already taken"
// @Failure 500 {string} string "Internal server error"
// @Router /user/update-username [put]
// @Security BearerAuth
func UpdateUsername(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value("claims").(*utils.Claims)
	if !ok {
		log.Println("Error extracting claims from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var req UpdateUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.First(&user, claims.UserID).Error; err != nil {
		log.Printf("User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var existingUser models.User
	if err := db.Where("username = ?", req.NewUsername).First(&existingUser).Error; err == nil {
		log.Println("Username already taken")
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	user.Username = req.NewUsername
	if err := db.Save(&user).Error; err != nil {
		log.Printf("Error updating username: %v", err)
		http.Error(w, "Error updating username", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Username updated successfully"})
}
