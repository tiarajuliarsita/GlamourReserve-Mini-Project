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
	svc:=services.NewSvcService(repo)
	handler:=handlers.NewServiceHandler(svc)
	

	
	app.GET("/sevices", handler.GetAllServices)

}
