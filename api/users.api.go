package api

import (
	. "github.com/jonhmchan/boilerplate/core"
	. "github.com/jonhmchan/boilerplate/models"

	"github.com/gin-gonic/gin"
)

func usersApi(r *gin.RouterGroup) {
	users := r.Group("/users")

	users.Use()
	{
		users.GET("", getAll)
	}
}

func getAll(c *gin.Context) {
	var users []User

	Storage.DB.Find(&users)

	c.JSON(200, users)
}
