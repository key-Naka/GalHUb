package query

import "time"

type DetailTag struct {
	ID uint64 `json:"id"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailCompany struct {
	ID uint64 `json:"id"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailGame struct {
	ID uint64 `json:"id"`

	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`

	Cover       string `json:"cover"`
	Description string `json:"description"`

	ReleaseDate *time.Time `json:"release_date"`

	ViewCount     int `json:"view_count"`
	FavoriteCount int `json:"favorite_count"`

	Status int8 `json:"status"`

	Tags      []*DetailTag     `json:"tags"`
	Companies []*DetailCompany `json:"companies"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
