package handlers

import (
	"glamour_reserve/entity"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
	"glamour_reserve/services"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterHandler(e echo.Context) error {
	userReq := request.UserRequest{}
	err := e.Bind(&userReq)
	if err != nil {
		return e.JSON(404, echo.Map{"error": err.Error()})

	}
	userInsert := entity.UserCore{
		UserName: userReq.UserName,
		Email:    userReq.Email,
		Password: userReq.Email,
		Phone:    userReq.Phone,
	}

	userInsert, err = h.userService.CreateUser(userInsert)
	if err != nil {
		return e.JSON(404, echo.Map{"error": err.Error()})

	}
	userResp := response.UserRespon{
		ID:        userInsert.ID,
		UserName:  userInsert.UserName,
		Email:     userInsert.Email,
		Phone:     userInsert.Phone,
		CreatedAt: userInsert.CreatedAt,
		UpdatedAt: userInsert.UpdatedAt,
	}
	return e.JSON(201, echo.Map{
		"message": "succes create user",
		"data":    userResp,
	})
}

func (h *userHandler) LoginUser(e echo.Context) error {
	userReq := request.UserRequest{}
	err := e.Bind(&userReq)
	if err != nil {
		return e.JSON(404, echo.Map{"error": err.Error()})

	}
	email := userReq.Email
	password := userReq.Password

	userData, token, err := h.userService.Login(email, password)
	if err != nil {
		return e.JSON(404, echo.Map{"error": err.Error()})
	}

	userResp := response.UserRespon{
		ID:        userData.ID,
		UserName:  userData.UserName,
		Email:     userData.Email,
		Phone:     userData.Phone,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return e.JSON(200, echo.Map{
		"message": "succes login",
		"data":    userResp,
		"token":   token,
	})

}

func (h *userHandler) GetAllUsers(e echo.Context) error {
	users, err := h.userService.FindAll()
	if err != nil {
		return e.JSON(404, echo.Map{"error": err.Error()})

	}
	usersResp := []response.UserRespon{}
	for _, v := range users {
		user := response.UserRespon{
			ID:        v.ID,
			UserName:  v.UserName,
			Email:     v.Email,
			Phone:     v.Phone,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		usersResp = append(usersResp, user)
	}
	return e.JSON(200, echo.Map{
		"user":usersResp,
	})
}
