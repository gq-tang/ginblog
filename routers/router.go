package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gq-tang/ginblog/controllers"
	"html/template"
)

func Engine() *gin.Engine {
	r := gin.Default()
	funcMaps := template.FuncMap{
		"config":   getConfig,
		"substr":   subStr,
		"str2html": str2html,
		"date_mh":  getDateMH,
		"date":     getDate,
		"avatar":   getGravatar,
	}
	r.HTMLRender = loadMultiTemplates("../views/", "inc/", "tpl", funcMaps)
	store := cookie.NewStore([]byte("verysecret"))
	r.Static("/static", "../static")
	r.Use(sessions.Sessions("mysession", store))
	{
		r.GET("/login", controllers.Login)
	}

	return r
}
