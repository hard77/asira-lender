package models

import (
	"time"

	"gitlab.com/asira-ayannah/basemodel"
)

type (
	Product struct {
		basemodel.BaseModel
		DeletedTime time.Time `json:"deleted_time" gorm:"column:deleted_time"`
		Name        string    `json:"name" gorm:"column:name;type:varchar(255)"`
		ServiceID   uint64    `json:"service_id" gorm:"column:service_id`
		Status      string    `json:"status" gorm:"column:status;type:varchar(255)"`
	}
)

func (model *Product) Create() error {
	err := basemodel.Create(&model)
	return err
}

func (model *Product) Save() error {
	err := basemodel.Save(&model)
	return err
}

func (model *Product) Delete() error {
	err := basemodel.Delete(&model)
	return err
}

func (model *Product) FindbyID(id int) error {
	err := basemodel.FindbyID(&model, id)
	return err
}

func (model *Product) PagedFindFilter(page int, rows int, order []string, sort []string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	products := []Product{}
	result, err = basemodel.PagedFindFilter(&products, page, rows, order, sort, filter)

	return result, err
}
