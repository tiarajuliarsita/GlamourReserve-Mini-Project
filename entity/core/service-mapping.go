package core

import "glamour_reserve/entity/models"

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
