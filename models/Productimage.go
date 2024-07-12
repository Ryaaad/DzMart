package models

import interfaces "DzMart/interface"

type ProductImage struct {
	interfaces.BaseImage
	Product string `gorm:"type:varchar(255)"`
}
