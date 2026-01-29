package request

type CreateTagInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTagInput struct {
	Name string `json:"name" binding:"required"`
}
