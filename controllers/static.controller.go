package controllers

import (
	. "github.com/jonhmchan/boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func staticController(e *gin.Engine) {
	router := e.Group("/static")
	router.Use(StaticHeaders())
	router.Static("/", "static")

	main := e.Group("/")
	main.Use(StaticHeaders())

	main.StaticFile("/favicon.ico", "./static/img/shared/favicon/favicon.ico")
	main.StaticFile("/sitemap.xml", "./static/xml/sitemap.xml")
	main.StaticFile("/robots.txt", "./static/txt/robots.txt")
}
