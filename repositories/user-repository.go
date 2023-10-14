package repositories

import (
	"glamour_reserve/entity"
	"glamour_reserve/entity/models"
	"glamour_reserve/helpers"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(entity.UserCore) (entity.UserCore, error)
	Login(email string, password string) (entity.UserCore, error)
	FindAll() ([]entity.UserCore, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB}
}

func (r *userRepository) FindAll() ([]entity.UserCore, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	usersCore := []entity.UserCore{}
	for _, v := range users {
		user := entity.UserCore{
			ID:        v.ID,
			UserName:  v.UserName,
			Email:     v.Email,
			Password:  v.Password,
			Phone:     v.Phone,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		usersCore = append(usersCore, user)
	}

	return usersCore, nil

}

func (r *userRepository) CreateUser(user entity.UserCore) (entity.UserCore, error) {
	userInsert := models.User{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	err := r.db.Create(&userInsert).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Login(email string, password string) (entity.UserCore, error) {
	var user models.User
	var dataUser entity.UserCore

	err := r.db.Where("email= ?", email).First(&user).Error
	if err != nil {

		return dataUser, err
	}

	comparePass, err := helpers.ComparePass([]byte(user.Password), []byte(password))
	if err != nil {
		return dataUser, err
	}

	if !comparePass {
		return dataUser, err
	}

	dataUser.CreatedAt = user.CreatedAt
	dataUser.UpdatedAt = user.UpdatedAt
	dataUser.UserName = user.UserName
	dataUser.Email = user.Email
	dataUser.Phone = user.Phone
	dataUser.ID = user.ID
	dataUser.Password = user.Password

	return dataUser, nil
}
