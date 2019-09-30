package models

import (
	"gitlab.com/asira-ayannah/basemodel"
)

type (
	Roles struct {
		basemodel.BaseModel
		Name        string `json:"name" gorm:"column:name"`
		Description string `json:"description" gorm:"column:description"`
		System      string `json:"system" gorm:"column:system"`
		Status      bool   `json:"status" gorm:"column:status;type:boolean" sql:"DEFAULT:TRUE"`
	}
)

func (b *Roles) Create() error {
	err := basemodel.Create(&b)
	return err
}

func (b *Roles) Save() error {
	err := basemodel.Save(&b)
	return err
}

func (b *Roles) Delete() error {
	err := basemodel.Delete(&b)
	return err
}

func (b *Roles) FindbyID(id int) error {
	err := basemodel.FindbyID(&b, id)
	return err
}

func (b *Roles) FilterSearchSingle(filter interface{}) error {
	err := basemodel.SingleFindFilter(&b, filter)
	return err
}

func (b *Roles) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	internal := []Roles{}
	order := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&internal, page, rows, order, sorts, filter)

	return result, err
}
