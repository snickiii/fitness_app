package utils

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var JwtKey = []byte("")

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(userID uint, username, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		log.Printf("Error signing JWT: %v", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		log.Printf("Error parsing JWT: %v", err)
		return nil, err
	}
	if !token.Valid {
		log.Println("Invalid JWT token")
		return nil, err
	}
	return claims, nil
}
