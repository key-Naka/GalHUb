package utill

import (
	"time"
)

func ParseReleaseDate(
	date string,
) (*time.Time, error) {

	if date == "" {
		return nil, nil
	}

	t, err := time.Parse(
		"2006-01-02",
		date,
	)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
