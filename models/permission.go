package models

import (
	"gitlab.com/asira-ayannah/basemodel"
)

type (
	Permissions struct {
		basemodel.BaseModel
		RoleID      int    `json:"role_id" gorm:"column:role_id"`
		Permissions string `json:"permissions" gorm:"column:permissions;type:varchar(255)"`
	}
)

func (b *Permissions) Create() error {
	err := basemodel.Create(&b)
	return err
}

func (b *Permissions) Save() error {
	err := basemodel.Save(&b)
	return err
}

func (b *Permissions) Delete() error {
	err := basemodel.Delete(&b)
	return err
}

func (b *Permissions) FindbyID(id int) error {
	err := basemodel.FindbyID(&b, id)
	return err
}

func (b *Permissions) FilterSearchSingle(filter interface{}) error {
	err := basemodel.SingleFindFilter(&b, filter)
	return err
}

func (b *Permissions) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	internal := []Permissions{}
	order := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&internal, page, rows, order, sorts, filter)

	return result, err
}

func (b *Permissions) FilterSearch(limit int, offset int, orderby string, sort string, filter interface{}) (result interface{}, err error) {
	internal := []Permissions{}
	order := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.FindFilter(&internal, order, sorts, limit, offset, filter)

	return result, err
}
