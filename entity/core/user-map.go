package core

import (
	"glamour_reserve/entity/models"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
)

func UserRequestToUserCore(userReq request.UserRequest) UserCore {
	userCore := UserCore{
		UserName: userReq.UserName,
		Email:    userReq.Email,
		Password: userReq.Password,
		Phone:    userReq.Phone,
	}
	return userCore
}

func UserCoreToUserModel(userCore UserCore) models.User {
	userModel := models.User{
		ID:        userCore.ID,
		UserName:  userCore.UserName,
		Email:     userCore.Email,
		Password:  userCore.Password,
		Phone:     userCore.Phone,
		CreatedAt: userCore.CreatedAt,
		UpdatedAt: userCore.UpdatedAt,
	}
	return userModel
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
	}
	return userCore
}

func UserCoreToUserResponse(userCore UserCore) response.UserRespon {
	userResp := response.UserRespon{
		ID:        userCore.ID,
		UserName:  userCore.UserName,
		Email:     userCore.Email,
		Phone:     userCore.Phone,
		CreatedAt: userCore.CreatedAt,
		UpdatedAt: userCore.UpdatedAt,
	}
	return userResp
}
