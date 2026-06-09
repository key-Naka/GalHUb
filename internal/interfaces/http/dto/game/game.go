package game

type CreateGameRequest struct {
	Title         string `json:"title" binding:"required,max=255"`
	OriginalTitle string `json:"original_title"`

	Cover string `json:"cover"`

	Description string `json:"description"`

	ReleaseDate string `json:"release_date"`
}
