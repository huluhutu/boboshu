package po

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string
}

func (u User) TableName() string {
	return "t_user"
}

type UserRelative struct {
	gorm.Model
	UserId      uint   `gorm:"unique"`
	FavoriteIds string `gorm:"unique"`
	Subscribes  string `gorm:"unique"`
	ArticleNum  int
	Header      string
	NickName    string
}

func (r UserRelative) TableName() string {
	return "t_user_relative"
}
