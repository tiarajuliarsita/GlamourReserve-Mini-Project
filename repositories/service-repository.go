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
	Delete(id string) (bool, error)
	Update(id string, updateSvc core.ServiceCore) (core.ServiceCore, error)
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

func (r *SvcRepository) Delete(id string) (bool, error) {
	service, err := r.FindById(id)

	dataSvc := core.ServiceCoreToModelsSevice(service)
	if err != nil {
		return false, err
	}

	err = r.db.Where("id = ?", id).Delete(&dataSvc).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *SvcRepository) Update(id string, updateSvc core.ServiceCore) (core.ServiceCore, error) {

	service := core.ServiceCoreToModelsSevice(updateSvc)

	data, err := r.FindById(id)
	if err != nil {
		return data, err
	}

	err = r.db.Where("id = ?", id).Updates(&service).Error
	if err != nil {
		return data, err
	}

	data.ID = id
	return data, nil
}
