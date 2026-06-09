package command

import (
	"time"
)

type UpdateGame struct {
	ID uint64

	Title string

	OriginalTitle string

	Cover string

	Description string

	ReleaseDate *time.Time

	Status int8
}
