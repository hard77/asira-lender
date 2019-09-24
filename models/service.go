package models

import "gitlab.com/asira-ayannah/basemodel"

type (
	Service struct {
		basemodel.BaseModel
		Name string `json:"name" gorm:"column:name;type:varchar(255)"`
	}
)
