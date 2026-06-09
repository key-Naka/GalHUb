package game

type CreateGameRequest struct {
	Title         string `json:"title" binding:"required,max=255"`
	OriginalTitle string `json:"original_title"`

	Cover string `json:"cover"`

	Description string `json:"description"`

	ReleaseDate string `json:"release_date"`
}
type UpdateGameRequest struct {
	ID uint64 `json:"id"`

	Title         string `json:"title" binding:"required,max=255"`
	OriginalTitle string `json:"original_title" binding:"required,max=255"`
	Cover         string `json:"cover"`
	Description   string `json:"description"`
	ReleaseDate   string `json:"release_date"`
	Status        int8   `json:"status"`
}
