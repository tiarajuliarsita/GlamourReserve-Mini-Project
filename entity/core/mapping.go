package core

import (
	"glamour_reserve/entity/models"
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

func UserModelToUserCore(userModel models.User) UserCore {
	userCore := UserCore{
		ID:        userModel.ID,
		UserName:  userModel.UserName,
		Email:     userModel.Email,
		Password:  userModel.Password,
		Phone:     userModel.Phone,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
		DeletedAt: userModel.DeletedAt.Time,
	}
	return userCore
}

func UserCoreToUserResponse(userCore UserCore)response.UserRespon{
	userResp:= response.UserRespon{
		ID:        userCore.ID,
		UserName:  userCore.UserName,
		Email:     userCore.Email,
		Phone:     userCore.Phone,
		CreatedAt: userCore.CreatedAt,
		UpdatedAt: userCore.UpdatedAt,
	}
	return userResp
}