package repositories

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/models"

	"gorm.io/gorm"
)

type ServiceRepoInterface interface {
	FindAll() ([]models.Service, error)
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

func (r *SvcRepository) FindAll() ([]models.Service, error) {
	var services []models.Service

	err := r.db.Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
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
