package router

import (
	"boboshu/controller"
	"boboshu/middleware"
	"boboshu/test"
	"github.com/gin-gonic/gin"
)

var Engine = gin.Default()

func GetEngine() *gin.Engine {
	return Engine
}

func init() {
	userGroup := Engine.Group("/user")
	{
		userGroup.POST("/signin", controller.Login)
		userGroup.POST("/signup", controller.Register)
		userGroup.POST("/upload", test.UpLoadFormImage)
		//userGroup.POST("/signout")
	}
	warkgroup := Engine.Group("/worktable")
	{
		warkgroup.Use(middleware.UserMiddleWare())
		warkgroup.GET("/", controller.GetUser)
	}
}
