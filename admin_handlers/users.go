package admin_handlers

import (
	"asira_lender/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

func GetAllUser(c echo.Context) error {
	defer c.Request().Body.Close()

	userM := models.User{}
	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	name := c.QueryParam("name")
	id := c.QueryParam("id")
	email := c.QueryParam("email")
	phone := c.QueryParam("phone")

	type Filter struct {
		ID       string `json:"id"`
		Username string `json:"username" condition:"LIKE"`
		Email    string `json:"email" condition:"LIKE"`
		Phone    string `json:"phone" condition:"LIKE"`
	}

	result, err := userM.PagedFilterSearch(page, rows, orderby, sort, &Filter{
		ID:       id,
		Username: name,
		Email:    email,
		Phone:    phone,
	})

	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "User tidak Ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}

func UserGetDetails(c echo.Context) error {
	defer c.Request().Body.Close()

	userM := models.User{}

	userID, _ := strconv.Atoi(c.Param("user_id"))
	err := userM.FindbyID(userID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "User ID tidak ditemukan")
	}

	return c.JSON(http.StatusOK, userM)
}

func AddUser(c echo.Context) error {
	defer c.Request().Body.Close()

	userM := models.User{}

	payloadRules := govalidator.MapData{
		"username": []string{"required"},
		"email":    []string{"required"},
		"phone":    []string{"required"},
		"role_id":  []string{"required", "role_id"},
		"status":   []string{},
	}

	validate := validateRequestPayload(c, payloadRules, &userM)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}
	userM.Password = RandString(8)
	err := userM.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat User")
	}

	return c.JSON(http.StatusCreated, userM)
}

func UpdateUser(c echo.Context) error {
	defer c.Request().Body.Close()
	userID, _ := strconv.Atoi(c.Param("user_id"))

	userM := models.User{}
	err := userM.FindbyID(userID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("User %v tidak ditemukan", userID))
	}

	payloadRules := govalidator.MapData{
		"username": []string{"required"},
		"email":    []string{"required"},
		"phone":    []string{"required"},
		"role_id":  []string{"required", "role_id"},
		"status":   []string{},
	}

	validate := validateRequestPayload(c, payloadRules, &userM)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err = userM.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update User %v", userID))
	}

	return c.JSON(http.StatusOK, userM)
}
