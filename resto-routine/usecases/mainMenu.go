package usecases

import "restoroutine/models"

type MainMenu map[string]*models.Menu

func RestoMainMenu() *MainMenu {
	return &MainMenu{
		"0001": models.NewMenu("Main Course", "Pikja itali", 2.3, 125000),
		"0002": models.NewMenu("Drink", "Kopi Tabrak Lari", 1.5, 35000),
		"0003": models.NewMenu("Drink", "Aer putih rasa hujan", 0.3, 1000),
		"0004": models.NewMenu("Dessert", "Sang Pisang", 1.2, 24000),
		"0005": models.NewMenu("Main Course", "Cumi Goreng Tepung", 3.5, 60000),
	}
}

func (mm *MainMenu) GetAllMenus() []*models.Menu {
	var menus []*models.Menu

	for _, menuInfo := range *mm {
		menus = append(menus, menuInfo)
	}
	return menus
}

func (mm *MainMenu) GetMenuByPLU(plu string) (menuName string, timeServing float64, price float64) {
	menuName = (*mm)[plu].MenuName
	timeServing = (*mm)[plu].TimeServing
	price = (*mm)[plu].Price
	return
}
