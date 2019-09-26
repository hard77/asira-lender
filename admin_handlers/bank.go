package admin_handlers

import (
	"asira_lender/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

func BankList(c echo.Context) error {
	defer c.Request().Body.Close()

	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	// filters
	name := c.QueryParam("name")
	id := c.QueryParam("id")

	type Filter struct {
		ID   string `json:"id"`
		Name string `json:"name" condition:"LIKE"`
	}

	bank := models.Bank{}
	result, err := bank.PagedFilterSearch(page, rows, orderby, sort, &Filter{
		ID:   id,
		Name: name,
	})
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}

func BankNew(c echo.Context) error {
	defer c.Request().Body.Close()

	bank := models.Bank{}

	payloadRules := govalidator.MapData{
		"name":           []string{"required"},
		"type":           []string{"bank_type_id"},
		"address":        []string{"required"},
		"province":       []string{"required"},
		"city":           []string{"required"},
		"services":       []string{},
		"products":       []string{},
		"pic":            []string{"required"},
		"phone":          []string{"required"},
		"adminfee_setup": []string{"required"},
		"convfee_setup":  []string{"required"},
		// "username": []string{"required", "unique:banks,username"},
		// "password": []string{"required"},
	}

	validate := validateRequestPayload(c, payloadRules, &bank)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err := bank.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat tipe bank baru")
	}

	return c.JSON(http.StatusCreated, bank)
}

func BankDetail(c echo.Context) error {
	defer c.Request().Body.Close()

	bank_id, _ := strconv.Atoi(c.Param("bank_id"))

	bank := models.Bank{}
	err := bank.FindbyID(bank_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("bank type %v tidak ditemukan", bank_id))
	}

	return c.JSON(http.StatusOK, bank)
}

func BankPatch(c echo.Context) error {
	defer c.Request().Body.Close()

	bank_id, _ := strconv.Atoi(c.Param("bank_id"))

	bank := models.Bank{}
	err := bank.FindbyID(bank_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("bank type %v tidak ditemukan", bank_id))
	}

	// dont allow admin to change bank credentials
	tempUsername := bank.Username
	tempPassword := bank.Password

	payloadRules := govalidator.MapData{
		"name":           []string{},
		"type":           []string{"bank_type_id"},
		"address":        []string{},
		"province":       []string{},
		"city":           []string{},
		"services":       []string{},
		"products":       []string{},
		"pic":            []string{},
		"phone":          []string{},
		"adminfee_setup": []string{},
		"convfee_setup":  []string{},
	}

	validate := validateRequestPayload(c, payloadRules, &bank)
	log.Println(bank)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	bank.Username = tempUsername
	bank.Password = tempPassword

	err = bank.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update bank tipe %v", bank_id))
	}

	return c.JSON(http.StatusOK, bank)
}

func BankDelete(c echo.Context) error {
	defer c.Request().Body.Close()

	bank_id, _ := strconv.Atoi(c.Param("bank_id"))

	bank := models.Bank{}
	err := bank.FindbyID(bank_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("bank type %v tidak ditemukan", bank_id))
	}

	err = bank.Delete()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update bank tipe %v", bank_id))
	}

	return c.JSON(http.StatusOK, bank)
}
