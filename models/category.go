package models

type Category struct {
	CatName string    `gorm:"primaryKey;column:CatName" json:"CatName" binding:"required,min=1,max=30"`
	Items   []Product `gorm:"foreignKey:Category" json:"Items" `
}
