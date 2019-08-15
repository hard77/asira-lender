package models

import "time"

type (
	BankType struct {
		BaseModel
		DeletedTime time.Time `json:"deleted_time" gorm:"column:deleted_time" sql:"DEFAULT:current_timestamp"`
		Name        string    `json:"name" gorm:"name"`
	}
)

func (b *BankType) Create() (*BankType, error) {
	err := Create(&b)
	return b, err
}

func (b *BankType) Save() (*BankType, error) {
	err := Save(&b)
	return b, err
}

func (l *Loan) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	loans := []Loan{}
	result, err = PagedFilterSearch(&loans, page, rows, orderby, sort, filter)

	return result, err
}
