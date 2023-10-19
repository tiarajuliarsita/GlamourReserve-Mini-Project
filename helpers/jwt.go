package helpers

import (
	"glamour_reserve/app/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func GenerateToken(id string, userName string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = id
	claims["user_name"] = userName
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cfg := config.AppConfig{}
	return token.SignedString([]byte(cfg.SECRET_KEY))
}

func ExtractTokenUserId(e echo.Context) (string, string, string) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		userName := claims["user_name"].(string)
		role := claims["role"].(string)
		return userId, userName, role
	}
	return "", "", ""
}

func Middleware() echo.MiddlewareFunc {
	cfg := config.AppConfig{}
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(cfg.SECRET_KEY),
		SigningMethod: "HS256",
	})
}
