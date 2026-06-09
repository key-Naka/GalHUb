package command

import "time"

type CreateGame struct {
	Title         string
	OriginalTitle string

	Cover string

	Description string

	ReleaseDate *time.Time
}
