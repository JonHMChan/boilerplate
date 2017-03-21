package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robvdl/pongo2gin"
)

func Start(e *gin.Engine) {
	e.HTMLRender = pongo2gin.Default()

	homeController(e)

	staticController(e)
}
