package routes

import (
	"glamour_reserve/features/handlers"
	"glamour_reserve/features/repositories"
	"glamour_reserve/features/services"
	"glamour_reserve/utils/helpers/authentication"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ServicesRoutes(app *echo.Echo, db *gorm.DB) {

	repo := repositories.NewServiceRepository(db)
	servive := services.NewSvcService(repo)
	handler := handlers.NewServiceHandler(servive)

	//all
	app.GET("services/", handler.GetAllServices)
	app.GET("services/id", handler.GetServiceByID)

	//admine
	e := app.Group("/services")
	e.Use(authentication.Middleware())
	e.POST("", handler.CreateService)
	e.DELETE("/:id", handler.DeletByID)
	e.PUT("/:id", handler.UpdateByID)

}
