package routes

import (
	"glamour_reserve/features/handlers"
	"glamour_reserve/features/repositories"
	"glamour_reserve/features/services"
	"glamour_reserve/utils/helpers/authentication"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(app *echo.Echo, db *gorm.DB) {

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userService)

	user := app.Group("users")
	{
		user.POST("/register", userHandlers.RegisterHandler)
		user.POST("/login", userHandlers.LoginUser)

	}
	
	beautyService := services.NewBeautyCare()
	beautyHandlers := handlers.NewBeautyCare(beautyService)
	app.POST("/ask-about-beauty-care", beautyHandlers.AskAboutBeauty)

	admin := app.Group("/users")
	admin.Use(authentication.Middleware())
	admin.GET("", userHandlers.GetAllUsers)

	

}
