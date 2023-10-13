package routes

import (
	"glamour_reserve/handlers"
	"glamour_reserve/repositories"
	"glamour_reserve/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(app *echo.Echo, db *gorm.DB) {
	 
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handlers := handlers.NewUserHandler(service)

	app.POST("/users/register", handlers.RegisterHandler)
	app.POST("/users/login", handlers.LoginUser)
}
