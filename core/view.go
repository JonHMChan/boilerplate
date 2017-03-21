package core

import (
	"os"

	. "github.com/jonhmchan/boilerplate/models"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

var View ViewStruct

type ViewStruct struct {
}

func (v *ViewStruct) Start() {
	pongo2.Globals["GOOGLE_ANALYTICS_ID"] = os.Getenv("GOOGLE_ANALYTICS_ID")
	pongo2.Globals["INTERCOM_ID"] = os.Getenv("INTERCOM_ID")
}

func (v *ViewStruct) AsPongo(val string) *pongo2.Value {
	return pongo2.AsValue(val)
}

func (v *ViewStruct) Img(path string) *pongo2.Value {
	return pongo2.AsValue("/static/img/" + path)
}

func (v *ViewStruct) Js(path string) *pongo2.Value {
	return pongo2.AsValue("/static/js/dist/" + path)
}

func (v *ViewStruct) Css(path string) *pongo2.Value {
	return pongo2.AsValue("/static/css/dist/" + path)
}

func Context(c *gin.Context, context ...interface{}) pongo2.Context {
	var user User
	if val, ok := c.Get("user"); ok {
		user = val.(User)
	}
	viewContext := pongo2.Context{
		"View":      &View,
		"User":      &user,
		"Request":   c.Request,
		"Settings":  Config.Settings,
		"CSRF":      nosurf.Token(c.Request),
		"Anonymous": user.IsAnonymous(),
	}
	if len(context) == 1 {
		viewContext["Model"] = context[0]
	}
	return viewContext
}
