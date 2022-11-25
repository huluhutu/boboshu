package middleware

import (
	"boboshu/config"
	"boboshu/dto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		username := ctx.GetHeader("username")
		result, err := config.GetRedisConnect().Get(username).Result()
		if err == nil && token == result {
			log.Print("continue to travel")
		} else {
			log.Print(result + " : " + token)
			ctx.JSON(http.StatusExpectationFailed, &dto.Result{
				Code: http.StatusExpectationFailed, Message: "you do not login",
			})
			ctx.Abort()
		}
	}
}
