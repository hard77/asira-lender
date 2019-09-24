package models

import "gitlab.com/asira-ayannah/basemodel"

type (
	BankService struct {
		basemodel.BaseModel
		Name    string `json:"name" gorm:"column:name"`
		ImageID int    `json:"image_id" gorm:"column:image_id"`
		Status  string `json:"status" gorm:"column:status"`
	}
)

func (b *BankService) Create() error {
	err := basemodel.Create(&b)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(b, "bank_service")

	return err
}

func (b *BankService) Save() error {
	err := basemodel.Save(&b)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(b, "bank_service")

	return err
}

func (b *BankService) Delete() error {
	err := basemodel.Delete(&b)
	if err != nil {
		return err
	}

	err = KafkaSubmitModel(b, "bank_service_delete")

	return err
}

func (b *BankService) FindbyID(id int) error {
	err := basemodel.FindbyID(&b, id)
	return err
}

func (b *BankService) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	bank_type := []BankService{}

	order := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&bank_type, page, rows, order, sorts, filter)

	return result, err
}
