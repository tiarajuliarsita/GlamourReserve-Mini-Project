package services

import (
	"errors"
	"glamour_reserve/entity/core"
	"glamour_reserve/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockUserRepoInterface(ctrl)
	service := NewUserService(mockRepo)

	userToCreate := core.UserCore{
		UserName: "tiara",
		Email:    "tiara@gmail.com",
		Password: "123",
		Phone:    "0867",
	}

	mockRepo.EXPECT().CreateUser(userToCreate).Return(userToCreate, nil)
	createdUser, err := service.CreateUser(userToCreate)

	assert.NoError(t, err)
	assert.Equal(t, userToCreate, createdUser)
}

func TestUserService_CreateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepoInterface(ctrl)
	service := NewUserService(mockRepo)

	userToCreate := core.UserCore{
		UserName: "tiara",
		Email:    "tiara@gmail.com",
		Password: "123",
		Phone:    "0867",
	}

	expectedErr := errors.New("an error occurred")

	mockRepo.EXPECT().CreateUser(userToCreate).Return(core.UserCore{}, expectedErr)
	createdUser, err := service.CreateUser(userToCreate)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, core.UserCore{}, createdUser)
}




func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepoInterface(ctrl)
	service := NewUserService(mockRepo)

	// Define login credentials
	email := "tiarajuliarsita.com"
	password := "password123"
	expectedUser := core.UserCore{
		ID:       "1",
		UserName: "tiara",
		Email:    "tiara@gmail.com",
		Password: "hashed_password",
		Role:     "user",
	}

	mockRepo.EXPECT().Login(email, password).Return(expectedUser, nil)
	user, token, err := service.Login(email, password)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	assert.NotEmpty(t, token)
}



func TestUserService_Login_ErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepoInterface(ctrl)
	service := NewUserService(mockRepo)

	email := "tiarajuliarsita.com"
	password := "password123"
	expectedErr := errors.New("authentication failed")

	mockRepo.EXPECT().Login(email, password).Return(core.UserCore{}, expectedErr)
	user, token, err := service.Login(email, password)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, user)
	assert.Empty(t, token)
}


func TestUserService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepoInterface(ctrl)
	service := NewUserService(mockRepo)
	expectedUsers := []core.UserCore{
		{
			ID:       "1",
			UserName: "yaya",
			Email:    "yaya@egmail.com",
		},
		{
			ID:       "2",
			UserName: "yaya",
			Email:    "tiara@gmail.com",
		},
	}

	mockRepo.EXPECT().FindAll().Return(expectedUsers, nil)
	users, err := service.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestUserService_FindAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepoInterface(ctrl)
	service := NewUserService(mockRepo)
	expectedErr := errors.New("an error occurred")

	mockRepo.EXPECT().FindAll().Return(nil, expectedErr)
	users, err := service.FindAll()
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, users)
}

