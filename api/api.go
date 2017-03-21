package api

import (
	"github.com/gin-gonic/gin"
)

func Start(e *gin.Engine) {

	api := e.Group("/api")

	v1 := api.Group("/v1")

	usersApi(v1)

}
