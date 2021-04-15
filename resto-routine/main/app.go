package main

import (
	"fmt"
	"restoroutine/models"
	"restoroutine/usecases"
	"strings"
	"sync"
	"time"
)

var restoMenus *usecases.MainMenu
var bills []*models.Bill
var wg sync.WaitGroup

func main() {
	restoMenus = usecases.RestoMainMenu()
	order("0001", "0002", "0001")
}

func order(orders ...string) {
	wg.Add(len(orders))
	for _, plu := range orders {
		go production(plu)
	}
	wg.Wait()
	printBill()
}

func production(plu string) {
	orderCheckIn := time.Now()
	defer wg.Done()

	menu, timeServing, price := restoMenus.GetMenuByPLU(plu)
	time.Sleep(time.Second * time.Duration(timeServing))
	orderCheckout := time.Now()

	checkInFmt := orderCheckIn.Format("Jan _2 15:04:05.000000")
	checkOutFmt := orderCheckout.Format("Jan _2 15:04:05.000000")
	duration := orderCheckout.Sub(orderCheckIn)
	fmt.Printf("Serve %v, ci: %v, co: %v, dur:%v\n", menu, checkInFmt, checkOutFmt, duration)
	bills = append(bills, models.NewBill(menu, price))
}

func printBill() {
	fmt.Printf("%-3v %-20v %10v\n", "No", "Menu", "Price")
	fmt.Printf("%s\n", strings.Repeat("=", 40))
	var total float64
	for idx, billItem := range bills {
		menuItem, price := billItem.BillDetail()
		total += price
		fmt.Printf("%-3v %-20v %10v\n", idx+1, menuItem, price)
	}
	fmt.Printf("%s\n", strings.Repeat("=", 40))
	fmt.Printf("%-3v %-20v %10v\n", "", "Total", total)
}
