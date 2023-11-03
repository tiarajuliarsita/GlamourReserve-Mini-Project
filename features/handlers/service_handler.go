package handlers

import (
	core "glamour_reserve/entity/core"
	"glamour_reserve/entity/request"
	"glamour_reserve/entity/response"
	"glamour_reserve/features/services"
	"glamour_reserve/utils/helpers/authentication"
	"glamour_reserve/utils/helpers/cloud"
	"glamour_reserve/utils/helpers/constanta"
	"net/http"

	"github.com/labstack/echo/v4"
)

type serviceHandler struct {
	svcService services.SvcServiceInterface
}

func NewServiceHandler(svcService services.SvcServiceInterface) *serviceHandler {
	return &serviceHandler{svcService}
}

func (h *serviceHandler) GetAllServices(e echo.Context) error {
	name := e.QueryParam("name")
	offset := e.QueryParam("offset")
	limit := e.QueryParam("limit")

	services, err := h.svcService.FindAll(name, offset, limit)
	if err != nil {
		return response.RespondJSON(e, http.StatusInternalServerError, err.Error(), nil)
	}

	servicesResponse := []response.ServiceRespon{}
	for _, v := range services {
		service := core.ServiceCoreToResponseService(v)
		servicesResponse = append(servicesResponse, service)
	}

	return response.RespondJSON(e, http.StatusOK, "succes", servicesResponse)
}

func (h *serviceHandler) CreateService(e echo.Context) error {

	_, _, role := authentication.ExtractTokenUserId(e)

	if role != constanta.ADMIN{
		return response.RespondJSON(e, http.StatusForbidden, "you don`t have permission", nil)
	}

	svcRequest := request.ServiceRequest{}
	if err := e.Bind(&svcRequest); err != nil {
		return response.RespondJSON(e, http.StatusBadRequest, err.Error(), nil)
	}

	file, err := e.FormFile("image")
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Failed to receive file")
	}

	client := cloud.ConfigCloud()
	imageurl := cloud.UploadFile(file, client)
	dataService := core.ServiceReqToServiceCore(svcRequest, imageurl)
	data, err := h.svcService.CreateService(dataService)
	if err != nil {
		return response.RespondJSON(e, http.StatusInternalServerError, err.Error(), nil)
	}

	servisResp := core.ServiceCoreToResponseService(data)
	return response.RespondJSON(e, http.StatusCreated, "succes", servisResp)
}

func (h *serviceHandler) GetServiceByID(e echo.Context) error {
	id := e.Param("id")
	service, err := h.svcService.FindById(id)
	if err != nil {
		return response.RespondJSON(e, http.StatusNotFound, err.Error(), nil)
	}

	svcResp := core.ServiceCoreToResponseService(service)
	return response.RespondJSON(e, http.StatusOK, "succes", svcResp)
}

func (h *serviceHandler) DeletByID(e echo.Context) error {
	_, _, role := authentication.ExtractTokenUserId(e)
	if role != constanta.ADMIN{
		return response.RespondJSON(e, http.StatusForbidden, "you don`t have permission", nil)
	}

	id := e.Param("id")
	ok, err := h.svcService.Delete(id)
	if err != nil {
		return response.RespondJSON(e, http.StatusInternalServerError, err.Error(), nil)
	}
	if !ok {
		return response.RespondJSON(e, http.StatusNotFound, err.Error(), nil)
	}
	return response.RespondJSON(e, http.StatusOK, "succes", nil)

}

func (h *serviceHandler) UpdateByID(e echo.Context) error {
	_, _, role := authentication.ExtractTokenUserId(e)
	if role != constanta.ADMIN {
		return response.RespondJSON(e, 401, "can`t update service", nil)
	}

	id := e.Param("id")
	newService := request.ServiceRequest{}

	err := e.Bind(&newService)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}

	file, err := e.FormFile("image")
	if err != nil {
		return e.JSON(400, "Failed to receive file")
	}

	client := cloud.ConfigCloud()
	imageurl := cloud.UploadFile(file, client)

	NewService := core.ServiceReqToServiceCore(newService, imageurl)
	dataService, err := h.svcService.Update(id, NewService)
	if err != nil {
		return response.RespondJSON(e, 400, err.Error(), nil)
	}
	dataService.Image = imageurl
	serviceRespon := core.ServiceCoreToResponseService(dataService)
	return response.RespondJSON(e, 200, "succes", serviceRespon)
}
