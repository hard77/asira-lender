package reports

import (
	"asira_lender/asira"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gitlab.com/asira-ayannah/basemodel"

	"github.com/labstack/echo"
)

func ConvenienceFeeReport(c echo.Context) error {
	defer c.Request().Body.Close()

	db := asira.App.DB

	type ConvenienceFeeReport struct {
		BankName       string    `json:"bank_name"`
		ServiceName    string    `json:"service_name"`
		LoanID         string    `json:"loan_id"`
		CreatedTime    time.Time `json:"created_time"`
		Plafond        float64   `json:"plafond"`
		ConvenienceFee float64   `json:"convenience_fee"`
	}
	var results []ConvenienceFeeReport
	var totalRows int

	// pagination parameters
	rows, _ := strconv.Atoi(c.QueryParam("rows"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	if rows <= 0 {
		rows = 25
	}
	offset := (page * rows) - rows

	db = db.Table("loans l").
		Select("b.name as bank_name, ss.name as service_name, l.id as loan_id, l.created_time, loan_amount as plafond, value->>'amount' as convenience_fee").
		Joins("JOIN LATERAL jsonb_array_elements(l.fees) j ON true").
		Joins("INNER JOIN banks b ON b.id = l.bank").
		Joins("INNER JOIN bank_products p ON p.id = l.product").
		Joins("INNER JOIN bank_services s ON s.id = p.bank_service_id").
		Joins("INNER JOIN services ss ON ss.id = s.service_id").
		Where("LOWER(value->>'description') LIKE 'convenience%'")

	// filters
	if bankName := c.QueryParam("bank_name"); len(bankName) > 0 {
		db = db.Where("LOWER(b.name) LIKE ?", "%"+strings.ToLower(bankName)+"%")
	}
	if serviceName := c.QueryParam("service_name"); len(serviceName) > 0 {
		db = db.Where("LOWER(ss.name) LIKE ?", "%"+strings.ToLower(serviceName)+"%")
	}
	if loanID := c.QueryParam("loan_id"); len(loanID) > 0 {
		db = db.Where("l.id = ?", loanID)
	}
	if plafond := c.QueryParam("plafond"); len(plafond) > 0 {
		db = db.Where("loan_amount = ?", plafond)
	}
	if convenienceFee := c.QueryParam("convenience_fee"); len(convenienceFee) > 0 {
		db = db.Where("value->>'amount' = ?", convenienceFee)
	}
	if startDate := c.QueryParam("start_date"); len(startDate) > 0 {
		endDate := c.QueryParam("end_date")
		if len(endDate) < 1 {
			endDate = startDate
		}
		db = db.Where("l.created_time BETWEEN ? AND ?", startDate, endDate)
	}

	err := db.Limit(rows).Offset(offset).
		Find(&results).Count(&totalRows).Error
	if err != nil {
		log.Println(err)
	}

	lastPage := int(math.Ceil(float64(totalRows) / float64(rows)))

	response := basemodel.PagedFindResult{
		TotalData:   totalRows,
		Rows:        rows,
		CurrentPage: page,
		LastPage:    lastPage,
		From:        offset + 1,
		To:          offset + rows,
		Data:        results,
	}

	return c.JSON(http.StatusOK, response)
}
