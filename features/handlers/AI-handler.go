package handlers

import (
	"fmt"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
	"glamour_reserve/features/services"
	"os"

	"github.com/labstack/echo/v4"
)

type beautyCareHandler struct {
	beautySvc services.BeautyCareInterface
}

func NewBeautyCare(beautySvc services.BeautyCareInterface) *beautyCareHandler {
	return &beautyCareHandler{beautySvc}
}

func (h *beautyCareHandler) AskAboutBeauty(e echo.Context) error {
	request := request.AskBeautyReq{}
	err := e.Bind(&request)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	userInput := fmt.Sprintf("Produk kecantikan %s yang sesuai untuk kulit %s dengan concern %s dengan biaya sebesar Rp %.0f.", request.Brand, request.SkinType, request.Concern, request.Budget)

	answer, err := h.beautySvc.AskAboutBeauty(userInput, request.Brand, request.SkinType, os.Getenv("OPEN_AI_KEY"))
	if err != nil {
		return response.RespondJSON(e, 500, err.Error(), nil)
	}

	return response.RespondJSON(e, 200, "succes", answer)

}
