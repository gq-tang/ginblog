package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gq-tang/ginblog/config"
	"github.com/gq-tang/ginblog/models"
	log "github.com/sirupsen/logrus"
)

type Login struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

// check session
func LoginPage(ctx *gin.Context) {
	if !ctx.GetBool("islogin") {
		ctx.HTML(http.StatusOK, "login.tpl", nil)
	} else {
		ctx.Redirect(http.StatusFound, "/article")
	}
}

// user login
func Loging(ctx *gin.Context) {
	var login Login
	err := ctx.ShouldBind(&login)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	if login.Phone == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "请输入用户ID",
		})
		return
	}
	if login.Password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "请输入密码",
		})
		return
	}

	user, err := models.LoginUser(config.C.MySQL.DB, login.Phone, login.Password)
	if err == nil {
		session := sessions.Default(ctx)
		session.Set("user", user.ID)
		session.Save()
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "登录成功",
			//"user":    user,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "登录失败:" + err.Error(),
		})
		return
	}
}

// GETLogout logut user
func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("user")
	session.Save()
	ctx.Redirect(http.StatusFound, "/login")
}

// GETAboutMe return about me
func AboutMe(ctx *gin.Context) {
	var id int64 = 1
	userid, ok := ctx.Get("userID")
	if ok {
		id = userid.(int64)
	}
	pro, err := models.GetUserProfile(config.C.MySQL.DB, id)
	if err != nil {
		log.Error(err)
		ctx.Redirect(http.StatusFound, "/404")
		return
	}
	ctx.HTML(http.StatusOK, "about.tpl", gin.H{
		"pro":     pro,
		"isLogin": ctx.GetBool("islogin"),
	})
}
