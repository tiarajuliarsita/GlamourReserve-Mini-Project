package handlers

import (
	core "glamour_reserve/entity/core"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
	"glamour_reserve/utils/helpers/authentication"
	"glamour_reserve/utils/helpers/constanta"
	"net/http"

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
	if role != constanta.USER {
		return response.RespondJSON(e, http.StatusUnauthorized, "you don't have permission", nil)
	}

	NewBooking := request.BookingRequest{}
	err := e.Bind(&NewBooking)
	if err != nil {
		return response.RespondJSON(e, http.StatusBadRequest, err.Error(), nil)
	}

	NewBookingsCore := []core.DetailsBookCore{}
	bookingInsert := core.BookingReqMap(NewBooking, NewBookingsCore, userId)

	createdBook, err := h.bookingSvc.Create(bookingInsert, userName)
	if err != nil {
		return response.RespondJSON(e, http.StatusInternalServerError, err.Error(), nil)
	}

	bookResp := core.BookCoreToBookResp(createdBook)
	bookResp.Name = userName

	bookResp = h.PriceAndNameValues(bookResp, createdBook.DetailsBook)
	if err != nil {
		return response.RespondJSON(e, http.StatusNotFound, err.Error(), nil)
	}

	return response.RespondJSON(e, http.StatusCreated, "succes", bookResp)
}

func (h *bookingHandler) GetAllHistories(e echo.Context) error {

	userId, _, role := authentication.ExtractTokenUserId(e)
	if role != constanta.USER {
		return response.RespondJSON(e, http.StatusUnauthorized, "you don't have permission", nil)
	}

	bookHistories, err := h.bookingSvc.GetAllHistories(userId)
	if err != nil {
		return response.RespondJSON(e, http.StatusNotFound, err.Error(), nil)
	}

	historiesResp := []response.BookingRespon{}

	for _, v := range bookHistories {
		historyResp := core.BookCoreToBookResp(v)
		historyResp = h.PriceAndNameValues(historyResp, v.DetailsBook)
		historiesResp = append(historiesResp, historyResp)
	}
	return response.RespondJSON(e, http.StatusOK, "succes", historiesResp)
}

func (h *bookingHandler) GetSpecificHistory(e echo.Context) error {
	userId, _, role := authentication.ExtractTokenUserId(e)
	if role != constanta.USER {
		return response.RespondJSON(e, http.StatusUnauthorized, "you don't have permission", nil)
	}

	bookingId := e.Param("id")

	history, err := h.bookingSvc.GetSpecificHistory(bookingId, userId)
	if err != nil {
		return response.RespondJSON(e, http.StatusNotFound, err.Error(), nil)
	}

	historyResp := core.BookCoreToBookResp(history)
	historyResp = h.PriceAndNameValues(historyResp, history.DetailsBook)
	return response.RespondJSON(e, http.StatusOK, "succes", historyResp)
}

func (h *bookingHandler) FindBookingByID(e echo.Context) error {
	_, _, role := authentication.ExtractTokenUserId(e)
	if role != constanta.ADMIN {
		return response.RespondJSON(e, http.StatusForbidden, "you don't have permission", nil)
	}

	id := e.Param("id")
	booking, userName, err := h.bookingSvc.FindBookingByID(id)
	if err != nil {
		return response.RespondJSON(e, http.StatusNotFound, err.Error(), nil)
	}
	bookResp := core.BookCoreToBookResp(booking)

	bookResp = h.PriceAndNameValues(bookResp, booking.DetailsBook)
	return response.RespJSONWithUserName(e, http.StatusOK, "succes", bookResp, userName)
}

func (h *bookingHandler) UpdateStatusBooking(e echo.Context) error {
	_, _, role := authentication.ExtractTokenUserId(e)
	if role != constanta.ADMIN {
		return response.RespondJSON(e, http.StatusForbidden, "you don't have permission", nil)
	}

	newStatus := request.NewStatusReq{}
	booking_id := e.Param("id")

	err := e.Bind(&newStatus)
	if err != nil {
		return response.RespondJSON(e, http.StatusBadRequest, err.Error(), nil)
	}

	dataCore := core.UpdateStatusToBookCore(newStatus)
	updatedStatus, userName, err := h.bookingSvc.UpdateStatusBooking(dataCore,booking_id)
	if err != nil {
		return response.RespondJSON(e, http.StatusInternalServerError, err.Error(), nil)
	}

	respUpdated := core.BookCoreToBookResp(updatedStatus)
	respUpdated = h.PriceAndNameValues(respUpdated, updatedStatus.DetailsBook)
	return response.RespJSONWithUserName(e, http.StatusOK, "succes", respUpdated, userName)

}

func (h *bookingHandler) GetAllBookings(e echo.Context) error {
	_, _, role := authentication.ExtractTokenUserId(e)
	if role != constanta.ADMIN {
		return response.RespondJSON(e, http.StatusForbidden, "you don`t have permission", nil)
	}

	user := e.QueryParam("users")

	bookingsData, err := h.bookingSvc.FindAllBookings(user)
	if err != nil {
		return response.RespondJSON(e, http.StatusNotFound, err.Error(), nil)
	}

	return response.RespondJSON(e, http.StatusOK, "succes", bookingsData)
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
