package models

type (
	InternalRoles struct {
		BaseModel
		Name        string `json:"name" gorm:"column:name"`
		Description string `json:"description" gorm:"column:description"`
		Status      bool   `json:"status" gorm:"column:status;type:boolean" sql:"DEFAULT:FALSE"`
		System      string `json:"system" gorm:"column:system"`
	}
)

func (b *InternalRoles) Create() (*InternalRoles, error) {
	err := Create(&b)
	return b, err
}

func (b *InternalRoles) Save() (*InternalRoles, error) {
	err := Save(&b)
	return b, err
}

func (b *InternalRoles) Delete() (*InternalRoles, error) {
	err := Delete(&b)
	return b, err
}

func (b *InternalRoles) FindbyID(id int) (*InternalRoles, error) {
	err := FindbyID(&b, id)
	return b, err
}
