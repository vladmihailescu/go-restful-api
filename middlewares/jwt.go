package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladmihailescu/go-restful-api/utils"
)

var (
	BEARER_SCHEMA = "Bearer "
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(BEARER_SCHEMA) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		id, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userId", id)
		c.Next()
	}
}
