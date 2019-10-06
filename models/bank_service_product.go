package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"

	"gitlab.com/asira-ayannah/basemodel"
)

type (
	ServiceProduct struct {
		basemodel.BaseModel
		Name            string         `json:"name" gorm:"column:name"`
		MinTimeSpan     int            `json:"min_timespan" gorm:"column:min_timespan"`
		MaxTimeSpan     int            `json:"max_timespan" gorm:"column:max_timespan"`
		Interest        float64        `json:"interest" gorm:"column:interest"`
		MinLoan         int            `json:"min_loan" gorm:"column:min_loan"`
		MaxLoan         int            `json:"max_loan" gorm:"column:max_loan"`
		Fees            postgres.Jsonb `json:"fees" gorm:"column:fees"`
		ASN_Fee         string         `json:"asn_fee" gorm:"column:asn_fee"`
		Service         int            `json:"service" gorm:"column:service"`
		Collaterals     postgres.Jsonb `json:"collaterals" gorm:"column:collaterals"`
		FinancingSector postgres.Jsonb `json:"financing_sector" gorm:"column:financing_sector"`
		Assurance       string         `json:"assurance" gorm:"column:assurance"`
		Status          string         `json:"status" gorm:"column:status"`
	}
)

func (p *ServiceProduct) Create() error {
	err := basemodel.Create(&p)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(p, "bank_service_product")

	return err
}

func (p *ServiceProduct) Save() error {
	err := basemodel.Save(&p)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(p, "bank_service_product")

	return err
}

func (p *ServiceProduct) Delete() error {
	err := basemodel.Delete(&p)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(p, "bank_service_product_delete")

	return err
}

func (p *ServiceProduct) FindbyID(id int) error {
	err := basemodel.FindbyID(&p, id)
	return err
}

func (p *ServiceProduct) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	product := []ServiceProduct{}
	order := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&product, page, rows, order, sorts, filter)

	return result, err
}
