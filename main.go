package main

import (
	"net/http"

	"github.com/jonhmchan/boilerplate/api"
	"github.com/jonhmchan/boilerplate/controllers"
	. "github.com/jonhmchan/boilerplate/core"
	. "github.com/jonhmchan/boilerplate/middlewares"

	"github.com/drone/gin-location"
	"github.com/gin-gonic/gin"
)

func main() {
	Config.Start()
	Storage.Start()

	e := gin.Default()
	e.Use(location.Default())

	// Csrf
	e.Use(CsrfMiddleware())

	// Security
	e.Use(SecurityHeaders())

	api.Start(e)
	controllers.Start(e)

	e.NoRoute(NotFound())

	http.ListenAndServe(":"+Config.Env.PORT, CsrfHandler(e))
}
