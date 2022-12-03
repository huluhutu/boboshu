package service

import (
	"boboshu/config"
	"boboshu/dao"
	"boboshu/dto"
	"boboshu/po"
	"boboshu/util"
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"time"
)

func UserLogin(username, password string) (bool, *dto.UserDto) {
	user, err := dao.SelectTUserByUserName(username)
	if err != nil {
		log.Println("user do not exist" + err.Error())
	}
	if user.Password == hex.EncodeToString(md5.New().Sum([]byte(password))) {
		token, _ := util.GenerateToken(username, time.Second*24*7*60*60)
		rdb := config.GetRedisConnect()
		rdb.Set(username, token, time.Second*24*7*60*60)
		userDto := SetUserDto(user)
		userDto.Token = token
		return true, userDto
	}
	return false, nil
}

func UserRegister(user *po.TUser) bool {
	user.CreateTime = time.Now()
	user.Password = hex.EncodeToString(md5.New().Sum([]byte(user.Password)))
	user.BlogUrl = "http://127.0.0.1:8808/" + "blog/" + user.Username
	user.FavoriteNum = 0
	user.SubscribeNum = 0
	err := dao.CreateTUser(user)
	if err != nil {
		log.Println("create user fail" + err.Error())
		return false
	}
	CreateUserRelative(user.Username)
	return true
}

func GetUserByUserName(username string) *dto.UserDto {
	user, err := dao.SelectTUserByUserName(username)
	if err != nil {
		log.Println(err)
	}
	return SetUserDto(user)
}

func CreateUserRelative(username string) {
	var user *po.TUser
	user, err := dao.SelectTUserByUserName(username)
	if err != nil {
		log.Println(err, user)
		return
	}
	mongodb := config.ConnectToMongoDB()
	cuserrelative := po.CUserRelative{
		UserId:     user.Id,
		Favorites:  &[]int{},
		Subscribes: &[]int{},
	}
	one, err := mongodb.Database("boboshu").Collection("c_user_relative").InsertOne(context.TODO(), &cuserrelative)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(one)
}

func SetUserDto(user *po.TUser) *dto.UserDto {
	return &dto.UserDto{
		UserId:      user.Id,
		Username:    user.Username,
		NickName:    user.Nickname,
		Header:      user.Header,
		FavoriteNum: user.FavoriteNum,
		SubcribeNum: user.SubscribeNum,
		Blogurl:     user.BlogUrl,
	}
}
