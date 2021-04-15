package models

type Menu struct {
	MenuType    string
	MenuName    string
	TimeServing float64
	Price       float64
}

func NewMenu(menuType string, menuName string, timeServing float64, price float64) *Menu {
	return &Menu{menuType, menuName, timeServing, price}
}
