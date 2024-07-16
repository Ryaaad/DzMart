package dtos

type LoginInput struct {
	Email    *string `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required"`
}
type UpdateUserInput struct {
	Name     *string `json:"name,omitempty" binding:"omitempty"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Password *string `json:"password,omitempty" binding:"omitempty"`
}

type Favorite struct {
	IDProduct *uint `json:"IDProduct" `
}
