package models

type Bill struct {
	menuName string
	price    float64
}

func NewBill(menuName string, price float64) *Bill {
	return &Bill{
		menuName,
		price,
	}
}

func (b *Bill) BillDetail() (string, float64) {
	return b.menuName, b.price
}
