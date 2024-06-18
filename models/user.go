package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"not null" json:"Name" binding:"required,min=3,max=32"`
	Email    string `gorm:"uniqueIndex;not null" json:"Email" binding:"required,email"`
	Credit   int    `gorm:"default:0" json:"Credit" binding:"gte=0,lte=100"`
	Password string `json:"Password" binding:"required,min=8"`
	Sub      bool   `json:"sub"`
}
