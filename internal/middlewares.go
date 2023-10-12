package internal

import (
	"github.com/cesarcruzc/nearshore_test/pkg/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuthMiddleware(username, password string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, pass, hasAuth := ctx.Request.BasicAuth()

		if hasAuth && user == username && helper.CheckPasswordHash(pass, password) {
			ctx.Next()
			return
		}

		ctx.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
