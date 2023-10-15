package handlers

import (
	"glamour_reserve/entity/models"
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
		return e.JSON(400, echo.Map{"error": err.Error()})
	}

	// Mengecek apakah svcRequest.Image tidak nil sebelum mengunggah.
	// if svcRequest.Image == nil {
	// 	return e.JSON(400, echo.Map{"error": "Image is required"})
	// }

	// fileName := helpers.RemoveExtention(svcRequest.Image.Filename)
	// uploadResult, err := helpers.UploadFile(svcRequest.Image, fileName)
	// if err != nil {
	// 	return e.JSON(500, echo.Map{"error": err.Error()})
	// }
	service := models.Service{
		Name:        svcRequest.Name,
		Description: svcRequest.Description,
		// Image:       uploadResult,
	}
	data, err := h.svcService.CreateService(&service)
	if err != nil {
		return e.JSON(500, echo.Map{"error": err.Error()})
	}

	return e.JSON(200, echo.Map{
		"message":  "Berhasil membuat layanan",
		"services": data,
	})
}
