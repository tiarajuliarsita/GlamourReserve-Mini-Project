package routes

import (
	"glamour_reserve/handlers"
	"glamour_reserve/repositories"
	"glamour_reserve/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func VariantRoutes(app *echo.Echo, db *gorm.DB) {
	repo := repositories.NewVariantRepository(db)
	service := services.NewVariantService(repo)
	handler := handlers.NewVariantHandler(service)

	app.POST("/variants", handler.CreateVariant)
	app.GET("/variants/:id", handler.GetByID)
	app.GET("/variants", handler.GetAll)
	app.DELETE("/variants/:id", handler.Delete)
	app.PUT("/variants/:id", handler.Update)
}
