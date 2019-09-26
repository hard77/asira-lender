package models

import (
	"gitlab.com/asira-ayannah/basemodel"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		basemodel.BaseModel
		RoleID   int    `json:"role_id" gorm:"column:role_id"`
		Username string `json:"username" gorm:"column:username;type:varchar(255);unique;not null"`
		Password string `json:"password" gorm:"column:password;type:text;not null"`
	}
)

// gorm callback hook
func (u *User) BeforeCreate() (err error) {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordByte)
	return nil
}

func (u *User) Create() error {
	err := basemodel.Create(&u)
	return err
}

// gorm callback hook
func (u *User) BeforeSave() (err error) {
	return nil
}

func (u *User) Save() error {
	err := basemodel.Save(&u)
	return err
}

func (u *User) FindbyID(id int) error {
	err := basemodel.FindbyID(&u, id)
	return err
}
