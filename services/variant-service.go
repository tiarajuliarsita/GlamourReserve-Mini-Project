package services

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/repositories"
)

type VariantServiceInterface interface {
	Create(variant core.VariantCore) (core.VariantCore, error)
}

type variantService struct {
	variantRepo repositories.VariantRepoInterface
}

func NewVariantService(variantRepo repositories.VariantRepoInterface) *variantService {
	return &variantService{variantRepo}
}

func (s *variantService) Create(variant core.VariantCore) (core.VariantCore, error) {
	variant, err := s.variantRepo.Create(variant)
	if err!=nil{
		return variant, err
	}
	return variant, nil
}
