package controllers

import (
	"github.com/gin-gonic/gin"
)

func homeController(e *gin.Engine) {
	r := e.Group("")

	r.Use()
	{
		r.GET("", homeIndex)
	}
}

func homeIndex(c *gin.Context) {
	c.String(200, "Hello world!")
}
