package routers

import (
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gq-tang/ginblog/controllers"
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
	r.SetFuncMap(funcMaps)
	files := getTempalteFiles("../views/", "tpl")
	r.LoadHTMLFiles(files...)
	store := cookie.NewStore([]byte("verysecret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Static("/static", "../static")
	{
		r.GET("/404", controllers.Go404)

		r.GET("/login", controllers.GETLogin)
		r.POST("/login", controllers.POSTLogin)
		r.GET("/logout", controllers.GETLogout)

		r.GET("/about", controllers.GETAboutMe)
	}

	return r
}
