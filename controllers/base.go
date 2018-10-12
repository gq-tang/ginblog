package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gq-tang/ginblog/config"
	"github.com/gq-tang/ginblog/models"
	"github.com/gq-tang/ginblog/utils"
	log "github.com/sirupsen/logrus"
)

func Go404(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.tpl", nil)
}

// get session middleware
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

// Upload single file
func Upload(ctx *gin.Context) {
	if !ctx.GetBool("islogin") {
		ctx.JSON(http.StatusOK, gin.H{
			"error":   1,
			"message": "你没有权限上传",
		})
		return
	}

	file, err := ctx.FormFile("imgFile")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error":   1,
			"message": err.Error(),
		})
		return
	}

	// mkdir
	dir, err := mkdir("../static/uploadfile/")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error":   1,
			"message": err.Error(),
		})
		return
	}
	// generate file name
	ext := filepath.Ext(file.Filename)
	fileName := utils.GetUUID() + ext
	filepath := dir + "/" + fileName
	err = ctx.SaveUploadedFile(file, filepath)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error":   1,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": 0,
		"url":   strings.TrimPrefix(strings.TrimPrefix(filepath, ".."), "."),
	})
}

// upload multi files
func UploadMulti(ctx *gin.Context) {
	if !ctx.GetBool("islogin") {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "你没有权限上传",
		})
		return
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	// mkdir
	dir, err := mkdir("../static/uploadfile/")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}

	// uploadMultiPic := ""
	// uploadMultiName := ""
	files := form.File["uploadFiles"]
	for i, _ := range files {
		filename := files[i].Filename
		//uploadMultiName += utils.ReplaceFileSuffix(filename) + "||"

		ext := filepath.Ext(filename)
		newName := utils.GetUUID() + ext
		filepath := dir + "/" + newName
		err := ctx.SaveUploadedFile(files[i], filepath)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": err.Error(),
			})
			return
		}
		item := models.Album{
			Title:   utils.ReplaceFileSuffix(filename),
			Picture: strings.TrimPrefix(strings.TrimPrefix(filepath, ".."), "."),
			Status:  1,
			Created: time.Now().Unix(),
		}
		_, err = models.CreateAlbum(config.C.MySQL.DB, &item)
		if err != nil {
			log.Error(err)
		}
		//uploadMultiPic += strings.TrimPrefix(strings.TrimPrefix(filepath, ".."), ".") + "||"
	}
	// session := sessions.Default(ctx)
	// session.Set("uploadMultiPic", uploadMultiPic)
	// session.Set("uploadMultiName", uploadMultiName)
	// session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "上传成功",
		// "url":     uploadMultiPic,
	})
}

func mkdir(path string) (string, error) {
	now := time.Now()
	dir := path + now.Format("2006-01") + "/" + strconv.Itoa(now.Day())
	return dir, os.MkdirAll(dir, 0755)
}
