package permission

import (
	"asira_lender/asira"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//ValidatePermissions handlers middleware
func ValidatePermissions(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		permissions := claims["permissions"]

		Method := c.Request().Method
		URL := c.Request().URL.String()

		perConfig := asira.App.Permission.GetStringMap(fmt.Sprintf("%s", Method))
		for key, value := range perConfig {
			for _, val := range permissions {
				if key == val; value == URL {
					return next(c)
				}
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("%s", "you are not allowed"))
	}
}
