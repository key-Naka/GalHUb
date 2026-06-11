package company

type CreateCompanyRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCompanyRequest struct {
	Name string `json:"name" binding:"required"`
}
