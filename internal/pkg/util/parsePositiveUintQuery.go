package util

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParsePositiveUintQuery(
	c *gin.Context,
	key string,
) (*uint64, error) {
	value := c.Query(key)
	if value == "" {
		return nil, nil
	}

	parsed, err := strconv.ParseUint(
		value,
		10,
		64,
	)
	if err != nil || parsed < 1 {
		return nil, fmt.Errorf("invalid %s", key)
	}

	return &parsed, nil
}

