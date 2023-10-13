package handlers

import (
	"glamour_reserve/entity/response"
	"glamour_reserve/services"

	"github.com/labstack/echo/v4"
)

type serviceHandler struct {
	svcService services.SvcServiceInterface
}

func NewServiceHandler(svcService services.SvcServiceInterface) *serviceHandler {
	return &serviceHandler{svcService}
}

func (h *serviceHandler) GetAllServices(e echo.Context) error {
	services, err := h.svcService.FindAll()
	if err != nil {
		return e.JSON(500, echo.Map{"error": err.Error()})

	}

	servicesResponse := []response.ServiceRespon{}
	for _, v := range services {
		service := response.ServiceRespon{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		servicesResponse = append(servicesResponse, service)
	}
	return e.JSON(200, echo.Map{
		"message":  "succes get all services",
		"services": servicesResponse,
	})

}
