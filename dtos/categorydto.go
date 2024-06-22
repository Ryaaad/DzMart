package dtos

type UpdateCategory struct {
	CatName *string `json:"CatName" binding:"required"`
}
