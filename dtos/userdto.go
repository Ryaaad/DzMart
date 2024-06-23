package dtos

type UpdateUserInput struct {
	Name     *string `json:"name,omitempty" binding:"omitempty"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Password *string `json:"password,omitempty" binding:"omitempty"`
}

type Favorite struct {
	IDProduct *uint `json:"IDProduct" `
}
