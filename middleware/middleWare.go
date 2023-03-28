package middleware

import (
	"golang_gin_gorm_jwt/helpers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		var stat bool
		cookie, err := c.Cookie("user")

		if err != nil {
			c.Set("stat", "false")
		}

		stat = helpers.ValidateTokens(cookie)
		c.Set("stat", stat)

		c.Next()
	}
}
