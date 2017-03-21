package controllers

import (
	"github.com/gin-gonic/gin"
)

func Start(e *gin.Engine) {
	homeController(e)
}
