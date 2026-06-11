package tag

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTagRequest struct {
	Name string `json:"name" binding:"required"`
}
