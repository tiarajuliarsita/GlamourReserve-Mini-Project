package services

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/repositories"
)

type SvcServiceInterface interface {
	FindAll() ([]core.ServiceCore, error)
	FindById(id string) (core.ServiceCore, error)
	CreateService(service core.ServiceCore) (core.ServiceCore, error)
	Delete(id string) (bool, error)
}

type svcService struct {
	svcRepo repositories.ServiceRepoInterface
}

func NewSvcService(svcRepo repositories.ServiceRepoInterface) *svcService {
	return &svcService{svcRepo}
}

func (s *svcService) FindById(id string) (core.ServiceCore, error) {
	dataService, err := s.svcRepo.FindById(id)
	if err != nil {
		return dataService, err
	}

	return dataService, nil
}

func (s *svcService) FindAll() ([]core.ServiceCore, error) {
	services, err := s.svcRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (s *svcService) CreateService(service core.ServiceCore) (core.ServiceCore, error) {
	serviceData, err := s.svcRepo.Create(service)
	if err != nil {
		return service, err
	}
	return serviceData, nil
}

func (s *svcService) Delete(id string) (bool, error) {
	ok, err := s.svcRepo.Delete(id)
	if err != nil {
		return false, err
	}
	return ok, nil
}
