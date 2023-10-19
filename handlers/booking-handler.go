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
	bookResp = h.AssignValuePriceAndNameDetailsBook(bookResp, dataResp.DetailsBook)

	return response.RespondJSON(e, 201, "succes", bookResp)

}

func (h *bookingHandler) GetAllHistories(e echo.Context) error {
	userId, _, _ := helpers.ExtractTokenUserId(e)

	dataBookings, err := h.bookingSvc.GetAllHistories(userId)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), dataBookings)
	}

	bookingsData := []response.BookingRespon{}
	for _, v := range dataBookings {
		bookingsResp := core.BookCoreToBookResp(v)
		bookingsResp=h.AssignValuePriceAndNameDetailsBook(bookingsResp,v.DetailsBook)
		bookingsData = append(bookingsData, bookingsResp)
	}

	return response.RespondJSON(e, 200, "succes", bookingsData)
}



func (h *bookingHandler) GetSpecificHistory(e echo.Context) error {
	userId, _, _ := helpers.ExtractTokenUserId(e)
	bookingId := e.Param("id")

	data, err := h.bookingSvc.GetSpecificHistory(bookingId, userId)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), data)
	}

	bookResp := core.BookCoreToBookResp(data)
	bookResp = h.AssignValuePriceAndNameDetailsBook(bookResp, data.DetailsBook)
	return response.RespondJSON(e, 200, "succes", bookResp)
}

func (h *bookingHandler) FindBookingByID(e echo.Context) error {
	id := e.Param("id")
	data, err := h.bookingSvc.FindServiceByID(id)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), data)
	}
	return response.RespondJSON(e, 200, "succes", data)
}

func (h *bookingHandler) AssignValuePriceAndNameDetailsBook(response response.BookingRespon, data []core.DetailsBookCore) response.BookingRespon {
	for _, v := range data {
		booking := core.DetailsBookCoreToDetailsBookResp(v)
		dataService, _ := h.bookingSvc.FindServiceByID(v.ServiceID)
		booking.Name = dataService.Name
		booking.Price = dataService.Price
		response.DetailsBooking = append(response.DetailsBooking, booking)
	}
	return response
}
