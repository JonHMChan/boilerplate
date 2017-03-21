package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	. "github.com/jonhmchan/boilerplate/core"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

const (
	CSRF_FORM_VALUE  = "csrf_token"
	CSRF_HEADER_NAME = "X-CSRF-Token"
)

func CsrfHandler(e *gin.Engine) *nosurf.CSRFHandler {
	handler := nosurf.New(e)
	handler.SetFailureHandler(http.HandlerFunc(failureHandler))

	handler.ExemptRegexp(".*")

	return handler
}

func CsrfMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsStringInSlice([]string{"GET", "HEAD", "OPTIONS", "TRACE"}, c.Request.Method) {

			var possibleTokens = []string{
				c.Request.Header.Get(CSRF_HEADER_NAME),
				c.PostForm(CSRF_FORM_VALUE),
			}

			realToken := nosurf.Token(c.Request)

			for _, token := range possibleTokens {
				if verifyToken(realToken, token) {
					c.Next()
					return
				}
			}
			c.String(403, "CSRF failure")
			c.Abort()
			return

		}
		c.Next()
	}
}

func verifyToken(realToken string, token string) bool {
	if !nosurf.VerifyToken(realToken, token) {
		return false
	}
	return true
}

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

func failureHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	methodColor := colorForMethod(method)

	fmt.Fprintf(os.Stdout, "[GIN] %v |%s %3d %s| %13v | %s |%s  %s %-7s %s\n%s",
		time.Now().Format("2006/01/02 - 15:04:05"),
		yellow, 403, reset,
		0.0,
		"[::?]:?????",
		methodColor, reset, method,
		r.URL.String(),
		nosurf.Reason(r).Error(),
	)
	http.Error(w, nosurf.Reason(r).Error(), 403)

}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}
