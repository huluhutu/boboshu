package controller

import (
	"boboshu/config"
	"boboshu/dto"
	"boboshu/service"
	"boboshu/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	login := service.UserLogin(username, password)
	if login {
		token, _ := util.GenerateToken(username, time.Second*24*7*60*60)
		rdb := config.GetRedisConnect()
		rdb.Set(username, token, time.Second*24*7*60*60)
		ctx.JSON(http.StatusOK, &dto.Result{
			Code: http.StatusOK, Message: "login success", ResBody: token,
		})
	} else {
		ctx.JSON(http.StatusNotFound, &dto.Result{
			Code: http.StatusNotFound, Message: "login fail",
		})
	}
}

func Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	register := service.UserRegister(username, password, email)
	if register {
		ctx.JSON(http.StatusOK, &dto.Result{
			Code: http.StatusOK, Message: "register success",
		})
	} else {
		ctx.JSON(http.StatusPreconditionFailed, &dto.Result{
			Code: http.StatusPreconditionFailed, Message: "register fail",
		})
	}
}

func GetUser(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	user := service.GetUserByUserName(username)
	ctx.JSON(http.StatusOK, &dto.Result{
		Code: http.StatusOK, Message: "user information", ResBody: user,
	})
}
