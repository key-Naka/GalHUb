package query

import "time"

type ListFavorite struct {
	UserID uint64

	Page int
	Size int

	Keyword string
}

type ListFavoriteItem struct {
	ID uint64 `json:"id"`

	UserID uint64 `json:"user_id"`
	GameID uint64 `json:"game_id"`

	GameTitle         string     `json:"game_title"`
	GameOriginalTitle string     `json:"game_original_title"`
	GameCover         string     `json:"game_cover"`
	ReleaseDate       *time.Time `json:"release_date"`
	GameStatus        int8       `json:"game_status"`
	FavoriteCount     int        `json:"favorite_count"`
	Tags              []string   `json:"tags"`
	Companies         []string   `json:"companies"`

	CreatedAt time.Time `json:"created_at"`
}

type ListFavoriteResult struct {
	List []*ListFavoriteItem `json:"list"`

	Total int64 `json:"total"`

	Page int `json:"page"`
	Size int `json:"size"`
}
