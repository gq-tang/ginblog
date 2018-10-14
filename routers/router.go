package routers

import (
	"html/template"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gq-tang/ginblog/controllers"
	"github.com/gq-tang/ginblog/static"
	log "github.com/sirupsen/logrus"
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

	t, err := loadTemplate(funcMaps, ".tpl")
	if err != nil {
		log.WithError(err).Fatal("load template error")
	}
	r.SetHTMLTemplate(t)
	store := cookie.NewStore([]byte("verysecret"))
	store.Options(sessions.Options{
		MaxAge:   24 * 60 * 60, // 1 day
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("mysession", store), controllers.IsLogin())
	r.StaticFS("static", &assetfs.AssetFS{
		Asset:     static.Asset,
		AssetDir:  static.AssetDir,
		AssetInfo: static.AssetInfo,
		Prefix:    "",
	})
	r.Static("/uploadfile", "../uploadfile/")
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

		r.POST("/comment/add", controllers.AddComment)
		r.POST("/comment/edit/status", controllers.EditComment)

		r.GET("/album", controllers.ListAlbum)
		r.GET("/album/upload", controllers.AlbumPage)
		r.POST("/album/edit", controllers.EditAlbum)

		r.POST("/upload", controllers.Upload)
		r.POST("/uploadmulti", controllers.UploadMulti)
	}

	return r
}
