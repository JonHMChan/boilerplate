package middlewares

import (
	. "github.com/jonhmchan/boilerplate/core"

	"github.com/gin-gonic/gin"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(404, "/system/404.html", Context(c))
	}
}
