package repositories

import (
	"errors"
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/models"
	"time"

	"gorm.io/gorm"
)

type BookingRepoInterface interface {
	Create(booking core.BookingCore) (core.BookingCore, error)
	FindServiceByID(id string) (core.ServiceCore, error)
	GetPriceService(id string) (int, error)
	CheckAvailableService(dateTime time.Time) error
	GetAllHistories(userID string) ([]core.BookingCore, error)
	GetSpecificHistory(userID, bookingID string) (core.BookingCore, error)
	FindBookingById(bookingId string) (core.BookingCore, error)
	UpdateStatusInovice(invoiceNum string, newBook core.BookingCore) (core.BookingCore, error)
	FindBookingByInvoice(invoiceNum string) (core.BookingCore, error)
	FindAllBookings() ([]core.BookingCore, error)
	FindUserName(userId string) string
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

func (r *bookingRepository) CheckAvailableService(dateTime time.Time) error {
	var detailBooking []models.DetailBooking
	err := r.db.Where("date_time = ? ", dateTime).Find(&detailBooking).Error

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
	return data, nil
}

func (r *bookingRepository) FindBookingById(bookingId string) (core.BookingCore, error) {
	booking := models.Booking{}
	err := r.db.Where("id = ?", bookingId).Preload("DetailsBooking").First(&booking).Error

	if err != nil {
		return core.BookingCore{}, err
	}

	data := core.BookingModelToBookingCore(booking)
	return data, nil
}

func (r *bookingRepository) UpdateStatusInovice(invoiceNum string, newBook core.BookingCore) (core.BookingCore, error) {
	booking := core.BookingCoreToBookingModels(newBook)

	err := r.db.Where("invoice_numb = ?", invoiceNum).Updates(&booking).Error
	if err != nil {
		return core.BookingCore{}, err
	}
	data := core.BookingModelToBookingCore(booking)
	return data, nil
}

func (r *bookingRepository) FindBookingByInvoice(invoiceNum string) (core.BookingCore, error) {
	data := models.Booking{}
	err := r.db.Where("invoice_numb = ?", invoiceNum).Preload("DetailsBooking").First(&data).Error
	if err != nil {
		if err != nil {
			return core.BookingCore{}, err
		}
	}

	dataBooking := core.BookingModelToBookingCore(data)
	return dataBooking, nil
}

func (r *bookingRepository) FindUserName(userId string) string {
	user := models.User{}
	r.db.Where("id = ?", userId).First(&user)
	return user.UserName
}

func (r *bookingRepository) FindAllBookings() ([]core.BookingCore, error) {
	data := []models.Booking{}
	err := r.db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	dataResp := []core.BookingCore{}
	for _, v := range data {
		dataCore := core.BookingModelToBookingCore(v)
		dataResp = append(dataResp, dataCore)
	}
	return dataResp, nil
}
