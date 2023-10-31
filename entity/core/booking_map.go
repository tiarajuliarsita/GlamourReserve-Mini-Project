package core

import (
	"glamour_reserve/entity/models"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
)

func BookingCoreToBookingModels(book BookingCore) models.Booking {
	dataBooking := models.Booking{
		ID:          book.ID,
		UserID:      book.UserID,
		InvoiceNumb: book.InvoiceNumb,
		Status:      book.Status,
		Total:       book.Total,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}

	for _, v := range book.DetailsBook {
		data := DetailsBookCoreToModelDetailsBook(v)
		dataBooking.DetailsBooking = append(dataBooking.DetailsBooking, data)
	}
	return dataBooking
}

func BookingModelToBookingCore(book models.Booking) BookingCore {
	dataBooking := BookingCore{
		ID:          book.ID,
		UserID:      book.UserID,
		InvoiceNumb: book.InvoiceNumb,
		Status:      book.Status,
		Total:       book.Total,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}

	for _, v := range book.DetailsBooking {
		detailBook := DetailBookingModelToDetailBookingCore(v)
		dataBooking.DetailsBook = append(dataBooking.DetailsBook, detailBook)
	}
	return dataBooking
}

func BookCoreToBookResp(book BookingCore) response.BookingRespon {
	dataBooking := response.BookingRespon{
		ID:          book.ID,
		Name:        "",
		Status:      book.Status,
		InvoiceNumb: book.InvoiceNumb,
		Total:       book.Total,
		CreatedAt:   book.CreatedAt,
	}

	return dataBooking

}

func UpdateStatusToBookCore(newStatus request.NewStatusReq, invoice string) BookingCore {
	dataCore := BookingCore{
		InvoiceNumb: invoice,
		Status:      newStatus.Status,
	}
	return dataCore
}

func BookingReqMap(req request.BookingRequest, detailBookings []DetailsBookCore, userId string) BookingCore {
	for _, detail := range req.Details {
		detailCore := BookingDataRequestToDetailsBookingCore(detail)
		detailBookings = append(detailBookings, detailCore)
	}

	bookingInsert := BookingCore{}
	bookingInsert.DetailsBook = detailBookings
	bookingInsert.UserID = userId

	return bookingInsert

}

func BookingCoreToBookingAll(book BookingCore) BookingAll {
	allBooking := BookingAll{
		Name:        "",
		BookingId:   book.ID,
		InvoiceNumb: book.InvoiceNumb,
		Total:       book.Total,
		Status:      book.Status,
		Created_at:  book.CreatedAt,
	}
	return allBooking
}
