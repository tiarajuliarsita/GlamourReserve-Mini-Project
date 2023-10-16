package core

import (
	"glamour_reserve/entity/models"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
)

func VariantCoreToVariantModel(variant VariantCore) models.Variant {
	variantModel := models.Variant{
		ID:          variant.ID,
		Name:        variant.Name,
		Description: variant.Description,
		Price:       variant.Price,
		ServiceID:   variant.ServiceID,
		CreatedAt:   variant.CreatedAt,
		UpdatedAt:   variant.UpdatedAt,
	}
	return variantModel
}

func VariantModelToVariantCore(variant models.Variant) VariantCore {
	variantCore := VariantCore{
		ID:          variant.ID,
		Name:        variant.Name,
		Description: variant.Description,
		Price:       variant.Price,
		ServiceID:   variant.ServiceID,
		CreatedAt:   variant.CreatedAt,
		UpdatedAt:   variant.UpdatedAt,
	}
	return variantCore

}

func VariantRequestToVariantCore(variant request.VariantRequest)VariantCore{
	variantCore:= VariantCore{
		Name:        variant.Name,
		Description: variant.Description,
		Price:       variant.Price,
		ServiceID:   variant.ServiceID,
	}
	return variantCore
}

func VariantCoreToVariantRespon( variant VariantCore)response.VariantRespon{
	variantResp:= response.VariantRespon{
		ID:          variant.ID,
		Name:        variant.Name,
		Description: variant.Description,
		Price:       variant.Price,
		ServiceID:   variant.ServiceID,
		CreatedAt:   variant.CreatedAt,
		UpdatedAt:   variant.UpdatedAt,
	}
	return variantResp
}

