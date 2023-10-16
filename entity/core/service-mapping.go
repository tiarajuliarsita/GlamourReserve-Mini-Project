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
		
		CreatedAt:   svcModel.CreatedAt,
		UpdatedAt:   svcModel.UpdatedAt,
		DeletedAt:   svcModel.DeletedAt.Time,
	}
	return svcCore
}

func ServiceReqToServiceCore(svcReq request.ServiceRequest) ServiceCore {
	svcCore := ServiceCore{
		Name: svcReq.Name,
		// Image:       imageurl,
		Description: svcReq.Description,
	}
	return svcCore
}
func ServiceCoreToModelsSevice(svcCore ServiceCore) models.Service {
	svcModel := models.Service{
		ID:          svcCore.ID,
		Name:        svcCore.Name,
		Description: svcCore.Description,
		// Image:       svcCore.Image,
		Variants:  []models.Variant{},
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
		Variants:    []models.Variant{},
		CreatedAt:   svcCore.CreatedAt,
		UpdatedAt:   svcCore.UpdatedAt,
	}
	return svcResponse
}
