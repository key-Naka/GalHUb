package game

import "time"

type GameModel struct {
	ID uint64 `gorm:"primaryKey"`

	Title         string
	OriginalTitle string

	Cover string

	Description string

	ReleaseDate *time.Time

	ViewCount     int
	FavoriteCount int

	Status int8

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (GameModel) TableName() string {
	return "games"
}
