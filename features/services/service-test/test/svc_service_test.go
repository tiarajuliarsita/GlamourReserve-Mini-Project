package services

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/features/services"
	"glamour_reserve/features/services/service-test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSvcService_FindById(t *testing.T) {
	MockServiceRepo := new(mock.MockServiceRepo)
	expectedService := core.ServiceCore{
		ID:   "1",
		Name: "manny paddy",
	}

	MockServiceRepo.On("FindById", "1").Return(expectedService, nil)
	svcService := services.NewSvcService(MockServiceRepo)

	service, err := svcService.FindById("1")

	assert.NoError(t, err, "Seharusnya tidak ada error")
	assert.Equal(t, expectedService, service, "Hasil tidak sesuai dengan ekspektasi")
	MockServiceRepo.AssertExpectations(t)
}

func TestSvcService_FindAll(t *testing.T) {
	MockServiceRepo := new(mock.MockServiceRepo)
	expectedServices := []core.ServiceCore{
		{ID: "1", Name: "Service 1"},
		{ID: "2", Name: "Service 2"},
	}
	MockServiceRepo.On("FindAll", "", 0, 0).Return(expectedServices, nil)

	svcService := services.NewSvcService(MockServiceRepo)

	services, err := svcService.FindAll("", "0", "0")

	MockServiceRepo.AssertExpectations(t)
	assert.Nil(t, err, "Seharusnya tidak ada error")
	assert.Equal(t, len(expectedServices), len(services), "Jumlah hasil sesuai dengan ekspektasi")
	assert.ElementsMatch(t, expectedServices, services, "Hasil sesuai dengan ekspektasi")
}

func TestSvcService_CreateService(t *testing.T) {

	MockServiceRepo := new(mock.MockServiceRepo)
	serviceInput := core.ServiceCore{ID: "1", Name: "New Service"}
	MockServiceRepo.On("Create", serviceInput).Return(serviceInput, nil)

	svcService := services.NewSvcService(MockServiceRepo)

	createdService, err := svcService.CreateService(serviceInput)

	assert.NoError(t, err)
	assert.Equal(t, serviceInput, createdService, "Hasil harus sesuai dengan ekspektasi")
	MockServiceRepo.AssertExpectations(t)
}

func TestSvcService_Delete(t *testing.T) {
	MockServiceRepo := new(mock.MockServiceRepo)
	serviceID := "1"
	MockServiceRepo.On("Delete", serviceID).Return(true, nil)

	svcService := services.NewSvcService(MockServiceRepo)

	deleted, err := svcService.Delete(serviceID)

	MockServiceRepo.AssertExpectations(t)
	assert.NoError(t, err, "Seharusnya tidak ada error, tetapi ada error: %v", err)
	assert.True(t, deleted, "Seharusnya hasil adalah true, tetapi got: false")
}

func TestSvcService_Update(t *testing.T) {
	MockServiceRepo := new(mock.MockServiceRepo)
	serviceID := "1"
	newService := core.ServiceCore{ID: "1", Name: "Updated Service"}
	MockServiceRepo.On("Update", serviceID, newService).Return(newService, nil)

	svcService := services.NewSvcService(MockServiceRepo)

	updatedService, err := svcService.Update(serviceID, newService)

	MockServiceRepo.AssertExpectations(t)
	assert.NoError(t, err, "Seharusnya tidak ada error, tetapi ada error: %v", err)
	assert.Equal(t, newService, updatedService, "Hasil tidak sesuai dengan ekspektasi")
}
