package routes

import (
	"glamour_reserve/handlers"
	"glamour_reserve/helpers"
	"glamour_reserve/repositories"
	"glamour_reserve/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BookingRoutes(app *echo.Echo, db *gorm.DB) {

	
	repo := repositories.NewBookingRepository(db)
	service := services.NewBookingService(repo)
	handler := handlers.NewBookingHandler(service)
	
	e:=app.Group("")
	e.Use(helpers.Middleware())
	e.POST("/bookings", handler.CreateBooking)
	e.GET("/bookings",handler.GetAllHistoriesBookingsUser)
}
