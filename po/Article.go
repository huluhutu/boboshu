package po

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	CategoryName string
	ArticleName  string
	Content      string
	Like         int
	Favorite     int
}

func (a Article) TableName() string {
	return "t_Article"
}
