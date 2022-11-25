package service

import (
	"boboshu/dao"
	"boboshu/po"
	"crypto/md5"
	"encoding/hex"
	"log"
)

func UserLogin(username, password string) bool {
	user, err := dao.SelectUserByUserName(username)
	if err != nil {
		log.Println("user do not exist" + err.Error())
	}
	if user.Password == hex.EncodeToString(md5.New().Sum([]byte(password))) {
		return true
	}
	return false
}

func UserRegister(username, password, email string) bool {
	user := &po.User{Username: username, Password: password, Email: email}
	user.Password = hex.EncodeToString(md5.New().Sum([]byte(password)))
	err := dao.CreateUser(user)
	if err != nil {
		log.Println("create user fail" + err.Error())
		return false
	}
	return true
}

func GetUserByUserName(username string) *po.User {
	user, err := dao.SelectUserByUserName(username)
	if err != nil {
		log.Println(err)
	}
	return user
}
