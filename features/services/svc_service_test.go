package services

import (
	"errors"
	"glamour_reserve/entity/core"
	"glamour_reserve/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSvcService_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	serviceID := "123"
	expectedService := core.ServiceCore{
		ID:   serviceID,
		Name: "Nail Art",
	}

	mockRepo.EXPECT().FindById(serviceID).Return(expectedService, nil)
	foundService, err := service.FindById(serviceID)
	assert.NoError(t, err)
	assert.Equal(t, expectedService, foundService)
}

func TestSvcService_FindById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	serviceID := "123"
	expectedErr := errors.New("service not found")

	mockRepo.EXPECT().FindById(serviceID).Return(core.ServiceCore{}, expectedErr)
	foundService, err := service.FindById(serviceID)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, core.ServiceCore{}, foundService)
}

func TestSvcService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	name := "Service"
	offset := "0"
	limit := "10"
	expectedServices := []core.ServiceCore{
		{ID: "1", Name: "Service1"},
		{ID: "2", Name: "Service2"},
	}

	mockRepo.EXPECT().FindAll(name, 0, 10).Return(expectedServices, nil)
	services, err := service.FindAll(name, offset, limit)
	assert.NoError(t, err)
	assert.Equal(t, expectedServices, services)
}



func TestSvcService_FindAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	name := "Service"
	offset := "0"
	limit := "10"
	expectedErr := errors.New("error fetching services")

	mockRepo.EXPECT().FindAll(name, 0, 10).Return([]core.ServiceCore{}, expectedErr)
	services, err := service.FindAll(name, offset, limit)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, services)
}

func TestSvcService_CreateService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	newService := core.ServiceCore{
		Name: "Nail Art",
	}

	expectedService := core.ServiceCore{
		ID:   "123",
		Name: newService.Name,
	}

	mockRepo.EXPECT().Create(newService).Return(expectedService, nil)
	createdService, err := service.CreateService(newService)
	assert.NoError(t, err)
	assert.Equal(t, expectedService, createdService)
}


func TestSvcService_CreateService_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	newService := core.ServiceCore{
		Name: "Nail Art",
	}
	expectedErr := errors.New("error creating service")

	mockRepo.EXPECT().Create(newService).Return(core.ServiceCore{}, expectedErr)
	_, err := service.CreateService(newService)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}


func TestSvcService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	serviceID := "123"
	deleteResult := true

	mockRepo.EXPECT().Delete(serviceID).Return(deleteResult, nil)
	ok, err := service.Delete(serviceID)

	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestSvcService_Delete_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	serviceID := "123"
	expectedErr := errors.New("error deleting service")
	deleteResult := false

	mockRepo.EXPECT().Delete(serviceID).Return(deleteResult, expectedErr)
	ok, err := service.Delete(serviceID)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.False(t, ok)
}


func TestSvcService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	serviceID := "123"
	updatedService := core.ServiceCore{
		ID:   serviceID,
		Name: "Nail Art",
	}

	mockRepo.EXPECT().Update(serviceID, updatedService).Return(updatedService, nil)
	updatedServiceResult, err := service.Update(serviceID, updatedService)
	assert.NoError(t, err)
	assert.Equal(t, updatedService, updatedServiceResult)
}

func TestSvcService_Update_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockServiceRepoInterface(ctrl)
	service := NewSvcService(mockRepo)

	serviceID := "123"
	updatedService := core.ServiceCore{
		ID:   serviceID,
		Name: "Nail Art",
	}
	expectedErr := errors.New("error updating service")

	mockRepo.EXPECT().Update(serviceID, updatedService).Return(core.ServiceCore{}, expectedErr)
	updatedServiceResult, err := service.Update(serviceID, updatedService)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, core.ServiceCore{}, updatedServiceResult)
}

