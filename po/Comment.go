package po

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ArticleId uint
	content   string
}
