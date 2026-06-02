package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		// TODO JWT验证

		c.Next()
	}
}
