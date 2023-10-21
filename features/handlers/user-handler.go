package handlers

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
	"glamour_reserve/features/services"

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
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	userInsert := core.UserRequestToUserCore(userReq)
	userdata, err := h.userService.CreateUser(userInsert)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	userResp := core.UserCoreToUserResponse(userdata)
	return response.RespondJSON(e, 201, "succes", userResp)

}

func (h *userHandler) LoginUser(e echo.Context) error {
	userReq := request.UserRequest{}
	err := e.Bind(&userReq)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	email := userReq.Email
	password := userReq.Password

	userData, token, err := h.userService.Login(email, password)
	if err != nil {
		return response.RespondJSON(e, 404, err.Error(), nil)
	}

	userResp := core.UserCoreToUserResponse(userData)
	return e.JSON(200, echo.Map{
		"message": "succes",
		"data":    userResp,
		"token":   token,
	})
}

func (h *userHandler) GetAllUsers(e echo.Context) error {
	users, err := h.userService.FindAll()
	if err != nil {
		return response.RespondJSON(e, 404, err.Error(), nil)
	}

	usersResp := []response.UserRespon{}
	for _, v := range users {
		user := core.UserCoreToUserResponse(v)
		usersResp = append(usersResp, user)
	}

	return response.RespondJSON(e, 200, "succes", usersResp)
}
