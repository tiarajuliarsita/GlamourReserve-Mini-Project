package core

import (
	"glamour_reserve/entity/models"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
)

func BookingDataRequestToDetailsBookingCore(book request.BookDataRequest) DetailsBookCore {
	dataDetailBook := DetailsBookCore{
		DateTime:     book.DateTime,
		ServiceID:    book.Service_id,
		TimeExpected: book.TimeExpected,
	}
	return dataDetailBook
}

func DetailsBookCoreToModelDetailsBook(book DetailsBookCore) models.DetailBooking {
	dataDetailBook := models.DetailBooking{
		ID:           book.ID,
		DateTime:     book.DateTime, // Time:      book.Time,
		TimeExpected: book.TimeExpected,
		BookingID:    book.BookingID,
		ServiceID:    book.ServiceID,
	}
	return dataDetailBook
}

func DetailsBookCoreToDetailsBookResp(book DetailsBookCore) response.DetailBookingRespon {
	dataDetailBook := response.DetailBookingRespon{
		ID:           book.ID,
		ServiceID:    book.ServiceID,
		TimeExpected: book.TimeExpected,
		Price:        0,
		Name:         "",
		DateTime:     book.DateTime,
	}
	return dataDetailBook

}

func DetailBookingModelToDetailBookingCore(book models.DetailBooking) DetailsBookCore {
	dataDetailBook := DetailsBookCore{
		ID:           book.ID,
		DateTime:     book.DateTime,
		TimeExpected: book.TimeExpected,
		BookingID:    book.BookingID,
		ServiceID:    book.ServiceID,
		CreatedAt:    book.CreatedAt,
		UpdatedAt:    book.CreatedAt,
	}

	return dataDetailBook
}
