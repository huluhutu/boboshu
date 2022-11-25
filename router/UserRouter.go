package router

import (
	"boboshu/controller"
	"boboshu/middleware"
	"github.com/gin-gonic/gin"
)

var Engine = gin.Default()

func GetEngine() *gin.Engine {
	return Engine
}

func init() {
	userGroup := Engine.Group("/user")
	{
		userGroup.POST("/login", controller.Login)
		userGroup.POST("/register", controller.Register)
	}
	warkgroup := Engine.Group("/worktable")
	{
		warkgroup.Use(middleware.UserMiddleWare())
		warkgroup.GET("/", controller.GetUser)
	}
}
