package controllers

import (
	"net/http"
	"strconv"

	"ginblog/config"
	"ginblog/models"
	"ginblog/utils/pagination"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// get upload album page
func AlbumPage(ctx *gin.Context) {
	if !ctx.GetBool("islogin") {
		ctx.Redirect(http.StatusPermanentRedirect, "/login")
		return
	}
	ctx.HTML(http.StatusOK, "album-upload.tpl", gin.H{
		"isLogin": ctx.GetBool("islogin"),
	})
}

// EditAlbum update album info
func EditAlbum(ctx *gin.Context) {
	var item models.Album
	err := ctx.ShouldBind(&item)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "相册修改出错:" + err.Error(),
		})
		return
	}

	err = models.UpdateAlbum(config.C.MySQL.DB, &item)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "相册修改出错:" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "相册修改成功",
	})
}

// ListAlbum list albums
func ListAlbum(ctx *gin.Context) {
	pagestr := ctx.Param("p")
	title := ctx.Param("title")
	keywords := ctx.Param("keywords")
	status := ctx.Param("status")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		page = 1
	}

	offset, err := config.C.Int("pageoffset")
	if err != nil {
		offset = 9
	}

	if !ctx.GetBool("islogin") {
		status = "1"
	}
	count, _ := models.CountAlbum(config.C.MySQL.DB, title, keywords, status)
	paginator := pagination.NewPaginator(ctx.Request, offset, count)
	albs, err := models.ListAlbum(config.C.MySQL.DB, page, offset, title, keywords, status)
	if err != nil {
		log.Error(err)
	}
	ctx.HTML(http.StatusOK, "album.tpl", gin.H{
		"paginator": paginator,
		"alb":       albs,
		"isLogin":   ctx.GetBool("islogin"),
	})
}
