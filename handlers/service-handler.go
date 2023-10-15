package handlers

import (
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/request"
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

func (h *serviceHandler) CreateService(e echo.Context) error {

	svcRequest := request.ServiceRequest{}
	if err := e.Bind(&svcRequest); err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	dataService := core.ServiceReqToServiceCore(svcRequest)
	data, err := h.svcService.CreateService(dataService)
	if err != nil {
		return response.RespondJSON(e, 500, err.Error(), nil)
	}

	servisResp := core.ServiceCoreToResponseService(data)
	return response.RespondJSON(e, 201, "succes", servisResp)
}

func (h *serviceHandler) GetServiceByID(e echo.Context) error {
	id := e.Param("id")
	service, err := h.svcService.FindById(id)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}
	svcResp := core.ServiceCoreToResponseService(service)
	return response.RespondJSON(e, 200, "succes", svcResp)
}
