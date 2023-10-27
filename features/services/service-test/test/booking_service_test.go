package services

import (
	"encoding/json"
	"glamour_reserve/entity/core"
	"testing"

	"glamour_reserve/features/services"
	"glamour_reserve/features/services/service-test/mock"

	"github.com/stretchr/testify/assert"
	mc "github.com/stretchr/testify/mock"
)

func TestMockBookingRepo_FindUserEmails(t *testing.T) {
	mockRepo := new(mock.MockBookingRepo)

	userID := "1"
	expectedUserEmail := "tiara@gmail.com"

	mockRepo.On("FindUserEmails", userID).Return(expectedUserEmail, nil)
	userEmail, err := mockRepo.FindUserEmails(userID)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedUserEmail, userEmail, "User email is not as expected")
}

func TestMockBookingRepo_FindUserName(t *testing.T) {
	mockRepo := new(mock.MockBookingRepo)

	userID := "1"
	expectedUserName := "tiarajuliarsita"

	mockRepo.On("FindUserName", userID).Return(expectedUserName)
	userName := mockRepo.FindUserName(userID)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, expectedUserName, userName, "User name is not as expected")
}

func TestMockBookingRepo_CheckAvailableService(t *testing.T) {
	mockRepo := new(mock.MockBookingRepo)

	serviceID := "1"
	start := "2023-10-27 10:00:00"
	end := "2023-10-27 11:00:00"

	mockRepo.On("CheckAvailableService", serviceID, start, end).Return(nil)
	err := mockRepo.CheckAvailableService(serviceID, start, end)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestMockBookingRepo_GetPriceService(t *testing.T) {
	mockRepo := new(mock.MockBookingRepo)

	serviceID := "1"
	expectedPrice := 100

	mockRepo.On("GetPriceService", serviceID).Return(expectedPrice, nil)
	price, err := mockRepo.GetPriceService(serviceID)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedPrice, price, "Price is not as expected")
}

func TestBookingService_FindBookingByID(t *testing.T) {
	mockRepo := new(mock.MockBookingRepo)

	bookingID := "1"
	expectedBookingJSON := `
{
    "data": {
        "id": "7602b450-178e-4cac-8536-8b15461e22c8",
        "invoice_num": "2023-10-20474e4c3145435048",
        "total": 3000,
        "details": [
            {
                "id": "034513b0-45e4-46c1-a4c3-72fd1c7efbd6",
                "service_id": "0fdda9a5-1e71-43a9-945c-89c041d3cd9a",
                "name": "service 1",
                "price": 3000,
                "date": "2023-10-2",
                "time": "15:20-16:20"
            }
        ],
        "created_at": "2023-10-20T00:24:35.955+08:00"
    },
    "status": "success",
    "user_name": "tiara"
}
`

	var expectedBooking core.BookingCore
	if err := json.Unmarshal([]byte(expectedBookingJSON), &expectedBooking); err != nil {
		t.Fatalf("Failed to unmarshal expectedBooking: %v", err)
	}

	expectedUserName := "tiara"

	mockRepo.On("FindBookingById", bookingID).Return(expectedBooking, nil)
	mockRepo.On("FindUserName", mc.Anything).Return(expectedUserName)

	bookingService := services.NewBookingService(mockRepo)

	booking, userName, err := bookingService.FindBookingByID(bookingID)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedBooking, booking, "Booking is not as expected")
	assert.Equal(t, expectedUserName, userName, "User name is not as expected")
}

func TestBookingService_FindServiceByID(t *testing.T) {
	mockRepo := new(mock.MockBookingRepo)
	serviceID := "1"
	expectedService := core.ServiceCore{ID: "1", Name: "tiara"}
	mockRepo.On("FindServiceByID", serviceID).Return(expectedService, nil)

	bookingService := services.NewBookingService(mockRepo)
	service, err := bookingService.FindServiceByID(serviceID)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err, "Seharusnya tidak ada error, tetapi ada error: %v", err)
	assert.Equal(t, expectedService, service, "Hasil tidak sesuai dengan ekspektasi")
}

func TestBookingService_GetAllHistories(t *testing.T) {
	MockBookingRepo := new(mock.MockBookingRepo)

	userID := "1"
	expectedHistories := []core.BookingCore{} 
	MockBookingRepo.On("GetAllHistories", userID).Return(expectedHistories, nil)

	bookingService := services.NewBookingService(MockBookingRepo)

	histories, err := bookingService.GetAllHistories(userID)

	MockBookingRepo.AssertExpectations(t)
	assert.NoError(t, err,  err)
	assert.Equal(t, expectedHistories, histories)
}

func TestBookingService_GetSpecificHistory(t *testing.T) {

	MockBookingRepo := new(mock.MockBookingRepo)

	userID := "1"
	bookingID := "2"
	expectedHistory := core.BookingCore{}
	MockBookingRepo.On("GetSpecificHistory", userID, bookingID).Return(expectedHistory, nil)

	bookingService := services.NewBookingService(MockBookingRepo)
	history, err := bookingService.GetSpecificHistory(bookingID, userID)
	MockBookingRepo.AssertExpectations(t)
	assert.NoError(t, err,  err)
	assert.Equal(t, expectedHistory, history)
}

func TestBookingService_FindAllBookings(t *testing.T) {
	mockRepo := new(mock.MockBookingRepo)

	expectedBookings := []core.BookingCore{}

	expectedUserNames := []string{
		"tiara ",
		"yaya",
	}

	mockRepo.On("FindAllBookings").Return(expectedBookings, nil)
	mockRepo.On("FindUserName", mc.Anything).Return(expectedUserNames)

	bookingService := services.NewBookingService(mockRepo)
	_, err := bookingService.FindAllBookings("user")
	assert.NoError(t, err)

}

func TestBookingService_UpdateStatusBooking(t *testing.T) {
	mockRepo := new(mock.MockBookingRepo)

	newDatacore := core.BookingCore{
		InvoiceNumb: "INV123",
	}

	expectedUpdatedStatus := core.BookingCore{
		InvoiceNumb: "INV123",
		Status:      "Updated",
	}
	expectedUserName := "tiara"
	mockRepo.On("FindBookingByInvoice", mc.Anything).Return(expectedUpdatedStatus, nil)
	mockRepo.On("UpdateStatusInovice", mc.Anything, newDatacore).Return(expectedUpdatedStatus, nil)
	mockRepo.On("FindBookingByInvoice", expectedUpdatedStatus.InvoiceNumb).Return(expectedUpdatedStatus, nil)
	mockRepo.On("FindUserName", mc.Anything).Return(expectedUserName)

	bookingService := services.NewBookingService(mockRepo)

	updatedStatus, userName, err := bookingService.UpdateStatusBooking(newDatacore)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedUpdatedStatus, updatedStatus, "Updated status is not as expected")
	assert.Equal(t, expectedUserName, userName, "User name is not as expected")

}

func TestBookingService_Create(t *testing.T) {
	// Create a mock repository
	mockRepo := new(mock.MockBookingRepo)

	// Create a sample booking
	booking := core.BookingCore{
		UserID: "1",
		DetailsBook: []core.DetailsBookCore{
			{
				ServiceID:    "2",
				ServiceStart: "2023-10-27 10:00:00",
				ServiceEnd:   "2023-10-27 11:00:00",
			},
		},
	}

	userName := "tiara"

	
	mockRepo.On("FindServiceByID", mc.Anything).Return(core.ServiceCore{}, nil)
	mockRepo.On("GetPriceService", mc.Anything).Return(0, nil)
	mockRepo.On("CheckAvailableService", mc.Anything, mc.Anything, mc.Anything).Return(nil)
	mockRepo.On("Create", mc.Anything).Return(booking, nil)
	mockRepo.On("FindUserEmails", mc.Anything).Return("user@example.com", nil)

	bookingService := services.NewBookingService(mockRepo)
	createdBook, err := bookingService.Create(booking, userName)
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.NotNil(t, createdBook)
}
