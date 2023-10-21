package routes

import (
	"glamour_reserve/features/handlers"
	"glamour_reserve/features/repositories"
	"glamour_reserve/features/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(app *echo.Echo, db *gorm.DB) {

	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handlers := handlers.NewUserHandler(service)

	user :=app.Group("users")
	user.POST("/register", handlers.RegisterHandler)
	user.POST("/login", handlers.LoginUser)
	user.GET("", handlers.GetAllUsers)
}
