package util

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParsePageAndSize(
	c *gin.Context,
	maxSize int,
) (int, int, error) {
	page, err := strconv.Atoi(
		c.DefaultQuery("page", "1"),
	)
	if err != nil || page < 1 {
		return 0, 0, fmt.Errorf("invalid page")
	}

	size, err := strconv.Atoi(
		c.DefaultQuery("size", "10"),
	)
	if err != nil || size < 1 || size > maxSize {
		return 0, 0, fmt.Errorf("invalid size")
	}

	return page, size, nil
}
