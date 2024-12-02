package handlers

import (
	"encoding/json"
	"fitness_back/models"
	"fitness_back/utils"
	"log"
	"net/http"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/bcrypt"
)

// Register регистрирует нового пользователя
// @Summary Register a new user
// @Description Register a new user with the given details, including hashing the password before saving to the database
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   user  body   models.User  true  "User registration data"
// @Success 201 {string} string "User registered successfully"
// // @Failure 400 {string} string "Invalid request or user already exists"
// // @Failure 500 {string} string "Internal server error"
// @Router /auth/register [post]
func Register(w http.ResponseWriter, r *http.Request) { //Добавить парольные политики
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "User already exists or data is invalid", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Printf("User %s registered successfully", user.Username)
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login аутентифицирует пользователя и выдает JWT
// @Summary Authenticate a user and issue a JWT
// @Description Authenticate a user by verifying their username and password, and return a JWT if successful
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param loginData body LoginData true "User login credentials"
// @Success 200 {object} string "JWT token"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Invalid login or password"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var loginData LoginData
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		log.Printf("Error decoding login data: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var dbUser models.User
	if err := db.Where("username = ?", loginData.Username).First(&dbUser).Error; err != nil {
		log.Printf("Error finding user: %v", err)
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginData.Password)); err != nil {
		log.Printf("Invalid password for user %s: %v", loginData.Username, err)
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(dbUser.UserID, dbUser.Username, dbUser.Email)
	if err != nil {
		log.Printf("Error generating JWT for user %s: %v", dbUser.Username, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User %s logged in successfully", dbUser.Username)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
