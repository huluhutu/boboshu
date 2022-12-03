package po

type Navigation struct {
	ID              int
	NavigationName  string
	NavigationOrder int
	Categorys       []Category
}

type TNavigation struct {
	ID              int `gorm:"primarykey"`
	NavigationName  string
	NavigationOrder int
	UserId          string
}

func (n TNavigation) TableName() string {
	return "t_navigation"
}

type CNavigations struct {
	UserId      int
	Navigations []Navigation
}
