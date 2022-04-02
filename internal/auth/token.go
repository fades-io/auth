package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Создание токена
func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Токен протухает через 1 день
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
