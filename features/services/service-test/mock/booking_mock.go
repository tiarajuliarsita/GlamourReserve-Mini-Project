package mock

import (
	"glamour_reserve/entity/core"

	"github.com/stretchr/testify/mock"
)

type MockBookingRepo struct {
	mock.Mock
}

func (m *MockBookingRepo) Create(booking core.BookingCore) (core.BookingCore, error) {
	args := m.Called(booking)
	return args.Get(0).(core.BookingCore), args.Error(1)
}

//sudah
func (m *MockBookingRepo) GetPriceService(idService string) (int, error) {
	args := m.Called(idService)
	return args.Int(0), args.Error(1)
}

//sudah
func (m *MockBookingRepo) CheckAvailableService(serviceID string, start string, end string) error {
	args := m.Called(serviceID, start, end)
	return args.Error(0)
}

//sudah
func (m *MockBookingRepo) FindUserName(userID string) string {
	args := m.Called(userID)
	return args.String(0)
}

//sudah
func (m *MockBookingRepo) FindUserEmails(userID string) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

func (m *MockBookingRepo) FindBookingByInvoice(invoiceNumb string) (core.BookingCore, error) {
	args := m.Called(invoiceNumb)
	return args.Get(0).(core.BookingCore), args.Error(1)
}

//sudah
func (m *MockBookingRepo) FindBookingById(bookingId string) (core.BookingCore, error) {
	args := m.Called(bookingId)
	return args.Get(0).(core.BookingCore), args.Error(1)
}

//sudah
func (m *MockBookingRepo) FindServiceByID(idService string) (core.ServiceCore, error) {
	args := m.Called(idService)
	return args.Get(0).(core.ServiceCore), args.Error(1)
}

//sudah
func (m *MockBookingRepo) UpdateStatusInovice(invoiceNumb string, newDatacore core.BookingCore) (core.BookingCore, error) {
	args := m.Called(invoiceNumb, newDatacore)
	return args.Get(0).(core.BookingCore), args.Error(1)
}

//sudah
func (m *MockBookingRepo) GetAllHistories(userID string) ([]core.BookingCore, error) {
	args := m.Called(userID)
	return args.Get(0).([]core.BookingCore), args.Error(1)
}

//sudah
func (m *MockBookingRepo) GetSpecificHistory(userId, bookingId string) (core.BookingCore, error) {
	args := m.Called(userId, bookingId)
	return args.Get(0).(core.BookingCore), args.Error(1)
}

//sudah
func (m *MockBookingRepo) FindAllBookings() ([]core.BookingCore, error) {
	args := m.Called()
	return args.Get(0).([]core.BookingCore), args.Error(1)
}
