package models

type Internals struct {
	BaseModel
	Name   string `json:"name" gorm:"column:name"`
	Secret string `json:"secret" gorm:"column:secret"`
	Key    string `json:"key" gorm:"column:key"`
	Role   string `json:"role" gorm:"column:role"`
}

func (i *Internals) Create() (*Internals, error) {
	err := Create(&i)
	return i, err
}

func (i *Internals) Save() (*Internals, error) {
	err := Save(&i)
	return i, err
}

func (i *Internals) Delete() (*Internals, error) {
	err := Delete(&i)
	return i, err
}

func (l *Internals) FilterSearchSingle(filter interface{}) (*Internals, error) {
	err := FilterSearchSingle(&l, filter)
	return l, err
}
