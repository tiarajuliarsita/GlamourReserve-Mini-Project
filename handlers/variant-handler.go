package handlers

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
	"glamour_reserve/services"

	"github.com/labstack/echo/v4"
)

type VariantHandler struct {
	variantService services.VariantServiceInterface
}

func NewVariantHandler(variantService services.VariantServiceInterface) *VariantHandler {
	return &VariantHandler{variantService}
}

func (h *VariantHandler) CreateVariant(e echo.Context) error {
	variantReq := request.VariantRequest{}
	err := e.Bind(&variantReq)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	variant := core.VariantRequestToVariantCore(variantReq)
	variantData, err := h.variantService.Create(variant)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	variantResp := core.VariantCoreToVariantRespon(variantData)
	return response.RespondJSON(e, 201, "succes", variantResp)
}

func (h *VariantHandler) GetByID(e echo.Context) error {
	id := e.Param("id")
	dataVariant, err := h.variantService.FindByID(id)
	if err!=nil{
		return response.RespondJSON(e, 400, err.Error(), nil)
	}
	variantResp:=core.VariantCoreToVariantRespon(dataVariant)
	return response.RespondJSON(e, 200, "succes", variantResp)

}
