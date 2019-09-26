package models

import (
	"gitlab.com/asira-ayannah/basemodel"
)

type (
	UserRelation struct {
		basemodel.BaseModel
		UserID uint64 `json:"user_id" gorm:"column:user_id;foreignkey"`
		BankID uint64 `json:"bank_id" gorm:"column:bank_id;foreignkey"`
	}
)

// gorm callback hook
func (u *UserRelation) BeforeCreate() (err error) {
	return nil
}

func (u *UserRelation) Create() error {
	err := basemodel.Create(&u)
	return err
}

// gorm callback hook
func (u *UserRelation) BeforeSave() (err error) {
	return nil
}

func (u *UserRelation) Save() error {
	err := basemodel.Save(&u)
	return err
}

func (u *UserRelation) FindbyID(id int) error {
	err := basemodel.FindbyID(&u, id)
	return err
}
