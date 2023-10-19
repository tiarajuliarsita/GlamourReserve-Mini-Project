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
	CheckAvailableService(date, time string) error
	GetAllHistories(userID string) ([]core.BookingCore, error)
	GetSpecificHistory(userID, bookingID string) (core.BookingCore, error)
	FindBookingById(bookingId string) (core.BookingCore, error)
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

func (r *bookingRepository) GetAllHistories(userID string) ([]core.BookingCore, error) {
	bookingsData := []models.Booking{}

	err := r.db.Where("user_id = ?", userID).Preload("DetailsBooking").Find(&bookingsData).Error
	if err != nil {
		return nil, err
	}

	BookingsCore := []core.BookingCore{}
	for _, v := range bookingsData {
		data := core.BookingModelToBookingCore(v)

		for _, details := range v.DetailsBooking {
			detailsCore := core.DetailBookingModelToDetailBookingCore(details)
			data.DetailsBook = append(data.DetailsBook, detailsCore)
		}
		BookingsCore = append(BookingsCore, data)
	}

	return BookingsCore, nil
}

func (r *bookingRepository) GetSpecificHistory(userID, bookingID string) (core.BookingCore, error) {
	bookingData := models.Booking{}

	err := r.db.Where("id = ? AND user_id = ?", bookingID, userID).Preload("DetailsBooking").First(&bookingData).Error
	if err != nil {
		return core.BookingCore{}, err
	}

	data := core.BookingModelToBookingCore(bookingData)
	for _, v := range bookingData.DetailsBooking {
		dataDetails := core.DetailBookingModelToDetailBookingCore(v)
		data.DetailsBook = append(data.DetailsBook, dataDetails)
	}

	return data, nil
}

func (r *bookingRepository) FindBookingById(bookingId string) (core.BookingCore, error) {
	booking := models.Booking{}
	err := r.db.Where("id = ?", bookingId).Preload("DetailsBooking").First(&booking).Error
	if err != nil {
		return core.BookingCore{}, err
	}

	data := core.BookingModelToBookingCore(booking)
	return data, err
}
