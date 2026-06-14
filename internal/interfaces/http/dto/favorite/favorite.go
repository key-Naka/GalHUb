package favorite

type CreateFavoriteRequest struct {
	GameID uint64 `json:"game_id" binding:"required"`
}

type UpdateFavoriteRequest struct {
	GameID uint64 `json:"game_id" binding:"required"`
}
