package controller

import (
	"boboshu/dto"
	"boboshu/po"
	"boboshu/service"
	"boboshu/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	login, userDto := service.UserLogin(username, password)
	if login {
		ctx.JSON(http.StatusOK, &dto.Result{
			Code: http.StatusOK, Message: "login success", ResBody: userDto,
		})
	} else {
		ctx.JSON(http.StatusNotFound, &dto.Result{
			Code: http.StatusNotFound, Message: "login fail",
		})
	}
}

func Register(ctx *gin.Context) {
	var user po.TUser
	user.Username = ctx.PostForm("username")
	user.Password = ctx.PostForm("password")
	user.Email = ctx.PostForm("email")
	user.Nickname = ctx.PostForm("nickname")
	user.Header = util.UpLoadHeaderAndReturnKey(ctx)
	register := service.UserRegister(&user)
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
	userDto := service.GetUserByUserName(username)
	ctx.JSON(http.StatusOK, &dto.Result{
		Code: http.StatusOK, Message: "user information", ResBody: userDto,
	})
}
