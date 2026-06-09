package middleware

import (
	jwtpkg "galhub/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func Auth(secret string) gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "未授权",
			})
			return
		}
		claims, err := jwtpkg.ParseToken(
			tokenString,
			secret,
		)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "无效的token",
			})
			return
		}

		c.Set(
			"user_id",
			claims.UserID,
		)

		c.Next()
	}
}
