package services

import (
	"glamour_reserve/entity/models"
	"glamour_reserve/repositories"
)

type SvcServiceInterface interface {
	FindAll() ([]models.Service, error)
}

type svcService struct {
	svcRepo repositories.ServiceRepoInterface
}



func NewSvcService(svcRepo repositories.ServiceRepoInterface) *svcService {
	return &svcService{svcRepo}
}

func (s *svcService) FindAll() ([]models.Service, error) {
	services, err := s.svcRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return services, nil
}
