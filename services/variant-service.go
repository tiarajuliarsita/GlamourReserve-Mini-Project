package services

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/repositories"
)

type VariantServiceInterface interface {
	Create(variant core.VariantCore) (core.VariantCore, error)
	FindByID(id string) (core.VariantCore, error)
	FindAll() ([]core.VariantCore, error)
	Delete(id string) error
	Update(id string, newVariant core.VariantCore) (core.VariantCore, error)
}

type variantService struct {
	variantRepo repositories.VariantRepoInterface
}

func NewVariantService(variantRepo repositories.VariantRepoInterface) *variantService {
	return &variantService{variantRepo}
}

func (s *variantService) Create(variant core.VariantCore) (core.VariantCore, error) {
	variant, err := s.variantRepo.Create(variant)
	if err != nil {
		return variant, err
	}
	return variant, nil
}

func (s *variantService) FindByID(id string) (core.VariantCore, error) {
	dataVariant, err := s.variantRepo.FindByID(id)
	if err != nil {
		return dataVariant, err
	}

	return dataVariant, nil
}

func (s *variantService) FindAll() ([]core.VariantCore, error) {
	variants, err := s.variantRepo.FindAll()
	if err != nil {
		return variants, err
	}
	return variants, nil

}

func (s *variantService) Delete(id string) error {
	_, err := s.variantRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.variantRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *variantService) Update(id string, newVariant core.VariantCore) (core.VariantCore, error) {
	variant, err := s.variantRepo.Update(id, newVariant)
	if err != nil {
		return variant, err
	}
	return variant, nil
}
