package dtos

type UpdateProductInput struct {
	Name        *string `json:"name,omitempty" binding:"omitempty"`
	Description *string `json:"Description,omitempty" binding:"omitempty"`
	Price       *int    `json:"price,omitempty" binding:"omitempty,gte=0"`
	Sold        *int    `json:"sold,omitempty" binding:"omitempty,gte=0,lte=100" `
	Qte         *int    `json:"Qte,omitempty" binding:"omitempty,gte=0" `
	Category    *string `json:"Category,omitempty"  binding:"omitempty"`
}
