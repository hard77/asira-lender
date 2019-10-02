package admin_handlers

import (
	"asira_lender/asira"
	"asira_lender/models"
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
	role_id := c.QueryParam("role_id")

	type Filter struct {
		ID         string `json:"id"`
		RoleID     string `json:"role_id"`
		Permission string `json:"permissions" condition:"LIKE"`
	}

	result, err := Permission.PagedFilterSearch(page, rows, orderby, sort, &Filter{
		ID:         id,
		RoleID:     role_id,
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

	type (
		validatePermissions struct {
			RoleID      string   `json:"role_id"`
			Permissions []string `json:"permissions"`
		}
	)
	valPermissions := validatePermissions{}

	payloadRules := govalidator.MapData{
		"role_id":     []string{"required", "role_id", "numeric"},
		"permissions": []string{},
	}

	validate := validateRequestPayload(c, payloadRules, &valPermissions)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}
	Permissions := []models.Permissions{}
	RoleID, _ := strconv.Atoi(valPermissions.RoleID)
	for _, n := range valPermissions.Permissions {
		Permissions = append(Permissions, models.Permissions{
			RoleID:      RoleID,
			Permissions: n,
		})
	}

	for _, per := range Permissions {
		per.Create()
	}
	return c.JSON(http.StatusCreated, valPermissions)
}

func UpdatePermission(c echo.Context) error {
	defer c.Request().Body.Close()
	type (
		validatePermissions struct {
			RoleID      string   `json:"role_id"`
			Permissions []string `json:"permissions"`
		}
	)
	valPermissions := validatePermissions{}
	Permissions := []models.Permissions{}

	payloadRules := govalidator.MapData{
		"role_id":     []string{"required", "role_id", "numeric"},
		"permissions": []string{},
	}

	validate := validateRequestPayload(c, payloadRules, &valPermissions)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	RoleID, _ := strconv.Atoi(valPermissions.RoleID)
	asira.App.DB.Where("role_id = ?", RoleID).Delete(&Permissions)

	for _, n := range valPermissions.Permissions {
		Permissions = append(Permissions, models.Permissions{
			RoleID:      RoleID,
			Permissions: n,
		})
	}

	for _, per := range Permissions {
		per.Create()
	}

	return c.JSON(http.StatusOK, valPermissions)
}
