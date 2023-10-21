package routes

import (
	"glamour_reserve/features/handlers"
	"glamour_reserve/features/repositories"
	"glamour_reserve/features/services"
	"glamour_reserve/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BookingRoutes(app *echo.Echo, db *gorm.DB) {

	repo := repositories.NewBookingRepository(db)
	service := services.NewBookingService(repo)
	handler := handlers.NewBookingHandler(service)

	user := app.Group("/users/bookings")
	user.Use(helpers.Middleware())

	//user
	user.POST("", handler.CreateBooking)
	user.GET("", handler.GetAllHistories)
	user.GET("/:id", handler.GetSpecificHistory)

	//admin
	admin:= app.Group("/admin/bookings")
	admin.GET("/:id", handler.FindBookingByID)
	admin.PUT("/:invoice", handler.UpdateStatusBooking)
	admin.GET("", handler.GetAllBookings)

}
