package jwt

import (
	"DobryySoul/project-with-API-interaction/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *entity.User, JWTSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(JWTSecret))
}
