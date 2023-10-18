package repositories

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/models"
	"glamour_reserve/helpers"

	"gorm.io/gorm"
)

type BookingRepoInterface interface {
	Create(booking core.BookingCore) (core.BookingCore, error)
	FindServiceByID(id string) (core.ServiceCore, error)
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
	listPrice := []int{}
	service := models.Service{}

	for _, v := range InsertBook.DetailsBooking {
		err := r.db.Where("id = ?", v.ServiceID).First(&service).Error
		if err != nil {
			tx.Rollback()
			return dataBook, err
		}
		listPrice = append(listPrice, service.Price)
	}

	total := helpers.SumTotal(listPrice)
	InsertBook.Total = total

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
