package config

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"time"
)

var accessKey = "vhC9w7fpcstsAUEnt8LhmLFGg5B2_J3UH_1Ss2td"
var secretKey = "scqn_3BCP-v3iwNn_ZLB-VrtqeHpSl6qbqY6MXRr"
var bucket = "boboshu"
var mac *qbox.Mac
var putPolicy storage.PutPolicy
var formUploader *storage.FormUploader
var ret storage.PutRet
var putExtra storage.PutExtra

func init() {
	mac = qbox.NewMac(accessKey, secretKey)
	putPolicy = storage.PutPolicy{
		Scope:   bucket,
		Expires: 3600 * 24 * 7,
	}
	//localFile := "/Users/jemy/Documents/github.png"
	//key := "github-x.png"
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	// 构建表单上传的对象
	formUploader = storage.NewFormUploader(&cfg)
	ret = storage.PutRet{}

	// 可选配置
	putExtra = storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	//err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(ret.Key, ret.Hash)
}

func getUpLoadToken() string {
	rdb := GetRedisConnect()
	result, err := rdb.Get("upLoadToken").Result()
	if err != nil {
		result = putPolicy.UploadToken(mac)
		rdb.Set("upLoadToken", result, time.Second*60*60*24*7)
	}
	return result
}

func UpLoadImage(key string, data io.Reader, dataSize int64) {
	err := formUploader.Put(context.Background(), &ret, getUpLoadToken(), key, data, dataSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
