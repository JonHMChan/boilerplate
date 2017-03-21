package middlewares

import (
	"github.com/gin-gonic/gin"
)

func StaticHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Vary", "Accept-Encoding")
		c.Next()
	}
}
