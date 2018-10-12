package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gq-tang/ginblog/config"
	"github.com/gq-tang/ginblog/models"
)

// AddComment add article comment.
func AddComment(ctx *gin.Context) {
	var item models.Comment
	err := ctx.ShouldBind(&item)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	item.Created = time.Now().Unix()
	id, err := models.CreateComment(config.C.MySQL.DB, &item)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "评论添加成功",
		"id":      id,
	})
}

// EditComment update comment status
func EditComment(ctx *gin.Context) {
	if !ctx.GetBool("islogin") {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": ctx.GetBool("islogin"),
		})
		return
	}

	param := struct {
		ID     int64 `form:"id"`
		Status int   `form:"status"`
	}{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	err = models.UpdateComment(config.C.MySQL.DB, param.ID, param.Status)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "状态更新成功",
	})
}
