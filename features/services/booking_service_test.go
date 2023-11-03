package services

import (
	"errors"
	"glamour_reserve/entity/core"
	"glamour_reserve/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//create
func TestBookingService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)

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

	const userName = "tiara"

	mockRepo.EXPECT().FindServiceByID(gomock.Any()).Return(core.ServiceCore{}, nil)
	mockRepo.EXPECT().GetPriceService(gomock.Any()).Return(0, nil)
	mockRepo.EXPECT().CheckAvailableService(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockRepo.EXPECT().Create(gomock.Any()).Return(booking, nil)
	mockRepo.EXPECT().FindUserEmails(gomock.Any()).Return("user@example.com", nil)

	bookingService := NewBookingService(mockRepo)
	createdBook, err := bookingService.Create(booking, userName)

	assert.NoError(t, err)
	assert.NotNil(t, createdBook)
}

func TestBookingService_Create_ErrorCases_FindServiceByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
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

	const userName = "tiara"

	// Simulasi pengembalian kesalahan dari FindServiceByID
	expectedServiceError := errors.New("Error finding service by ID")
	mockRepo.EXPECT().FindServiceByID(gomock.Any()).Return(core.ServiceCore{}, expectedServiceError)

	bookingService := NewBookingService(mockRepo)
	_, err := bookingService.Create(booking, userName)
	assert.Error(t, err)
	assert.Equal(t, expectedServiceError, err)
}

func TestBookingService_Create_ErrorCases_GetPriceService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)

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

	const userName = "tiara"

	// Simulasi pengembalian kesalahan dari GetPriceService
	expectedPriceError := errors.New("Error getting service price")
	mockRepo.EXPECT().FindServiceByID(gomock.Any()).Return(core.ServiceCore{}, nil)
	mockRepo.EXPECT().GetPriceService(gomock.Any()).Return(0, expectedPriceError)

	bookingService := NewBookingService(mockRepo)
	_, err := bookingService.Create(booking, userName)
	assert.Error(t, err)
	assert.Equal(t, expectedPriceError, err)
}

func TestBookingService_Create_ErrorCases_CheckAvailableService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
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

	const userName = "tiara"

	// Simulasi pengembalian kesalahan dari CheckAvailableService
	expectedAvailabilityError := errors.New("Error checking service availability")
	mockRepo.EXPECT().FindServiceByID(gomock.Any()).Return(core.ServiceCore{}, nil)
	mockRepo.EXPECT().GetPriceService(gomock.Any()).Return(0, nil)
	mockRepo.EXPECT().CheckAvailableService(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedAvailabilityError)

	bookingService := NewBookingService(mockRepo)
	_, err := bookingService.Create(booking, userName)
	assert.Error(t, err)
	assert.Equal(t, expectedAvailabilityError, err)
}

func TestBookingService_Create_ErrorCases_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)

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

	const userName = "tiara"

	expectedCreateError := errors.New("Error creating booking")
	mockRepo.EXPECT().FindServiceByID(gomock.Any()).Return(core.ServiceCore{}, nil)
	mockRepo.EXPECT().GetPriceService(gomock.Any()).Return(0, nil)
	mockRepo.EXPECT().CheckAvailableService(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockRepo.EXPECT().Create(gomock.Any()).Return(core.BookingCore{}, expectedCreateError)

	bookingService := NewBookingService(mockRepo)
	_, err := bookingService.Create(booking, userName)
	assert.Error(t, err)
	assert.Equal(t, expectedCreateError, err)
}

//find service by id
func TestBookingService_FindServiceByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)

	serviceID := "123"
	expectedServiceData := core.ServiceCore{
		ID:          serviceID,
		Name:        "service 1",
		Image:       "http.//glamour.id",
		Price:       100,
		Description: "the description",
	}

	mockRepo.EXPECT().FindServiceByID(serviceID).Return(expectedServiceData, nil)
	serviceData, err := service.FindServiceByID(serviceID)
	assert.NoError(t, err)
	assert.Equal(t, expectedServiceData, serviceData)
}

func TestBookingService_FindServiceByID_ErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)

	serviceID := "123"
	expectedGetError := errors.New("Error finding service by ID")
	mockRepo.EXPECT().FindServiceByID(serviceID).Return(core.ServiceCore{}, expectedGetError)

	_, err := service.FindServiceByID(serviceID)
	assert.Error(t, err)
	assert.Equal(t, expectedGetError, err)
}

//get all histories
func TestBookingService_GetAllHistories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)

	userID := "user123"
	detailsExpected := []core.DetailsBookCore{
		{
			ID:           "1",
			ServiceID:    "1",
			ServiceStart: "2023-02-07 10:00",
			ServiceEnd:   "2023-02-07 11:00",
		},
	}
	expectedHistories := []core.BookingCore{
		{
			ID:          "1",
			UserID:      userID,
			InvoiceNumb: "123",
			Total:       100,
			Status:      "pending",
			DetailsBook: detailsExpected,
		},
	}

	mockRepo.EXPECT().GetAllHistories(userID).Return(expectedHistories, nil)
	histories, err := service.GetAllHistories(userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedHistories, histories)
}

func TestBookingService_GetAllHistories_ErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)

	userID := "user123"
	expectedGetError := errors.New("Error getting all booking histories")
	mockRepo.EXPECT().GetAllHistories(userID).Return(nil, expectedGetError)

	_, err := service.GetAllHistories(userID)
	assert.Error(t, err)
	assert.Equal(t, expectedGetError, err)
}

//get spesific history
func TestBookingService_GetSpecificHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)

	userID := "user123"
	bookingID := "1"
	detailsExpected := []core.DetailsBookCore{
		{
			ID:           "1",
			ServiceID:    "1",
			ServiceStart: "2023-02-07 10:00",
			ServiceEnd:   "2023-02-07 11:00",
		},
	}

	expectedHistories := core.BookingCore{
		ID:          "1",
		UserID:      userID,
		InvoiceNumb: "123",
		Total:       100,
		Status:      "pending",
		DetailsBook: detailsExpected,
	}
	mockRepo.EXPECT().GetSpecificHistory(userID, bookingID).Return(expectedHistories, nil)
	histories, err := service.GetSpecificHistory(bookingID, userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedHistories, histories)
}

func TestBookingService_GetSpecificHistory_ErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)
	userID := "user123"
	bookingID := "1"

	expectedGetError := errors.New("Error getting specific booking history")
	mockRepo.EXPECT().GetSpecificHistory(userID, bookingID).Return(core.BookingCore{}, expectedGetError)

	_, err := service.GetSpecificHistory(bookingID, userID)
	assert.Error(t, err)
	assert.Equal(t, expectedGetError, err)
}

//find booking by id
func TestBookingService_FindBookingByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)

	bookingID := "1"
	expectedBooking := core.BookingCore{
		ID:          bookingID,
		UserID:      "user123",
		InvoiceNumb: "123",
		Total:       100,
		Status:      "pending",
		DetailsBook: []core.DetailsBookCore{
			{
				ID:           "1",
				ServiceID:    "1",
				ServiceStart: "2023-02-07 10:00",
				ServiceEnd:   "2023-02-07 11:00",
			},
		},
	}
	expectedUserName := "John Doe"

	mockRepo.EXPECT().FindBookingById(bookingID).Return(expectedBooking, nil)
	mockRepo.EXPECT().FindUserName(expectedBooking.UserID).Return(expectedUserName)

	booking, userName, err := service.FindBookingByID(bookingID)
	assert.NoError(t, err)
	assert.Equal(t, expectedBooking, booking)
	assert.Equal(t, expectedUserName, userName)
}

func TestBookingService_FindBookingByID_ErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)

	bookingID := "1"
	expectedFindError := errors.New("Error finding booking by ID")
	mockRepo.EXPECT().FindBookingById(bookingID).Return(core.BookingCore{}, expectedFindError)

	_, _, err := service.FindBookingByID(bookingID)
	assert.Error(t, err)
	assert.Equal(t, expectedFindError, err)
}

//update status booking
func TestBookingService_UpdateStatusBooking(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	newDatacore := core.BookingCore{
		InvoiceNumb: "INV123",
	}

	expectedUpdatedStatus := core.BookingCore{
		InvoiceNumb: "INV123",
		Status:      "Updated",
	}
	expectedUserName := "tiara"

	mockRepo.EXPECT().FindBookingByInvoice(newDatacore.InvoiceNumb).Return(expectedUpdatedStatus, nil)
	mockRepo.EXPECT().UpdateStatusInovice(newDatacore.InvoiceNumb, newDatacore).Return(expectedUpdatedStatus, nil)
	mockRepo.EXPECT().FindBookingByInvoice(expectedUpdatedStatus.InvoiceNumb).Return(expectedUpdatedStatus, nil)
	mockRepo.EXPECT().FindUserName(newDatacore.UserID).Return(expectedUserName)

	bookingService := NewBookingService(mockRepo)
	updatedStatus, userName, err := bookingService.UpdateStatusBooking(newDatacore,"1")

	assert.NoError(t, err)
	assert.Equal(t, expectedUpdatedStatus, updatedStatus)
	assert.Equal(t, expectedUserName, userName)
}

func Test_FindBook_UpdateStatusBooking_ErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	newDatacore := core.BookingCore{
		InvoiceNumb: "INV123",
	}

	expectedError := errors.New("Error finding booking by invoice")
	mockRepo.EXPECT().FindBookingByInvoice(newDatacore.InvoiceNumb).Return(core.BookingCore{}, expectedError)

	bookingService := NewBookingService(mockRepo)
	_, _, err := bookingService.UpdateStatusBooking(newDatacore,"1")
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
func TestBookingService_UpdateStatusBooking_ErrorCases_FindByInvoiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)

	newDatacore := core.BookingCore{
		InvoiceNumb: "INV123",
	}

	expectedError := errors.New("Error finding booking by invoice")
	mockRepo.EXPECT().FindBookingByInvoice(newDatacore.InvoiceNumb).Return(core.BookingCore{}, expectedError)

	bookingService := NewBookingService(mockRepo)
	_, _, err := bookingService.UpdateStatusBooking(newDatacore,"1")
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestBookingService_UpdateStatusBooking_ErrorCases_UpdateStatusError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)

	newDatacore := core.BookingCore{
		InvoiceNumb: "INV123",
	}

	expectedError := errors.New("Error updating status")
	mockRepo.EXPECT().FindBookingByInvoice(newDatacore.InvoiceNumb).Return(core.BookingCore{}, nil)
	mockRepo.EXPECT().UpdateStatusInovice(newDatacore.InvoiceNumb, newDatacore).Return(core.BookingCore{}, expectedError)

	bookingService := NewBookingService(mockRepo)
	_, _, err := bookingService.UpdateStatusBooking(newDatacore,"1")
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

//find all booking
func TestBookingService_FindAllBookings(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)
	userID := "user123"
	bookingData := []core.BookingCore{
		{
			ID:          "1",
			UserID:      userID,
			InvoiceNumb: "123",
			Total:       100,
			Status:      "pending",
			DetailsBook: []core.DetailsBookCore{
				{
					ID:           "1",
					ServiceID:    "1",
					ServiceStart: "2023-02-07 10:00",
					ServiceEnd:   "2023-02-07 11:00",
				},
			},
		},
	}

	mockRepo.EXPECT().FindAllBookings().Return(bookingData, nil)
	for _, booking := range bookingData {
		name := "User Name"
		mockRepo.EXPECT().FindUserName(booking.UserID).Return(name)
	}

	_, err := service.FindAllBookings(userID)
	assert.NoError(t, err)
}

func TestBookingService_FindAllBookings_ErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookingRepoInterface(ctrl)
	service := NewBookingService(mockRepo)

	userID := "user123"
	expectedError := errors.New("Error retrieving bookings")
	mockRepo.EXPECT().FindAllBookings().Return(nil, expectedError)

	_, err := service.FindAllBookings(userID)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
