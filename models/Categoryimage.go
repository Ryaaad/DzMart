package models

import interfaces "DzMart/interface"

type CategoryImage struct {
	interfaces.BaseImage
	CatName string `gorm:"type:varchar(255)"`
}
