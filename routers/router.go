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
	r.Use(sessions.Sessions("mysession", store), controllers.IsLogin())
	r.Static("/static", "../static")
	{
		r.GET("/", controllers.ListArticle)
		r.GET("/404", controllers.Go404)

		r.GET("/login", controllers.LoginPage)
		r.POST("/login", controllers.Loging)
		r.GET("/logout", controllers.Logout)

		r.GET("/about", controllers.AboutMe)

		r.GET("/article", controllers.ListArticle)
		r.GET("/article/detail/:id", controllers.ArticleDetail)
		r.GET("/article/add", controllers.AddArticlePage)
		r.POST("/article/add", controllers.AddArticle)
		r.GET("/article/edit/:id", controllers.GetEditArticle)
		r.POST("/article/edit/:id", controllers.EditArticle)
	}

	return r
}
