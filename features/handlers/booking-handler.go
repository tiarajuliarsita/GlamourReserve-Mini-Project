package handlers

import (
	core "glamour_reserve/entity/core"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
	"glamour_reserve/utils/helpers/authentication"

	"glamour_reserve/features/services"

	"github.com/labstack/echo/v4"
)

type bookingHandler struct {
	bookingSvc services.BookingServiceInterface
}

func NewBookingHandler(bookingSvc services.BookingServiceInterface) *bookingHandler {
	return &bookingHandler{bookingSvc}
}

func (h *bookingHandler) CreateBooking(e echo.Context) error {

	userId, userName, role := authentication.ExtractTokenUserId(e)
	if role != "user" {
		return response.RespondJSON(e, 401, "can't create booking", nil)
	}

	NewBooking := request.BookingRequest{}
	err := e.Bind(&NewBooking)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	NewBookingsCore := []core.DetailsBookCore{}
	bookingInsert := core.BookingReqMap(NewBooking, NewBookingsCore, userId)

	createdBook, err := h.bookingSvc.Create(bookingInsert, userName)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	bookResp := core.BookCoreToBookResp(createdBook)
	bookResp.Name = userName

	bookResp = h.PriceAndNameValues(bookResp, createdBook.DetailsBook)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	return response.RespondJSON(e, 201, "succes", bookResp)
}

func (h *bookingHandler) GetAllHistories(e echo.Context) error {

	userId, _, role := authentication.ExtractTokenUserId(e)
	if role != "user" {
		return response.RespondJSON(e, 401, "can't create booking", nil)
	}

	bookHistories, err := h.bookingSvc.GetAllHistories(userId)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	historiesResp := []response.BookingRespon{}

	for _, v := range bookHistories {
		historyResp := core.BookCoreToBookResp(v)
		historyResp = h.PriceAndNameValues(historyResp, v.DetailsBook)
		historiesResp = append(historiesResp, historyResp)
	}
	return response.RespondJSON(e, 200, "succes", historiesResp)
}

func (h *bookingHandler) GetSpecificHistory(e echo.Context) error {
	userId, _, role := authentication.ExtractTokenUserId(e)
	if role != "user" {
		return response.RespondJSON(e, 401, "can't create booking", nil)
	}

	bookingId := e.Param("id")

	history, err := h.bookingSvc.GetSpecificHistory(bookingId, userId)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), history)
	}

	historyResp := core.BookCoreToBookResp(history)
	historyResp = h.PriceAndNameValues(historyResp, history.DetailsBook)
	return response.RespondJSON(e, 200, "succes", historyResp)
}

func (h *bookingHandler) FindBookingByID(e echo.Context) error {
	_, _, role := authentication.ExtractTokenUserId(e)
	if role != "admin" {
		return response.RespondJSON(e, 401, "can't create booking", nil)
	}

	id := e.Param("id")
	booking, userName, err := h.bookingSvc.FindBookingByID(id)

	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}
	bookResp := core.BookCoreToBookResp(booking)

	bookResp = h.PriceAndNameValues(bookResp, booking.DetailsBook)
	return response.RespJSONWithUserName(e, 200, "succes", bookResp, userName)
}

func (h *bookingHandler) UpdateStatusBooking(e echo.Context) error {
	_, _, role := authentication.ExtractTokenUserId(e)
	if role != "admin" {
		return response.RespondJSON(e, 401, "can't create booking", nil)
	}

	newStatus := request.NewStatusReq{}
	NoInvoice := e.Param("invoice")

	err := e.Bind(&newStatus)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	dataCore := core.UpdateStatusToBookCore(newStatus, NoInvoice)
	updatedStatus, userName, err := h.bookingSvc.UpdateStatusBooking(dataCore)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	respUpdated := core.BookCoreToBookResp(updatedStatus)
	respUpdated = h.PriceAndNameValues(respUpdated, updatedStatus.DetailsBook)
	return response.RespJSONWithUserName(e, 200, "succes", respUpdated, userName)

}

func (h *bookingHandler) GetAllBookings(e echo.Context) error {
	_, _, role := authentication.ExtractTokenUserId(e)
	if role != "admin" {
		return response.RespondJSON(e, 401, "can't create booking", nil)
	}

	bookingsData, err := h.bookingSvc.FindAllBookings()
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	bookingsrespons := []response.BookingRespon{}
	for _, v := range bookingsData {
		bookResp := core.BookCoreToBookResp(v)
		bookingsrespons = append(bookingsrespons, bookResp)
	}
	return response.RespondJSON(e, 200, "succes", bookingsrespons)
}

func (h *bookingHandler) PriceAndNameValues(response response.BookingRespon, data []core.DetailsBookCore) response.BookingRespon {
	for _, v := range data {
		booking := core.DetailsBookCoreToDetailsBookResp(v)
		dataService, _ := h.bookingSvc.FindServiceByID(v.ServiceID)
		booking.Name = dataService.Name
		booking.Price = dataService.Price
		response.DetailsBooking = append(response.DetailsBooking, booking)
	}
	return response
}
