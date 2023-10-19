package handlers

import (
	core "glamour_reserve/entity/core"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
	"glamour_reserve/helpers"
	"glamour_reserve/services"

	"github.com/labstack/echo/v4"
)

type bookingHandler struct {
	bookingSvc services.BookingServiceInterface
}

func NewBookingHandler(bookingSvc services.BookingServiceInterface) *bookingHandler {
	return &bookingHandler{bookingSvc}
}

func (h *bookingHandler) CreateBooking(e echo.Context) error {
	//extract token
	userId, userName, _ := helpers.ExtractTokenUserId(e)

	bookingReq := request.BookingRequest{}

	err := e.Bind(&bookingReq)

	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	detailBookings := []core.DetailsBookCore{}

	for _, v := range bookingReq.Details {
		data := core.BookingDataRequestToDetailsBookingCore(v)
		detailBookings = append(detailBookings, data)
	}

	bookingInsert := core.BookingCore{}
	bookingInsert.DetailsBook = detailBookings
	bookingInsert.UserID = userId

	dataResp, err := h.bookingSvc.Create(bookingInsert)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	bookResp := core.BookCoreToBookResp(dataResp)
	bookResp.Name = userName

	for _, v := range dataResp.DetailsBook {
		booking := core.DetailsBookCoreToDetailsBookResp(v)
		dataService, _ := h.bookingSvc.FindServiceByID(v.ServiceID)
		booking.Name = dataService.Name
		booking.Price = dataService.Price
		bookResp.DetailsBooking = append(bookResp.DetailsBooking, booking)

	}

	return response.RespondJSON(e, 201, "succes", bookResp)

}
