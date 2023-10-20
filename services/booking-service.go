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
	GetAllHistories(userID string) ([]core.BookingCore, error)
	GetSpecificHistory(bookingId, userId string) (core.BookingCore, error)
	FindBookingByID(bookingId string) (core.BookingCore, string, error)
	UpdateStatusBooking(newDatacore core.BookingCore) (core.BookingCore, string, error)
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

func (s *bookingService) GetAllHistories(userID string) ([]core.BookingCore, error) {
	dataBookings, err := s.bookRepo.GetAllHistories(userID)
	if err != nil {
		return nil, err
	}
	return dataBookings, nil
}

func (s *bookingService) GetSpecificHistory(bookingId, userId string) (core.BookingCore, error) {
	data, err := s.bookRepo.GetSpecificHistory(userId, bookingId)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *bookingService) FindBookingByID(bookingId string) (core.BookingCore, string, error) {
	data, err := s.bookRepo.FindBookingById(bookingId)

	if err != nil {
		return data, "", err
	}
	userName := s.bookRepo.FindUserName(data.UserID)
	return data, userName, nil
}

func (s *bookingService) UpdateStatusBooking(newDatacore core.BookingCore) (core.BookingCore, string, error) {
	_, err := s.bookRepo.FindBookingByInvoice(newDatacore.InvoiceNumb)
	if err != nil {
		if err != nil {
			return core.BookingCore{}, "", err
		}
	}

	data, err := s.bookRepo.UpdateStatusInovice(newDatacore.InvoiceNumb, newDatacore)
	if err != nil {
		return data, "", err
	}

	updatedStatus, err := s.bookRepo.FindBookingByInvoice(data.InvoiceNumb)
	if err != nil {
		return updatedStatus, "", err
	}
	userName := s.bookRepo.FindUserName(updatedStatus.UserID)
	return updatedStatus, userName, nil

}
