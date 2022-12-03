package po

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ArticleId   int
	ArticleName string
}

type TArticle struct {
	ID           int
	ArticleName  string
	ArticleOrder int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Like         int
	Favorite     int
	UserId       uint
}

func (a TArticle) TableName() string {
	return "t_Article"
}

type CArticle struct {
	gorm.Model
	ArticleId    int
	ArticleName  string
	ArticleOrder int
	Content      string
	Like         int
	Favorite     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
