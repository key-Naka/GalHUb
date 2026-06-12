package query

import "time"

type ListGame struct {
	Page int
	Size int

	Keyword string

	TagID *uint64

	CompanyID *uint64

	Status *int8
}

type ListGameItem struct {
	ID uint64 `json:"id"`

	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`

	Cover       string `json:"cover"`
	Description string `json:"description"`

	ReleaseDate *time.Time `json:"release_date"`

	ViewCount     int `json:"view_count"`
	FavoriteCount int `json:"favorite_count"`

	Status int8 `json:"status"`

	Tags []string `json:"tags"`

	Companies []string `json:"companies"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListGameResult struct {
	List []*ListGameItem `json:"list"`

	Total int64 `json:"total"`

	Page int `json:"page"`
	Size int `json:"size"`
}
