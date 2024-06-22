package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"not null" json:"Name" binding:"required,min=1,max=30"`
	Email    string `gorm:"uniqueIndex;not null" json:"Email" binding:"required,email"`
	Credit   int    `gorm:"default:0" json:"Credit" binding:"gte=0"`
	Password string `json:"Password" binding:"required,min=8"`
	Sub      bool   `json:"sub" binding:"boolean" `
}
