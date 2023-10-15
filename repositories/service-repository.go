package repositories

import (
	"glamour_reserve/entity/models"

	"gorm.io/gorm"
)

type ServiceRepoInterface interface {
	FindAll() ([]models.Service, error)
	Create(service *models.Service) (models.Service, error)
}
type SvcRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *SvcRepository {
	return &SvcRepository{db}
}

func (r *SvcRepository) FindAll() ([]models.Service, error) {
	var services []models.Service
	err := r.db.Find(&services).Error
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (r *SvcRepository) Create(service *models.Service) (models.Service, error) {

	err := r.db.Create(&service).Error
	if err != nil {
		return *service, err
	}
	return *service, nil
}
