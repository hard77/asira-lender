package admin_handlers

import (
	"asira_lender/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

func BankServiceProductList(c echo.Context) error {
	defer c.Request().Body.Close()

	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	// filters
	name := c.QueryParam("name")

	type Filter struct {
		Name string `json:"name" condition:"LIKE"`
	}

	bank_service := models.ServiceProduct{}
	result, err := bank_service.PagedFilterSearch(page, rows, orderby, sort, &Filter{
		Name: name,
	})
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}

func BankServiceProductNew(c echo.Context) error {
	defer c.Request().Body.Close()

	serviceProduct := models.ServiceProduct{}

	payloadRules := govalidator.MapData{
		"name":             []string{"required"},
		"min_timespan":     []string{"numeric"},
		"max_timespan":     []string{"numeric"},
		"interest":         []string{"numeric"},
		"min_loan":         []string{"numeric"},
		"max_loan":         []string{"numeric"},
		"fees":             []string{},
		"asn_fee":          []string{"regex:^(\\d|\\d%)+$"}, // fixed number or percentage
		"service":          []string{"numeric"},
		"collaterals":      []string{},
		"financing_sector": []string{},
		"assurance":        []string{},
		"status":           []string{"active_inactive"},
	}

	validate := validateRequestPayload(c, payloadRules, &serviceProduct)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err := serviceProduct.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat layanan bank baru")
	}

	return c.JSON(http.StatusCreated, serviceProduct)
}

func BankServiceProductDetail(c echo.Context) error {
	defer c.Request().Body.Close()

	product_id, _ := strconv.Atoi(c.Param("product_id"))

	serviceProduct := models.ServiceProduct{}
	err := serviceProduct.FindbyID(product_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("product %v tidak ditemukan", serviceProduct))
	}

	return c.JSON(http.StatusOK, serviceProduct)
}

func BankServiceProductPatch(c echo.Context) error {
	defer c.Request().Body.Close()

	product_id, _ := strconv.Atoi(c.Param("product_id"))

	serviceProduct := models.ServiceProduct{}
	err := serviceProduct.FindbyID(product_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("product %v tidak ditemukan", product_id))
	}

	productPayloadRules := govalidator.MapData{
		"name":             []string{},
		"min_timespan":     []string{"numeric"},
		"max_timespan":     []string{"numeric"},
		"interest":         []string{"numeric"},
		"min_loan":         []string{"numeric"},
		"max_loan":         []string{"numeric"},
		"fees":             []string{},
		"asn_fee":          []string{"regex:^(\\d|\\d%)+$"}, // fixed number or percentage
		"service":          []string{"numeric"},
		"collaterals":      []string{},
		"financing_sector": []string{},
		"assurance":        []string{},
		"status":           []string{"active_inactive"},
	}
	validate := validateRequestPayload(c, productPayloadRules, &serviceProduct)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err = serviceProduct.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update product %v", product_id))
	}

	return c.JSON(http.StatusOK, serviceProduct)
}

func BankServiceProductDelete(c echo.Context) error {
	defer c.Request().Body.Close()

	product_id, _ := strconv.Atoi(c.Param("product_id"))

	serviceProduct := models.ServiceProduct{}
	err := serviceProduct.FindbyID(product_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("bank type %v tidak ditemukan", product_id))
	}

	err = serviceProduct.Delete()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update bank tipe %v", product_id))
	}

	return c.JSON(http.StatusOK, serviceProduct)
}
