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

func ProductList(c echo.Context) error {
	defer c.Request().Body.Close()

	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	order := strings.Split(c.QueryParam("orderby"), ",")
	sort := strings.Split(c.QueryParam("sort"), ",")

	// filters
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	serviceId := c.QueryParam("service_id")
	status := c.QueryParam("status")

	type Filter struct {
		ID        string `json:"id"`
		Name      string `json:"name" condition:"LIKE"`
		ServiceID string `json:"service_id`
		Status    string `json:"status"`
	}

	product := models.Product{}
	result, err := product.PagedFindFilter(page, rows, order, sort, &Filter{
		ID:        id,
		Name:      name,
		ServiceID: serviceId,
		Status:    status,
	})
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}

func ProductNew(c echo.Context) error {
	defer c.Request().Body.Close()

	product := models.Product{}

	payloadRules := govalidator.MapData{
		"name":       []string{"required"},
		"service_id": []string{"required", "valid_id:services"},
		"status":     []string{"required", "active_inactive"},
	}

	validate := validateRequestPayload(c, payloadRules, &product)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err := product.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat layanan bank baru")
	}

	return c.JSON(http.StatusCreated, product)
}

func ProductDetail(c echo.Context) error {
	defer c.Request().Body.Close()

	productId, _ := strconv.Atoi(c.Param("id"))

	product := models.Product{}
	err := product.FindbyID(productId)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("product %v tidak ditemukan", productId))
	}

	return c.JSON(http.StatusOK, product)
}

func ProductPatch(c echo.Context) error {
	defer c.Request().Body.Close()

	productId, _ := strconv.Atoi(c.Param("id"))

	product := models.Product{}
	err := product.FindbyID(productId)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("product %v tidak ditemukan", productId))
	}

	payloadRules := govalidator.MapData{
		"name":       []string{"required"},
		"service_id": []string{"required"},
		"status":     []string{"required", "active_inactive"},
	}
	validate := validateRequestPayload(c, payloadRules, &product)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err = product.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update layanan %v", productId))
	}

	return c.JSON(http.StatusOK, product)
}

func ProductDelete(c echo.Context) error {
	defer c.Request().Body.Close()

	productId, _ := strconv.Atoi(c.Param("id"))

	product := models.Product{}
	err := product.FindbyID(productId)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("product %v tidak ditemukan", productId))
	}

	err = product.Delete()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update bank tipe %v", productId))
	}

	return c.JSON(http.StatusOK, product)
}
