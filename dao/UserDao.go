package dao

import (
	"boboshu/config"
	"boboshu/po"
)

func SelectUserByUserName(username string) (*po.User, error) {
	var user *po.User
	db := config.GetConnectToMysql()
	err := db.Where("username = ?", username).First(&user).Error
	return user, err
}

func CreateUser(user *po.User) error {
	db := config.GetConnectToMysql()
	return db.Create(user).Error
}
