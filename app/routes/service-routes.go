package routes

import (
	"glamour_reserve/handlers"
	"glamour_reserve/repositories"
	"glamour_reserve/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ServicesRoutes(app *echo.Echo, db *gorm.DB) {

	repo := repositories.NewServiceRepository(db)
	servive := services.NewSvcService(repo)
	handler := handlers.NewServiceHandler(servive)

	app.GET("/services", handler.GetAllServices)
	app.POST("/services", handler.CreateService)
	app.GET("/services/:id", handler.GetServiceByID)
	app.DELETE("/services/:id", handler.DeletByID)
	app.PUT("/services/:id", handler.UpdateByID)

}
