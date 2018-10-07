package controllers

import (
	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// v := session.Get("user")
	// if v == nil {
	ctx.HTML(http.StatusOK, "login.tpl", nil)
	// } else {
	// 	ctx.Redirect(http.StatusFound, "/article")
	// }
}

func Test(ctx *gin.Context) {
	ctx.HTML(200, "test.tpl", nil)
}
