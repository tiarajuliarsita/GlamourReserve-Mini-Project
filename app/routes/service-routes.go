package routes

import (
	"glamour_reserve/handlers"
	"glamour_reserve/helpers"
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

	e := app.Group("")
	app.GET("/services/:id", handler.GetServiceByID)
	e.Use(helpers.Middleware())
	e.POST("/services", handler.CreateService)
	e.DELETE("/services/:id", handler.DeletByID)
	e.PUT("/services/:id", handler.UpdateByID)

}
