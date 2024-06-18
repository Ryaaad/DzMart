package models

type Category struct {
	IDCategory uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"uniqueIndex;not null" json:"Name" binding:"required,min=1,max=30"`
}
