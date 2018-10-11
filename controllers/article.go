package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gq-tang/ginblog/utils/pagination"

	"github.com/gin-gonic/gin"
	"github.com/gq-tang/ginblog/config"
	"github.com/gq-tang/ginblog/models"
)

// get add article page.
func AddArticlePage(ctx *gin.Context) {
	if ctx.GetBool("islogin") {
		art := models.Article{Status: 1}
		ctx.HTML(http.StatusOK, "article-form.tpl", gin.H{
			"art": art,
		})
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}
}

// create article
func AddArticle(ctx *gin.Context) {
	if !ctx.GetBool("islogin") {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "请先登录",
		})
		return
	}
	var art models.Article
	err := ctx.ShouldBind(&art)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	art.Created = time.Now().Unix()
	id, err := models.CreateArticle(config.C.MySQL.DB, &art)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "博客添加成功",
		"id":      id,
	})
}

func GetEditArticle(ctx *gin.Context) {
	if !ctx.GetBool("islogin") {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "请先登录",
		})
		return
	}
	idstr := ctx.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 0)

	art, err := models.GetArticle(config.C.MySQL.DB, id)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/404")
		return
	}
	ctx.HTML(http.StatusOK, "article-form.tpl", gin.H{
		"art": art,
	})
}

// edit article
func EditArticle(ctx *gin.Context) {
	var art models.Article
	err := ctx.ShouldBind(&art)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	if art.ID <= 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "博客不存在",
		})
		return
	}
	err = models.UpdateArticle(config.C.MySQL.DB, art)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "博客修改成功",
		"id":      art.ID,
	})
}

// list article
func ListArticle(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("p", "1")
	title := ctx.Query("title")
	keywords := ctx.Query("keywords")
	status := ctx.Query("status")
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}

	if !ctx.GetBool("islogin") {
		status = "1"
	} else {
		if status == "" {
			status = "status"
		}
	}
	offset, err := config.C.Int("pageoffset")
	if err != nil {
		offset = 9
	}

	count, err := models.CountArticle(config.C.MySQL.DB, status, title, keywords)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}

	paginator := pagination.NewPaginator(ctx.Request, offset, count)
	arts, err := models.ListArticle(config.C.MySQL.DB, page, offset, status, title, keywords)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "article.tpl", gin.H{
		"art":       arts,
		"paginator": paginator,
		"isLogin":   ctx.GetBool("islogin"),
	})
}

// Get article detail
func ArticleDetail(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, _ := strconv.ParseInt(idstr, 10, 0)

	art, err := models.GetArticle(config.C.MySQL.DB, id)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/404")
		return
	}
	if !ctx.GetBool("islogin") {
		if art.Status == 0 {
			ctx.Redirect(http.StatusFound, "/404")
		}
		return
	}

	// 评论分页
	pageStr := ctx.Query("p")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	offset, err := config.C.Int("pageoffset")
	if err != nil {
		offset = 9
	}
	status := "status"
	if !ctx.GetBool("islogin") {
		status = "1"
	}
	count, _ := models.CountComment(config.C.MySQL.DB, id, status)

	paginator := pagination.NewPaginator(ctx.Request, offset, count)
	items, _ := models.ListComment(config.C.MySQL.DB, page, offset, id, status)

	ctx.HTML(http.StatusOK, "article-detail.tpl", gin.H{
		"coms":      items,
		"paginator": paginator,
		"art":       art,
	})
}
