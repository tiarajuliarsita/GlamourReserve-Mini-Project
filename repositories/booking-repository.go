package repositories

import (
	"errors"
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/models"

	"gorm.io/gorm"
)

type BookingRepoInterface interface {
	Create(booking core.BookingCore) (core.BookingCore, error)
	FindServiceByID(id string) (core.ServiceCore, error)
	GetPriceService(id string) (int, error)
	CheckAvailableService(date, time string)  error
	// GetBookingsUser(userID string)([]core.BookingCore, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *bookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) Create(bookingNew core.BookingCore) (core.BookingCore, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return core.BookingCore{}, tx.Error
	}

	InsertBook := core.BookingCoreToBookingModels(bookingNew)
	dataBook := core.BookingCore{}

	err := tx.Create(&InsertBook).Error
	if err != nil {
		tx.Rollback()
		return dataBook, err
	}

	dataBook = core.BookingModelToBookingCore(InsertBook)
	tx.Commit()

	for _, v := range InsertBook.DetailsBooking {
		detailBook := core.DetailBookingModelToDetailBookingCore(v)
		dataBook.DetailsBook = append(dataBook.DetailsBook, detailBook)
	}

	return dataBook, nil
}

func (r *bookingRepository) FindServiceByID(id string) (core.ServiceCore, error) {

	service := models.Service{}
	dataService := core.ServiceCore{}

	err := r.db.Where("id = ?", id).Find(&service).Error
	if err != nil {
		return dataService, err
	}

	dataService = core.ServiceModelToServiceCore(service)
	return dataService, nil
}

func (r *bookingRepository) GetPriceService(id string) (int, error) {
	service := models.Service{}
	err := r.db.Where("id = ?", id).Find(&service).Error
	if err != nil {

		return service.Price, err
	}
	return service.Price, nil

}

func (r *bookingRepository) CheckAvailableService(date, time string) error {
	var detailBooking []models.DetailBooking
	err := r.db.Where("Date = ? AND Time = ?", date, time).Find(&detailBooking).Error

	if err != nil {
		return err
	}

	if len(detailBooking) > 0 {
		return errors.New("service not available")
	}
	return nil
}

