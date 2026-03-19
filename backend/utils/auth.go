package utils

import (

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string (bytes), err
}

func CheckPasswordHash(Password, Hash string) bool {
		err := bcrypt.CompareHashAndPassword([]byte(Hash), []byte(Password))
		return err == nil
}

func GenerateToken(userID string, role string) (string, error) {
		secret := os.Getenv("JWT_SECRET")
		claims := jwt.MapClaims{
			"user_id": userID,
			"role": role,
			"exp": time.Now().Add(time.Hour * 72).Unix(),		
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(secret))
}

