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

func BankProductList(c echo.Context) error {
	defer c.Request().Body.Close()

	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	order := strings.Split(c.QueryParam("order"), ",")
	sort := strings.Split(c.QueryParam("sort"), ",")

	// filters
	productID := c.QueryParam("product_id")
	bankServiceID := c.QueryParam("bank_service_id")
	minTimespan := c.QueryParam("min_timespan")
	maxTimespan := c.QueryParam("max_timespan")
	interest := c.QueryParam("interest")
	minLoan := c.QueryParam("min_loan")
	maxLoan := c.QueryParam("max_loan")
	fee := c.QueryParam("fee")
	collaterals := c.QueryParam("collaterals")
	financingSector := c.QueryParam("financing_sector")
	assurance := c.QueryParam("assurance")
	status := c.QueryParam("status")

	type Filter struct {
		ProductID       string `json:"product_id"`
		BankServiceID   string `json:"bank_service_id"`
		MinTimeSpan     string `json:"min_timespan"`
		MaxTimeSpan     string `json:"max_timespan"`
		Interest        string `json:"interest" condition:"LIKE"`
		MinLoan         string `json:"min_loan"`
		MaxLoan         string `json:"max_loan"`
		Fees            string `json:"fees" condition:"LIKE"`
		Collaterals     string `json:"collaterals" condition:"LIKE"`
		FinancingSector string `json:"financing_sector" condition:"LIKE"`
		Assurance       string `json:"assurance" condition:"LIKE"`
		Status          string `json:"status" condition:"LIKE"`
	}

	bank_product := models.BankProduct{}
	result, err := bank_product.PagedFindFilter(page, rows, order, sort, &Filter{
		ProductID:       productID,
		BankServiceID:   bankServiceID,
		MinTimeSpan:     minTimespan,
		MaxTimeSpan:     maxTimespan,
		Interest:        interest,
		MinLoan:         minLoan,
		MaxLoan:         maxLoan,
		Fees:            fee,
		Collaterals:     collaterals,
		FinancingSector: financingSector,
		Assurance:       assurance,
		Status:          status,
	})
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}

func BankProductNew(c echo.Context) error {
	defer c.Request().Body.Close()

	bankProduct := models.BankProduct{}

	payloadRules := govalidator.MapData{
		"product_id":       []string{"required"},
		"bank_service_id":  []string{"required"},
		"min_timespan":     []string{"required", "numeric"},
		"max_timespan":     []string{"required", "numeric"},
		"interest":         []string{"required", "numeric"},
		"min_loan":         []string{"required", "numeric"},
		"max_loan":         []string{"required", "numeric"},
		"fees":             []string{},
		"collaterals":      []string{"required"},
		"financing_sector": []string{"required"},
		"assurance":        []string{"required"},
		"status":           []string{"required", "active_inactive"},
	}

	validate := validateRequestPayload(c, payloadRules, &bankProduct)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err := bankProduct.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat layanan bank baru")
	}

	return c.JSON(http.StatusCreated, bankProduct)
}

func BankProductDetail(c echo.Context) error {
	defer c.Request().Body.Close()

	product_id, _ := strconv.Atoi(c.Param("id"))

	bankProduct := models.BankProduct{}
	err := bankProduct.FindbyID(product_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("product %v tidak ditemukan", bankProduct))
	}

	return c.JSON(http.StatusOK, bankProduct)
}

func BankProductPatch(c echo.Context) error {
	defer c.Request().Body.Close()

	product_id, _ := strconv.Atoi(c.Param("id"))

	bankProduct := models.BankProduct{}
	err := bankProduct.FindbyID(product_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("product %v tidak ditemukan", product_id))
	}

	productPayloadRules := govalidator.MapData{
		"product_id":       []string{},
		"bank_service_id":  []string{},
		"min_timespan":     []string{"numeric"},
		"max_timespan":     []string{"numeric"},
		"interest":         []string{"numeric"},
		"min_loan":         []string{"numeric"},
		"max_loan":         []string{"numeric"},
		"fees":             []string{},
		"collaterals":      []string{},
		"financing_sector": []string{},
		"assurance":        []string{},
		"status":           []string{"active_inactive"},
	}
	validate := validateRequestPayload(c, productPayloadRules, &bankProduct)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err = bankProduct.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update product %v", product_id))
	}

	return c.JSON(http.StatusOK, bankProduct)
}

func BankProductDelete(c echo.Context) error {
	defer c.Request().Body.Close()

	product_id, _ := strconv.Atoi(c.Param("id"))

	bankProduct := models.BankProduct{}
	err := bankProduct.FindbyID(product_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("bank type %v tidak ditemukan", product_id))
	}

	err = bankProduct.Delete()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update bank tipe %v", product_id))
	}

	return c.JSON(http.StatusOK, bankProduct)
}
