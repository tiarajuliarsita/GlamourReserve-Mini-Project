package core

import (
	"glamour_reserve/entity/models"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
)

func ServiceModelToServiceCore(svcModel models.Service) ServiceCore {
	svcCore := ServiceCore{
		ID:          svcModel.ID,
		Name:        svcModel.Name,
		Description: svcModel.Description,
		Image:       svcModel.Image,
		Price:       svcModel.Price,
		CreatedAt:   svcModel.CreatedAt,
		UpdatedAt:   svcModel.UpdatedAt,
	}
	return svcCore
}

func ServiceReqToServiceCore(svcReq request.ServiceRequest, imageUrl string) ServiceCore {
	svcCore := ServiceCore{
		Name:        svcReq.Name,
		Image:       imageUrl,
		Description: svcReq.Description,
		Price:       svcReq.Price,
	}
	return svcCore
}
func ServiceCoreToModelsSevice(svcCore ServiceCore) models.Service {
	svcModel := models.Service{
		ID:          svcCore.ID,
		Name:        svcCore.Name,
		Description: svcCore.Description,
		Image:       svcCore.Image,
		// Variants:  []models.Variant{},
		Price:     svcCore.Price,
		CreatedAt: svcCore.CreatedAt,
		UpdatedAt: svcCore.UpdatedAt,
	}
	return svcModel
}

func ServiceCoreToResponseService(svcCore ServiceCore) response.ServiceRespon {
	svcResponse := response.ServiceRespon{
		ID:          svcCore.ID,
		Name:        svcCore.Name,
		Description: svcCore.Description,
		Image:       svcCore.Image,
		// Variants:    []models.Variant{},
		Price:     svcCore.Price,
		CreatedAt: svcCore.CreatedAt,
		UpdatedAt: svcCore.UpdatedAt,
	}
	return svcResponse
}
