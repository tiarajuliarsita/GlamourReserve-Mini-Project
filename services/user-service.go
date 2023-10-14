package services

import (
	"glamour_reserve/entity"
	"glamour_reserve/helpers"
	"glamour_reserve/repositories"
)

type UserServiceInterface interface {
	CreateUser(user entity.UserCore) (entity.UserCore, error)
	Login(email string, password string) (entity.UserCore, string, error)
	FindAll() ([]entity.UserCore, error)
}

type userService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) *userService {
	return &userService{repo}
}

func (s *userService) CreateUser(user entity.UserCore) (entity.UserCore, error) {
	user, err := s.repo.CreateUser(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) Login(email string, password string) (entity.UserCore, string, error) {
	userData, err := s.repo.Login(email, password)
	if err != nil {
		return userData, "", err
	}

	token := helpers.GenerateToken(userData.ID, userData.UserName, userData.Email)
	return userData, token, nil
}

func (s *userService) FindAll() ([]entity.UserCore, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
