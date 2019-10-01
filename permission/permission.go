package permission

import (
	"asira_lender/asira"
	"asira_lender/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//ValidatePermissions handlers middleware
func ValidatePermissions(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//get role_id from JWT
		user := c.Get("user")
		token := user.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		RoleID := claims["role_id"]

		//method and url from request
		Method := c.Request().Method
		URL := c.Request().URL.String()
		PermissionsModel := models.Permissions{}
		//check permissions
		perConfig := asira.App.Permission.GetStringMap(fmt.Sprintf("%s", Method))
		for key, value := range perConfig {
			if strings.Contains(URL, value.(string)) {
				if !asira.App.DB.Where("lower(permissions) = ? AND role_id = ?", key, RoleID).Find(&PermissionsModel).RecordNotFound() {
					return next(c)
				}
			}
		}

		return echo.NewHTTPError(http.StatusMethodNotAllowed, fmt.Sprintf("%s", "you are not allowed"))
	}
}
