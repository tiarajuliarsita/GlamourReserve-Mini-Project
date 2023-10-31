package core

import (
	"glamour_reserve/entity/models"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
)

func BookingDataRequestToDetailsBookingCore(book request.BookDataRequest) DetailsBookCore {
	dataDetailBook := DetailsBookCore{
		ServiceID:    book.Service_id,
		ServiceStart: book.ServiceStart,
		ServiceEnd:   book.ServiceEnd,
	}
	return dataDetailBook
}

func DetailsBookCoreToModelDetailsBook(book DetailsBookCore) models.DetailBooking {
	dataDetailBook := models.DetailBooking{
		ID:           book.ID,
		ServiceStart: book.ServiceStart,
		ServiceEnd:   book.ServiceEnd,
		BookingID:    book.BookingID,
		ServiceID:    book.ServiceID,
	}
	return dataDetailBook
}

func DetailsBookCoreToDetailsBookResp(book DetailsBookCore) response.DetailBookingRespon {
	dataDetailBook := response.DetailBookingRespon{
		ID:           book.ID,
		ServiceID:    book.ServiceID,
		Price:        0,
		Name:         "",
		ServiceStart: book.ServiceStart,
		ServiceEnd:   book.ServiceEnd,
	}
	return dataDetailBook

}

func DetailBookingModelToDetailBookingCore(book models.DetailBooking) DetailsBookCore {
	dataDetailBook := DetailsBookCore{
		ID:           book.ID,
		ServiceStart: book.ServiceStart,
		ServiceEnd:   book.ServiceEnd,
		BookingID:    book.BookingID,
		ServiceID:    book.ServiceID,
		CreatedAt:    book.CreatedAt,
		UpdatedAt:    book.CreatedAt,
	}

	return dataDetailBook
}
