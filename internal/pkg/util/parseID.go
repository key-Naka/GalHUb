package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseID(
	c *gin.Context,
) (uint64, error) {

	return strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
}

