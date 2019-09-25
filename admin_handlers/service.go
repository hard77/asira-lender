package admin_handlers

import (
	"asira_lender/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

func ServiceList(c echo.Context) error {
	defer c.Request().Body.Close()

	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	order := strings.Split(c.QueryParam("orderby"), ",")
	sort := strings.Split(c.QueryParam("sort"), ",")

	// filters
	id := c.QueryParam("id")

	type Filter struct {
		ID string `json:"id"`
	}

	service := models.Service{}
	result, err := service.PagedFindFilter(page, rows, order, sort, &Filter{
		ID: id,
	})
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}

func ServiceNew(c echo.Context) error {
	defer c.Request().Body.Close()

	service := models.Service{}

	payloadRules := govalidator.MapData{
		"name":   []string{"required"},
		"status": []string{"required", "active_inactive"},
	}

	validate := validateRequestPayload(c, payloadRules, &service)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err := service.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat layanan bank baru")
	}

	return c.JSON(http.StatusCreated, service)
}

func ServiceDetail(c echo.Context) error {
	defer c.Request().Body.Close()

	serviceId, _ := strconv.Atoi(c.Param("id"))

	service := models.Service{}
	err := service.FindbyID(serviceId)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("layanan %v tidak ditemukan", serviceId))
	}

	return c.JSON(http.StatusOK, service)
}

func ServicePatch(c echo.Context) error {
	defer c.Request().Body.Close()

	serviceId, _ := strconv.Atoi(c.Param("id"))

	service := models.Service{}
	err := service.FindbyID(serviceId)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("layanan %v tidak ditemukan", serviceId))
	}

	payloadRules := govalidator.MapData{
		"name":   []string{"required"},
		"status": []string{"active_inactive"},
	}
	validate := validateRequestPayload(c, payloadRules, &service)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err = service.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update layanan %v", serviceId))
	}

	return c.JSON(http.StatusOK, service)
}

func ServiceDelete(c echo.Context) error {
	defer c.Request().Body.Close()

	serviceId, _ := strconv.Atoi(c.Param("id"))

	service := models.Service{}
	err := service.FindbyID(serviceId)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("bank type %v tidak ditemukan", serviceId))
	}

	err = service.Delete()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update bank tipe %v", serviceId))
	}

	return c.JSON(http.StatusOK, service)
}
