package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Go404(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.tpl", nil)
}
