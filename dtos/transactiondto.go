package dtos

type CreateTransactionInput struct {
	UserID    *uint `json:"UserID" binding:"required"`
	ProductID *uint `json:"ProductID" binding:"required"`
	Qte       *int  `json:"Qte" binding:"required,gte=0" `
}

type UpdateTransactionInput struct {
	Qte *int `json:"Qte" binding:"required,gte=0" `
}
