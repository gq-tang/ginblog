package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Go404(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.tpl", nil)
}

// get session
func IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var isLogin bool
		session := sessions.Default(ctx)
		user := session.Get("user")

		if user != nil {
			isLogin = true
		}
		ctx.Set("islogin", isLogin)
		ctx.Next()
	}
}
