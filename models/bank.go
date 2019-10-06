package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/asira-ayannah/basemodel"
	"golang.org/x/crypto/bcrypt"
)

type (
	Bank struct {
		basemodel.BaseModel
		DeletedTime         time.Time      `json:"deleted_time" gorm:"column:deleted_time" sql:"DEFAULT:current_timestamp"`
		Name                string         `json:"name" gorm:"column:name;type:varchar(255)"`
		Type                int            `json:"type" gorm:"column:type;type:varchar(255)"`
		Address             string         `json:"address" gorm:"column:address;type:text"`
		Province            string         `json:"province" gorm:"column:province;type:varchar(255)"`
		City                string         `json:"city" gorm:"column:city;type:varchar(255)"`
		AdminFeeSetup       string         `json:"adminfee_setup" gorm:"column:adminfee_setup;type:varchar(255)"`
		ConvinienceFeeSetup string         `json:"convfee_setup" gorm:"column:convfee_setup;type:varchar(255)"`
		Services            postgres.Jsonb `json:"services" gorm:"column:services;type:jsonb"`
		Products            postgres.Jsonb `json:"products" gorm:"column:products;type:jsonb"`
		PIC                 string         `json:"pic" gorm:"column:pic;type:varchar(255)"`
		Phone               string         `json:"phone" gorm:"column:phone;type:varchar(255)"`
		Username            string         `json:"username" gorm:"column:username;type:varchar(255);unique;not null"`
		Password            string         `json:"password" gorm:"column:password;type:text;not null"`
	}
)

// gorm callback hook
func (b *Bank) BeforeCreate() (err error) {
	log.Printf("new bank : %v", b)
	if len(b.Username) < 1 {
		b.Username = uuid.New().String()
	}
	if len(b.Username) < 1 {
		b.Password = uuid.New().String()
	}
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	b.Password = string(passwordByte)
	return nil
}

func (b *Bank) Create() error {
	err := basemodel.Create(&b)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(b, "bank")

	return err
}

// gorm callback hook
func (b *Bank) BeforeSave() (err error) {
	return nil
}

func (b *Bank) Save() error {
	err := basemodel.Save(&b)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(b, "bank")

	return err
}

func (b *Bank) Delete() error {
	err := basemodel.Delete(&b)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(b, "bank_delete")

	return err
}

func (b *Bank) FindbyID(id int) error {
	err := basemodel.FindbyID(&b, id)
	return err
}

func (b *Bank) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	bank_type := []Bank{}
	order := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&bank_type, page, rows, order, sorts, filter)

	return result, err
}
