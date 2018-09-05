package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// jwtSecret hide .evn.docker
var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

// GenerateToken ...
func GenerateToken(userID string) (string, error) {
	// Create token use jwt and set some claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

// ValidToken ...
func ValidToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return token, err
}
