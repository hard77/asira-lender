package permission

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// type (
// 	Method string
// 	Url    string
// )

func ValidatePermissions(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		permissions := claims["permissions"]

		Method := c.Request.Method
		Url := c.Request.URL.String()

		log.Println(c.Request().Method)

		return next(c)
		// return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("%s", "invalid token"))
	}
}
