package company

import "time"

type Company struct {
	ID uint64

	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}
