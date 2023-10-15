package repositories

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/models"

	"gorm.io/gorm"
)

type ServiceRepoInterface interface {
	FindAll() ([]core.ServiceCore, error)
	FindById(id string) (core.ServiceCore, error)
	Create(service core.ServiceCore) (core.ServiceCore, error)
}
type SvcRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *SvcRepository {
	return &SvcRepository{db}
}

func (r *SvcRepository) FindById(id string) (core.ServiceCore, error) {

	service := models.Service{}
	dataService := core.ServiceCore{}

	err := r.db.Where("id = ?", id).First(&service).Error
	if err != nil {
		return dataService, err
	}

	dataService = core.ServiceModelToServiceCore(service)
	return dataService, nil
}

func (r *SvcRepository) FindAll() ([]core.ServiceCore, error) {
	var services []models.Service
	var dataServices []core.ServiceCore

	err := r.db.Find(&services).Error
	if err != nil {
		return nil, err
	}
	for _, v := range services {
		svcCore := core.ServiceModelToServiceCore(v)
		dataServices = append(dataServices, svcCore)
	}

	return dataServices, nil
}

func (r *SvcRepository) Create(service core.ServiceCore) (core.ServiceCore, error) {

	serviceInput := core.ServiceCoreToModelsSevice(service)
	err := r.db.Create(&serviceInput).Error
	if err != nil {
		return service, err
	}

	result := core.ServiceModelToServiceCore(serviceInput)
	return result, nil
}
