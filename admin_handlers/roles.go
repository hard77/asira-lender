package admin_handlers

import (
	"asira_lender/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllRole(c echo.Context) error {
	defer c.Request().Body.Close()

	Iroles := models.InternalRoles{}
	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	var Filter struct{}
	result, err := Iroles.PagedFilterSearch(page, rows, orderby, sort, &Filter)

	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "Internal Role tidak Ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}

func RoleGetDetails(c echo.Context) error {
	defer c.Request().Body.Close()

	Iroles := models.InternalRoles{}

	IrolesID, _ := strconv.Atoi(c.Param("role_id"))
	err := Iroles.FindbyID(IrolesID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "Role ID tidak ditemukan")
	}

	return c.JSON(http.StatusOK, Iroles)
}
