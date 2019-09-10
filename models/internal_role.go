package models

type (
	Internal_Roles struct {
		BaseModel
		Name        string `json:"name" gorm:"column:name"`
		Description string `json:"description" gorm:"column:description"`
		Status      bool   `json:"status" gorm:"column:status;type:boolean" sql:"DEFAULT:FALSE"`
		System      string `json:"system" gorm:"column:system"`
	}
)

func (b *Internal_Roles) Create() (*Internal_Roles, error) {
	err := Create(&b)
	return b, err
}

func (b *Internal_Roles) Save() (*Internal_Roles, error) {
	err := Save(&b)
	return b, err
}

func (b *Internal_Roles) Delete() (*Internal_Roles, error) {
	err := Delete(&b)
	return b, err
}

func (b *Internal_Roles) FindbyID(id int) (*Internal_Roles, error) {
	err := FindbyID(&b, id)
	return b, err
}
