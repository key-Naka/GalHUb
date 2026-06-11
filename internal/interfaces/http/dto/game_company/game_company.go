package game_company

type ReplaceGameCompaniesRequest struct {
	CompanyIDs []uint64 `json:"company_ids" binding:"required"`
}
