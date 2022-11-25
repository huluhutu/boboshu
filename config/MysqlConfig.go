package config

import (
	"boboshu/po"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

var username = "root"
var password = "123456"
var host = "127.0.0.1"
var port = 3306
var dbname = "boboshu"
var timeout = "10s"
var maxOpenConnect = 100
var maxIdleConnect = 20
var maxIdleTime = time.Second * 100

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connect to mysql fail : " + err.Error())
	} else {
		log.Println("connect to mysql success.......")
	}
	set, _ := db.DB()
	set.SetMaxOpenConns(maxOpenConnect)
	set.SetMaxIdleConns(maxIdleConnect)
	set.SetConnMaxIdleTime(maxIdleTime)
	DB = db

	err = db.AutoMigrate(&po.User{}, &po.UserRelative{}, &po.Navigation{}, &po.Category{}, &po.Article{})
	if err != nil {
		log.Println(err)
	}
}

func GetConnectToMysql() *gorm.DB {
	return DB
}
