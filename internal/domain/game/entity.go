package game

import "time"

type Game struct {
	ID uint64

	Title         string
	OriginalTitle string

	Cover       string
	Description string

	ReleaseDate *time.Time

	ViewCount     int
	FavoriteCount int

	Status int8

	CreatedAt time.Time
	UpdatedAt time.Time
}
