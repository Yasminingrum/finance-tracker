package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken generates a JWT token for a user
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
