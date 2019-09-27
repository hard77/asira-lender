package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"

	"gitlab.com/asira-ayannah/basemodel"
)

type (
	BankProduct struct {
		basemodel.BaseModel
		DeletedTime     time.Time      `json:"deleted_time" gorm:"column:deleted_time"`
		ProductID       uint64         `json:"product_id" gorm:"column:product_id"`
		BankServiceID   uint64         `json:"bank_service_id" gorm:"column:bank_service_id`
		MinTimeSpan     int            `json:"min_timespan" gorm:"column:min_timespan"`
		MaxTimeSpan     int            `json:"max_timespan" gorm:"column:max_timespan"`
		Interest        float64        `json:"interest" gorm:"column:interest"`
		MinLoan         int            `json:"min_loan" gorm:"column:min_loan"`
		MaxLoan         int            `json:"max_loan" gorm:"column:max_loan"`
		Fees            postgres.Jsonb `json:"fees" gorm:"column:fees"`
		Collaterals     pq.StringArray `json:"collaterals" gorm:"column:collaterals"`
		FinancingSector pq.StringArray `json:"financing_sector" gorm:"column:financing_sector"`
		Assurance       string         `json:"assurance" gorm:"column:assurance"`
		Status          string         `json:"status" gorm:"column:status"`
	}
)

func (model *BankProduct) Create() error {
	err := basemodel.Create(&model)
	if err != nil {
		return err
	}

	// err = KafkaSubmitModel(model, "bank_service_product")

	return err
}

func (model *BankProduct) Save() error {
	err := basemodel.Save(&model)
	if err != nil {
		return err
	}

	// err = KafkaSubmitModel(model, "bank_service_product")

	return err
}

func (model *BankProduct) Delete() error {
	err := basemodel.Delete(&model)
	if err != nil {
		return err
	}

	// err = KafkaSubmitModel(model, "bank_service_product_delete")

	return err
}

func (model *BankProduct) FindbyID(id int) error {
	err := basemodel.FindbyID(&model, id)
	return err
}

func (model *BankProduct) PagedFindFilter(page int, rows int, order []string, sort []string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	products := []BankProduct{}
	result, err = basemodel.PagedFindFilter(&products, page, rows, order, sort, filter)

	return result, err
}
