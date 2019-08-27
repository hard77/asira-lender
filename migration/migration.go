package migration

import (
	"asira_lender/asira"
	"asira_lender/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

func Seed() {
	seeder := asira.App.DB.Begin()
	defer seeder.Commit()

	if asira.App.ENV == "development" {
		// seed images
		file, _ := os.Open("migration/image_dummy.txt")
		defer file.Close()
		b64image, _ := ioutil.ReadAll(file)
		images := []models.Image{
			models.Image{
				Image_string: string(b64image),
			},
		}
		for _, image := range images {
			image.Create()
		}

		// seed bank types
		bankTypes := []models.BankType{
			models.BankType{
				Name:        "BPD",
				Description: "Layanan BPD",
			},
			models.BankType{
				Name:        "BPR",
				Description: "Layanan BPR",
			},
			models.BankType{
				Name:        "Koperasi",
				Description: "Layanan Koperasi",
			},
		}
		for _, bankType := range bankTypes {
			bankType.Create()
		}

		// seed bank services
		bankServices := []models.BankService{
			models.BankService{
				Name:    "Pinjaman PNS",
				ImageID: 1,
				Status:  "active",
			},
			models.BankService{
				Name:    "Pinjaman Pensiun",
				ImageID: 1,
				Status:  "active",
			},
			models.BankService{
				Name:    "Pinjaman UMKN",
				ImageID: 1,
				Status:  "active",
			},
			models.BankService{
				Name:    "Pinjaman Mikro",
				ImageID: 1,
				Status:  "inactive",
			},
			models.BankService{
				Name:    "Pinjaman Lainnya",
				ImageID: 1,
				Status:  "inactive",
			},
		}
		for _, bankService := range bankServices {
			bankService.Create()
		}

		// seed service products
		feesMarshal, _ := json.Marshal(map[string]interface{}{
			"description": "Admin Fee",
			"amount":      2500,
		})
		collateralMarshal, _ := json.Marshal([]string{"Surat Tanah", "BPKB"})
		financeMarshal, _ := json.Marshal([]string{"Pendidikan"})
		serviceProducts := []models.ServiceProduct{
			models.ServiceProduct{
				Name:            "Product A",
				MinTimeSpan:     1,
				MaxTimeSpan:     6,
				Interest:        5,
				MinLoan:         1000000,
				MaxLoan:         10000000,
				Fees:            postgres.Jsonb{feesMarshal},
				ASN_Fee:         "1%",
				Service:         1,
				Collaterals:     postgres.Jsonb{collateralMarshal},
				FinancingSector: postgres.Jsonb{financeMarshal},
				Assurance:       "an Assurance",
				Status:          "active",
			},
			models.ServiceProduct{
				Name:            "Product B",
				MinTimeSpan:     3,
				MaxTimeSpan:     12,
				Interest:        5,
				MinLoan:         5000000,
				MaxLoan:         8000000,
				Fees:            postgres.Jsonb{feesMarshal},
				ASN_Fee:         "1%",
				Service:         1,
				Collaterals:     postgres.Jsonb{collateralMarshal},
				FinancingSector: postgres.Jsonb{financeMarshal},
				Assurance:       "an Assurance",
				Status:          "active",
			},
		}
		for _, serviceProduct := range serviceProducts {
			serviceProduct.Create()
		}

		// seed lenders
		array := []int{1, 2, 3, 5, 8}
		jMarshal, _ := json.Marshal(array)
		lenders := []models.Bank{
			models.Bank{
				Name:     "Bank A",
				Type:     1,
				Address:  "Bank A Address",
				Province: "Province A",
				City:     "City A",
				Services: postgres.Jsonb{jMarshal},
				Products: postgres.Jsonb{jMarshal},
				PIC:      "Bank A PIC",
				Phone:    "081234567890",
				Username: "Banktoib",
				Password: "password",
			},
			models.Bank{
				Name:     "Bank B",
				Type:     2,
				Address:  "Bank B Address",
				Province: "Province B",
				City:     "City B",
				Services: postgres.Jsonb{jMarshal},
				Products: postgres.Jsonb{jMarshal},
				PIC:      "Bank B PIC",
				Phone:    "081234567891",
				Username: "Banktoic",
				Password: "password",
			},
		}
		for _, lender := range lenders {
			lender.Create()
		}

		// seed borrowers
		borrowers := []models.Borrower{
			models.Borrower{
				Status:               "active",
				Fullname:             "Full Name A",
				Gender:               "M",
				IdCardNumber:         "9876123451234567789",
				TaxIDnumber:          "0987654321234567890",
				IdCardImageID:        "1",
				TaxIDImageID:         "1",
				Email:                "emaila@domain.com",
				Birthday:             time.Now(),
				Birthplace:           "a birthplace",
				LastEducation:        "a last edu",
				MotherName:           "a mom",
				Phone:                "081234567890",
				MarriedStatus:        "single",
				SpouseName:           "a spouse",
				SpouseBirthday:       time.Now(),
				SpouseLastEducation:  "master",
				Dependants:           0,
				Address:              "a street address",
				Province:             "a province",
				City:                 "a city",
				NeighbourAssociation: "a rt",
				Hamlets:              "a rw",
				HomePhoneNumber:      "021837163",
				Subdistrict:          "a camat",
				UrbanVillage:         "a lurah",
				HomeOwnership:        "privately owned",
				LivedFor:             5,
				Occupation:           "accupation",
				EmployerName:         "amployer",
				EmployerAddress:      "amployer address",
				Department:           "a department",
				BeenWorkingFor:       2,
				DirectSuperior:       "a boss",
				EmployerNumber:       "02188776655",
				MonthlyIncome:        5000000,
				OtherIncome:          2000000,
				RelatedPersonName:    "a big sis",
				RelatedPhoneNumber:   "08987654321",
				BankAccountNumber:    "520384716",
				Bank: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
			},
			models.Borrower{
				Status:               "active",
				Fullname:             "Full Name B",
				Gender:               "F",
				IdCardNumber:         "9876123451234567781",
				TaxIDnumber:          "0987654321234567891",
				IdCardImageID:        "1",
				TaxIDImageID:         "1",
				Email:                "emailb@domain.com",
				Birthday:             time.Now(),
				Birthplace:           "b birthplace",
				LastEducation:        "b last edu",
				MotherName:           "b mom",
				Phone:                "081234567891",
				MarriedStatus:        "single",
				SpouseName:           "b spouse",
				SpouseBirthday:       time.Now(),
				SpouseLastEducation:  "master",
				Dependants:           0,
				Address:              "b street address",
				Province:             "b province",
				City:                 "b city",
				NeighbourAssociation: "b rt",
				Hamlets:              "b rw",
				HomePhoneNumber:      "021837163",
				Subdistrict:          "b camat",
				UrbanVillage:         "b lurah",
				HomeOwnership:        "privately owned",
				LivedFor:             5,
				Occupation:           "bccupation",
				EmployerName:         "bmployer",
				EmployerAddress:      "bmployer address",
				Department:           "b department",
				BeenWorkingFor:       2,
				DirectSuperior:       "b boss",
				EmployerNumber:       "02188776655",
				MonthlyIncome:        5000000,
				OtherIncome:          2000000,
				RelatedPersonName:    "b big sis",
				RelatedPhoneNumber:   "08987654321",
				RelatedAddress:       "big sis address",
				Bank: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
			},
			models.Borrower{
				Status:               "inactive",
				Fullname:             "Full Name C",
				Gender:               "F",
				IdCardNumber:         "9876123451234567789",
				TaxIDnumber:          "0987654321234567890",
				IdCardImageID:        "1",
				TaxIDImageID:         "1",
				Email:                "emailc@domain.com",
				Birthday:             time.Now(),
				Birthplace:           "c birthplace",
				LastEducation:        "c last edu",
				MotherName:           "c mom",
				Phone:                "081234567892",
				MarriedStatus:        "single",
				Dependants:           2,
				Address:              "c street address",
				Province:             "c province",
				City:                 "c city",
				NeighbourAssociation: "c rt",
				Hamlets:              "c rw",
				HomePhoneNumber:      "021837164",
				Subdistrict:          "c camat",
				UrbanVillage:         "c lurah",
				HomeOwnership:        "privately owned",
				LivedFor:             5,
				Occupation:           "house wife",
				RelatedPersonName:    "c big sis",
				RelatedPhoneNumber:   "08987654323",
				BankAccountNumber:    "520384718",
				Bank: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
			},
			models.Borrower{
				Status:               "inactive",
				Fullname:             "Full Name D",
				Gender:               "M",
				IdCardNumber:         "9876123451234567719",
				TaxIDnumber:          "0987654321234567890",
				IdCardImageID:        "1",
				TaxIDImageID:         "1",
				Email:                "emaild@domain.com",
				Birthday:             time.Now(),
				Birthplace:           "d birthplace",
				LastEducation:        "d last edu",
				MotherName:           "d mom",
				Phone:                "081234567895",
				MarriedStatus:        "single",
				SpouseName:           "d spouse",
				SpouseBirthday:       time.Now(),
				SpouseLastEducation:  "master",
				Dependants:           0,
				Address:              "d street address",
				Province:             "d province",
				City:                 "d city",
				NeighbourAssociation: "d rt",
				Hamlets:              "d rw",
				HomePhoneNumber:      "021837167",
				Subdistrict:          "d camat",
				UrbanVillage:         "d lurah",
				LivedFor:             8,
				Occupation:           "dccupation",
				EmployerName:         "dmployer",
				EmployerAddress:      "dmployer address",
				Department:           "d department",
				BeenWorkingFor:       2,
				DirectSuperior:       "d boss",
				EmployerNumber:       "021887766554",
				MonthlyIncome:        5000000,
				OtherIncome:          2000000,
				RelatedPersonName:    "d big sis",
				RelatedPhoneNumber:   "08987654321",
				BankAccountNumber:    "520384716",
				Bank: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
			},
			models.Borrower{
				Status:               "inactive",
				Fullname:             "Full Name E",
				Gender:               "M",
				IdCardNumber:         "9876123451234567784129",
				TaxIDnumber:          "0987654321234567890",
				IdCardImageID:        "1",
				TaxIDImageID:         "1",
				Email:                "emaile@domain.com",
				Birthday:             time.Now(),
				Birthplace:           "e birthplace",
				LastEducation:        "e last edu",
				MotherName:           "e mom",
				Phone:                "08123456789767",
				MarriedStatus:        "married",
				SpouseName:           "e spouse",
				SpouseBirthday:       time.Now(),
				SpouseLastEducation:  "master",
				Dependants:           0,
				Address:              "e street address",
				Province:             "e province",
				City:                 "e city",
				NeighbourAssociation: "e rt",
				Hamlets:              "e rw",
				HomePhoneNumber:      "0218371631",
				Subdistrict:          "e camat",
				UrbanVillage:         "e lurah",
				Occupation:           "eccupation",
				EmployerName:         "employer",
				EmployerAddress:      "employer address",
				Department:           "e department",
				BeenWorkingFor:       2,
				DirectSuperior:       "e boss",
				EmployerNumber:       "02188776655133",
				MonthlyIncome:        5000000,
				OtherIncome:          2000000,
				RelatedPersonName:    "e big sis",
				RelatedPhoneNumber:   "0895387654321",
				BankAccountNumber:    "520384716",
				Bank: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
			},
		}
		for _, borrower := range borrowers {
			borrower.Create()
		}

		// seed loans
		fees := []models.LoanFee{
			models.LoanFee{
				Description: "fee 1",
				Amount:      1000,
			},
		}
		jMarshal, _ = json.Marshal(fees)
		loans := []models.Loan{
			models.Loan{
				Bank: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				Owner: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				OwnerName:        "Full Name A",
				LoanAmount:       5000000,
				Installment:      8,
				LoanIntention:    "a loan 1 intention",
				IntentionDetails: "a loan 1 intention details",
				Fees:             postgres.Jsonb{jMarshal},
				Interest:         1.5,
				TotalLoan:        float64(6500000),
				LayawayPlan:      500000,
			},
			models.Loan{
				Bank: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				Owner: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				OwnerName:        "Full Name B",
				LoanAmount:       2000000,
				Installment:      3,
				LoanIntention:    "a loan 1 intention",
				IntentionDetails: "a loan 1 intention details",
				Fees:             postgres.Jsonb{jMarshal},
				Interest:         1.5,
				TotalLoan:        float64(3000000),
				LayawayPlan:      200000,
			},
			models.Loan{
				Bank: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				Owner: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				OwnerName:        "Full Name C",
				LoanAmount:       29000000,
				Installment:      3,
				LoanIntention:    "a loan 1 intention",
				IntentionDetails: "a loan 1 intention details",
				Fees:             postgres.Jsonb{jMarshal},
				Interest:         1.5,
				TotalLoan:        float64(6500000),
				LayawayPlan:      500000,
			},
			models.Loan{
				Bank: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				Owner: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				OwnerName:        "Full Name D",
				LoanAmount:       3000000,
				Installment:      3,
				LoanIntention:    "a loan 1 intention",
				IntentionDetails: "a loan 1 intention details",
				Fees:             postgres.Jsonb{jMarshal},
				Interest:         1.5,
				TotalLoan:        float64(3000000),
				LayawayPlan:      200000,
			},
			models.Loan{
				Bank: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				Owner: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				OwnerName:        "Full Name E",
				LoanAmount:       9123456,
				Installment:      3,
				LoanIntention:    "a loan 3 intention",
				IntentionDetails: "a loan 5 intention details",
				Fees:             postgres.Jsonb{jMarshal},
				Interest:         1.5,
				TotalLoan:        float64(3000000),
				LayawayPlan:      200000,
			},
			models.Loan{
				Bank: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				Owner: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				OwnerName:        "Full Name F",
				LoanAmount:       80123456,
				Installment:      11,
				LoanIntention:    "a loan 3 intention",
				IntentionDetails: "a loan 5 intention details",
				Fees:             postgres.Jsonb{jMarshal},
				Interest:         1.5,
				TotalLoan:        float64(3000000),
				LayawayPlan:      200000,
			},
		}
		for _, loan := range loans {
			loan.Create()
		}
	}
}

// truncate defined tables. []string{"all"} to truncate all tables.
func Truncate(tableList []string) (err error) {
	if len(tableList) > 0 {
		if tableList[0] == "all" {
			tableList = []string{
				"bank_types",
				"banks",
				"bank_services",
				"borrowers",
				"loans",
				"images",
			}
		}

		tables := strings.Join(tableList, ", ")
		sqlQuery := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tables)
		err = asira.App.DB.Exec(sqlQuery).Error
		return err
	}

	return fmt.Errorf("define tables that you want to truncate")
}
