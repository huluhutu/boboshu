package util

import (
	"boboshu/config"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
)

//const (
//	defaultHeader1 = "domain.com/defaultHeader1"
//	defaultHeader2 = "domain.com/defaultHeader2"
//	defaultHeader3 = "domain.com/defaultHeader3"
//	defaultHeader4 = "domain.com/defaultHeader4"
//	defaultHeader5 = "domain.com/defaultHeader5"
//)

func UpLoadHeaderAndReturnKey(ctx *gin.Context) string {
	defaultHeader := ctx.PostForm("defaultHeader")
	if len(defaultHeader) > 0 {
		return defaultHeader
	}
	file, err := ctx.FormFile("headerImg")
	if err != nil {
		fmt.Println(err)
	}
	uuid := uuid2.NewV4().String()
	key := file.Filename + uuid
	data, _ := file.Open()
	dataSize := file.Size
	config.UpLoadImage(key, data, dataSize)
	return key
}
