package repositories

import (
	"glamour_reserve/entity/models"

	"gorm.io/gorm"
)

type ServiceRepoInterface interface {
	FindAll() ([]models.Service, error)
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
