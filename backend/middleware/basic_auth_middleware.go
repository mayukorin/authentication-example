package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func basicAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if ok == false {
		return false
	}
	return username == "mayukorin" && password == "password"
}

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if basicAuth(c.Request) == false {
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm=hello with basic auth")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}
		c.Next()
	}
}
