package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// upload
func GetAlbum(ctx *gin.Context) {
	session := sessions.Default(ctx)
	v := session.Get("user")
	if v == nil {
		ctx.Redirect(http.StatusPermanentRedirect, "/login")
	}
	ctx.HTML(http.StatusOK, "album-upload.tpl", nil)
}
