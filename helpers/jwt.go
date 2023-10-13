package helpers

import (
	"glamour_reserve/app/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id string, userName string, email string) string {
	claims := jwt.MapClaims{
		"id":        id,
		"email":     email,
		"user_name": userName,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cfg := config.AppConfig{}
	signedToken, err := token.SignedString([]byte(cfg.SECRETKEY))
	if err != nil {
		return err.Error()
	}
	return signedToken
}
