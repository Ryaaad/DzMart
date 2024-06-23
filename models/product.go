package models

type Product struct {
	IDProduct   uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"not null" json:"Name" binding:"required,min=1,max=30"` // uniqueIndex
	Description string  `gorm:"not null" json:"Description" binding:"required,min=1"`
	Price       int     `json:"Price" binding:"gte=0"`
	Sold        int     `gorm:"default:0" json:"Sold" binding:"gte=0,lte=100" `
	Rating      float64 `gorm:"default:0" json:"Rating" binding:"gte=0,lte=5" `
	Qte         int     `json:"Qte" binding:"gte=0" `
	Category    string  `json:"Category"  binding:"required"`
	Users       []*User `gorm:"many2many:Favorite;onDelete:CASCADE"`
}
