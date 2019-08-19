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
	g.GET("/bank", admin_handlers.BankList)
	g.POST("/bank", admin_handlers.BankNew)
	g.GET("/bank/:bank_id", admin_handlers.BankDetail)
	g.PATCH("/bank/:bank_id", admin_handlers.BankPatch)
	g.DELETE("/bank/:bank_id", admin_handlers.BankDelete)

	// Bank Services
	g.GET("/bank_services", admin_handlers.BankServiceList)
	g.POST("/bank_services", admin_handlers.BankServiceNew)
	g.GET("/bank_services/:bank_service_id", admin_handlers.BankServiceDetail)
	g.PATCH("/bank_services/:bank_service_id", admin_handlers.BankServicePatch)
	g.DELETE("/bank_services/:bank_service_id", admin_handlers.BankServiceDelete)
}
