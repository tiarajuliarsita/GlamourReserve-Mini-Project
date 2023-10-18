package services

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/helpers"
	"glamour_reserve/repositories"
	"time"
)

type BookingServiceInterface interface {
	Create(booking core.BookingCore) (core.BookingCore, error)
	FindServiceByID(id string) (core.ServiceCore, error)
}

type bookingService struct {
	bookRepo repositories.BookingRepoInterface
}

func NewBookingService(bookRepo repositories.BookingRepoInterface) *bookingService {
	return &bookingService{bookRepo}
}

func (s *bookingService) Create(booking core.BookingCore) (core.BookingCore, error) {
	listPrice := []int{}

	for _, v := range booking.DetailsBook {
		price, err := s.bookRepo.GetPriceService(v.ServiceID)
		if err != nil {
			return core.BookingCore{}, err
		}
		listPrice = append(listPrice, price)
	}

	for _, v := range booking.DetailsBook {
		err := s.bookRepo.CheckAvailableService(v.Date, v.Time)
		if err != nil {
			return core.BookingCore{}, err
		}
	}

	total := helpers.SumTotal(listPrice)
	invoice := helpers.CreateInvoice(time.Now())

	booking.InvoiceNumb = invoice
	booking.Total = total

	dataBook, err := s.bookRepo.Create(booking)
	if err != nil {
		return dataBook, err
	}

	return dataBook, nil
}

func (s *bookingService) FindServiceByID(id string) (core.ServiceCore, error) {
	dataService, err := s.bookRepo.FindServiceByID(id)
	if err != nil {
		return dataService, err
	}
	return dataService, nil
}
