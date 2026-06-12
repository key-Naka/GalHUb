package util

import (
	"strconv"
	"strings"
)

func ParseGameStatus(
	value string,
) (int8, error) {
	switch strings.ToLower(value) {
	case "draft":
		return 0, nil
	case "published":
		return 1, nil
	default:
		parsed, err := strconv.ParseInt(
			value,
			10,
			8,
		)
		if err != nil {
			return 0, err
		}

		return int8(parsed), nil
	}
}

