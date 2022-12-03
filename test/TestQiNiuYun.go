package test

import (
	"boboshu/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UpLoadFormImage(ctx *gin.Context) {
	fmt.Println("TestUpLoadImage")
	file, err := ctx.FormFile("test")
	if err != nil {
		fmt.Println(err)
	}
	fileName := file.Filename
	datafile, _ := file.Open()
	dataSize := file.Size
	config.UpLoadImage(fileName, datafile, dataSize)
}
