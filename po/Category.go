package po

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `gorm:"unique"`
	UserId       uint   `gorm:"unique"`
	NavigationId uint   `gorm:"unique"`
	ArticleNum   int
}

func (c Category) TableName() string {
	return "t_category"
}
