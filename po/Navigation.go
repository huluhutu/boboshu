package po

import "gorm.io/gorm"

type Navigation struct {
	gorm.Model
	NavigationName string `gorm:"unique"`
	UserId         string
}

func (n Navigation) TableName() string {
	return "t_navigation"
}
