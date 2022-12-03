package po

import (
	"time"
)

type TUser struct {
	Id           int `gorm:"primarykey;autoIncrement"`
	Username     string
	Password     string
	Email        string
	Nickname     string
	Header       string
	Status       int
	CreateTime   time.Time
	BlogUrl      string
	FavoriteNum  int //articleId
	SubscribeNum int // userId
}

func (u TUser) TableName() string {
	return "t_user"
}

type CUserRelative struct {
	UserId     int
	Favorites  *[]int //articleId
	Subscribes *[]int // userId
}
