package services

import (
	"glamour_reserve/entity/core"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) FindById(id string) (core.ServiceCore, error) {
	args := m.Called(id)
	return args.Get(0).(core.ServiceCore), args.Error(1)
}

func (m *MockRepo) FindAll(name string, offset, limit int) ([]core.ServiceCore, error) {
	args := m.Called(name, offset, limit)
	return args.Get(0).([]core.ServiceCore), args.Error(1)
}

func (m *MockRepo) Create(service core.ServiceCore) (core.ServiceCore, error) {
	args := m.Called(service)
	return args.Get(0).(core.ServiceCore), args.Error(1)
}

func (m *MockRepo) Delete(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepo) Update(id string, updateSvc core.ServiceCore) (core.ServiceCore, error) {
	args := m.Called(id, updateSvc)
	return args.Get(0).(core.ServiceCore), args.Error(1)
}

func TestSvcService_FindById(t *testing.T) {
	mockRepo := new(MockRepo)
	expectedService := core.ServiceCore{
		ID:   "1",
		Name: "manny paddy",
	}

	mockRepo.On("FindById", "1").Return(expectedService, nil)
	svcService := NewSvcService(mockRepo)

	service, err := svcService.FindById("1")

	assert.NoError(t, err, "Seharusnya tidak ada error")
	assert.Equal(t, expectedService, service, "Hasil tidak sesuai dengan ekspektasi")
	mockRepo.AssertExpectations(t)
}

func TestSvcService_FindAll(t *testing.T) {
	mockRepo := new(MockRepo)
	expectedServices := []core.ServiceCore{
		{ID: "1", Name: "Service 1"},
		{ID: "2", Name: "Service 2"},
	}
	mockRepo.On("FindAll", "", 0, 0).Return(expectedServices, nil)

	svcService := NewSvcService(mockRepo)

	services, err := svcService.FindAll("", "0", "0")

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err, "Seharusnya tidak ada error")
	assert.Equal(t, len(expectedServices), len(services), "Jumlah hasil sesuai dengan ekspektasi")
	assert.ElementsMatch(t, expectedServices, services, "Hasil sesuai dengan ekspektasi")
}

func TestSvcService_CreateService(t *testing.T) {

	mockRepo := new(MockRepo)
	serviceInput := core.ServiceCore{ID: "1", Name: "New Service"}
	mockRepo.On("Create", serviceInput).Return(serviceInput, nil)

	svcService := NewSvcService(mockRepo)

	createdService, err := svcService.CreateService(serviceInput)

	assert.NoError(t, err)
	assert.Equal(t, serviceInput, createdService, "Hasil harus sesuai dengan ekspektasi")
	mockRepo.AssertExpectations(t)
}

func TestSvcService_Delete(t *testing.T) {
	mockRepo := new(MockRepo)
	serviceID := "1"
	mockRepo.On("Delete", serviceID).Return(true, nil)

	svcService := NewSvcService(mockRepo)

	deleted, err := svcService.Delete(serviceID)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err, "Seharusnya tidak ada error, tetapi ada error: %v", err)
	assert.True(t, deleted, "Seharusnya hasil adalah true, tetapi got: false")
}

func TestSvcService_Update(t *testing.T) {
	mockRepo := new(MockRepo)
	serviceID := "1"
	newService := core.ServiceCore{ID: "1", Name: "Updated Service"}
	mockRepo.On("Update", serviceID, newService).Return(newService, nil)

	svcService := NewSvcService(mockRepo)

	updatedService, err := svcService.Update(serviceID, newService)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err, "Seharusnya tidak ada error, tetapi ada error: %v", err)
	assert.Equal(t, newService, updatedService, "Hasil tidak sesuai dengan ekspektasi")
}
