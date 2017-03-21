package main

import (
	"github.com/jonhmchan/boilerplate/api"
	"github.com/jonhmchan/boilerplate/controllers"
	. "github.com/jonhmchan/boilerplate/core"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	Config.Start()
	Storage.Start()
	api.Start(e)
	controllers.Start(e)

	e.Run(":8080")
}
