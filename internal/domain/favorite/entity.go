package favorite

import "time"

type Favorite struct {
	ID uint64

	UserID uint64
	GameID uint64

	CreatedAt time.Time
}
