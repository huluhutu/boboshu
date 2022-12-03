package po

type TCategory struct {
	ID            int `gorm:"primarykey"`
	CategoryName  string
	CategoryOrder int
	UserId        uint
}

func (c TCategory) TableName() string {
	return "t_category"
}

type Category struct {
	CategoryID   int
	CategoryName string
	Articles     []Article
}
