package admin_handlers

import (
	"asira_lender/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

func GetAllPermission(c echo.Context) error {
	defer c.Request().Body.Close()

	Permission := models.Permissions{}
	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	name := c.QueryParam("name")
	id := c.QueryParam("id")

	type Filter struct {
		ID         string `json:"id"`
		Permission string `json:"permissions" condition:"LIKE"`
	}

	result, err := Permission.PagedFilterSearch(page, rows, orderby, sort, &Filter{
		ID:         id,
		Permission: name,
	})

	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "Permission tidak Ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}

func PermissionGetDetails(c echo.Context) error {
	defer c.Request().Body.Close()

	Permission := models.Permissions{}

	PermissionID, _ := strconv.Atoi(c.Param("permission_id"))
	err := Permission.FindbyID(PermissionID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "Permission tidak ditemukan")
	}

	return c.JSON(http.StatusOK, Permission)
}

func AddPermission(c echo.Context) error {
	defer c.Request().Body.Close()

	Permission := models.Permissions{}

	payloadRules := govalidator.MapData{
		"role_id":     []string{"required", "role_id"},
		"permissions": []string{},
	}

	validate := validateRequestPayload(c, payloadRules, &Permission)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err := Permission.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat Permissions")
	}

	return c.JSON(http.StatusCreated, Permission)
}

func UpdatePermission(c echo.Context) error {
	defer c.Request().Body.Close()
	Permission_id, _ := strconv.Atoi(c.Param("permission_id"))

	Permission := models.Permissions{}
	err := Permission.FindbyID(Permission_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("Permission %v tidak ditemukan", Permission_id))
	}

	payloadRules := govalidator.MapData{
		"role_id":     []string{"required", "role_id"},
		"permissions": []string{},
	}

	validate := validateRequestPayload(c, payloadRules, &Permission)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err = Permission.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update Permissions %v", Permission_id))
	}

	return c.JSON(http.StatusOK, Permission)
}
