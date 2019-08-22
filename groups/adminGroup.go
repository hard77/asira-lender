package groups

import (
	"asira_lender/admin_handlers"
	"asira_lender/handlers"
	"asira_lender/middlewares"

	"github.com/labstack/echo"
)

func AdminGroup(e *echo.Echo) {
	g := e.Group("/admin")
	middlewares.SetClientJWTmiddlewares(g, "admin")

	// OTP
	g.GET("/info", handlers.AsiraAppInfo)

	// Bank Types
	g.GET("/bank_types", admin_handlers.BankTypeList)
	g.POST("/bank_types", admin_handlers.BankTypeNew)
	g.GET("/bank_types/:bank_id", admin_handlers.BankTypeDetail)
	g.PATCH("/bank_types/:bank_id", admin_handlers.BankTypePatch)
	g.DELETE("/bank_types/:bank_id", admin_handlers.BankTypeDelete)

	// Banks
	g.GET("/banks", admin_handlers.BankList)
	g.POST("/banks", admin_handlers.BankNew)
	g.GET("/banks/:bank_id", admin_handlers.BankDetail)
	g.PATCH("/banks/:bank_id", admin_handlers.BankPatch)
	g.DELETE("/banks/:bank_id", admin_handlers.BankDelete)
}
