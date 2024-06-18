package dtos

type UpdateCategory struct {
	Name *string `json:"name" binding:"required"`
}
