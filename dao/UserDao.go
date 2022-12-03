package dao

import (
	"boboshu/config"
	"boboshu/po"
)

func SelectTUserByUserName(username string) (*po.TUser, error) {
	var user *po.TUser
	db := config.GetConnectToMysql()
	err := db.Where("username = ?", username).First(&user).Error
	return user, err
}

func CreateTUser(user *po.TUser) error {
	return config.GetConnectToMysql().Create(user).Error
}
