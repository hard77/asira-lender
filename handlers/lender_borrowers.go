package handlers

import (
	"asira_lender/models"
	"database/sql"
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func LenderBorrowerList(c echo.Context) error {
	defer c.Request().Body.Close()

	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	lenderID, _ := strconv.Atoi(claims["jti"].(string))

	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	// filters
	fullname := c.QueryParam("fullname")
	status := c.QueryParam("status")
	id := c.QueryParam("id")

	type Filter struct {
		Bank     sql.NullInt64 `json:"bank"`
		Fullname string        `json:"fullname"`
		Status   string        `json:"status"`
		ID       string        `json:"id"`
	}

	borrower := models.Borrower{}
	result, err := borrower.PagedFilterSearch(page, rows, orderby, sort, &Filter{
		Bank: sql.NullInt64{
			Int64: int64(lenderID),
			Valid: true,
		},
		Fullname: fullname,
		Status:   status,
		ID:       id,
	})

	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "query result error")
	}

	return c.JSON(http.StatusOK, result)
}

func LenderBorrowerListDetail(c echo.Context) error {
	defer c.Request().Body.Close()

	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	lenderID, _ := strconv.Atoi(claims["jti"].(string))

	borrower_id, err := strconv.Atoi(c.Param("borrower_id"))
	if err != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, err, "error parsing borrower id")
	}
	type Filter struct {
		Bank sql.NullInt64 `json:"bank"`
		ID   int           `json:"id"`
	}

	borrower := models.Borrower{}
	result, err := borrower.FilterSearchSingle(&Filter{
		Bank: sql.NullInt64{
			Int64: int64(lenderID),
			Valid: true,
		},
		ID: borrower_id,
	})

	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "query result error")
	}

	return c.JSON(http.StatusOK, result)
}

func LenderBorrowerListDownload(c echo.Context) error {
	defer c.Request().Body.Close()

	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	lenderID, _ := strconv.Atoi(claims["jti"].(string))

	// pagination parameters
	rows, _ := strconv.Atoi(c.QueryParam("rows"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	// filters
	fullname := c.QueryParam("fullname")
	status := c.QueryParam("status")
	id := c.QueryParam("id")

	type Filter struct {
		Bank     sql.NullInt64 `json:"bank"`
		Fullname string        `json:"fullname"`
		Status   string        `json:"status"`
		ID       string        `json:"id"`
	}

	borrower := models.Borrower{}
	result, _ := borrower.PagedFilterSearch(page, rows, orderby, sort, &Filter{
		Bank: sql.NullInt64{
			Int64: int64(lenderID),
			Valid: true,
		},
		Fullname: fullname,
		Status:   status,
		ID:       id,
	})

	// write results to a new csv
	outfile, _ := os.Create("files/csvdownload.csv")
	defer outfile.Close()

	writer := csv.NewWriter(outfile)
	defer writer.Flush()
	writer.Write([]string{"fullname", "gender", "idcard_number", "idcard_imageid", "taxid_number", "taxid_imageid", "email", "birthday", "birthplace", "last_education", "mother_name", "phone", "married_status", "spouse_name", "spouse_birthday", "spouse_lasteducation", "dependants", "address", "province", "city", "neighbour_association", "hamlets", "home_phonenumber", "subdistrict", "urban_village", "home_ownership", "lived_for", "occupation", "employee_id", "employer_name", "employer_address", "department", "been_workingfor", "direct_superiorname", "employer_number", "monthly_income", "other_income", "other_incomesource", "field_of_work", "related_personname", "related_relation", "related_phonenumber", "related_homenumber", "bank", "bank_account_number"})

	for i, record := range result.Data.([]*models.Borrower) {
		// skip header row
		if i == 0 {
			writer.Write([]string{
				record.Fullname,
				record.Gender,
				record.IdCardNumber,
				record.IdCardImageID,
				record.TaxIDnumber,
				record.TaxIDImageID,
				record.Email,
				record.Birthday.Format("2006-01-02"),
				record.Birthplace,
				record.LastEducation,
				record.MotherName,
				record.Phone,
				record.MarriedStatus,
				record.SpouseName,
				record.SpouseBirthday.Format("2006-01-02"),
				record.SpouseLastEducation,
				string(record.Dependants),
				record.Address,
				record.Province,
				record.City,
				record.NeighbourAssociation,
				record.Hamlets,
				record.HomePhoneNumber,
				record.Subdistrict,
				record.UrbanVillage,
				record.HomeOwnership,
				string(record.LivedFor),
				record.Occupation,
				record.EmployeeID,
				record.EmployerName,
				record.EmployerAddress,
				record.Department,
				string(record.BeenWorkingFor),
				record.DirectSuperior,
				record.EmployerNumber,
				string(record.MonthlyIncome),
				string(record.OtherIncome),
				record.OtherIncomeSource,
				record.FieldOfWork,
				record.RelatedPersonName,
				record.RelatedRelation,
				record.RelatedPhoneNumber,
				record.RelatedHomePhone,
				record.RelatedAddress,
				record.BankAccountNumber,
			})
			continue
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Download",
	})
}
