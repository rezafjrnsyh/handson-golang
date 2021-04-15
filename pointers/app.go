package main

import "fmt"

type Person struct {
    firstName string
    lastName  string
}

// parse by pointer
func changeName(p *Person) {
	// jika sebuah struct tidak perlu menggunakan pointer untuk merubah
	p.firstName = "Eja"
	fmt.Println(p)
}

func changeNumberPtr(num *int) {
	// jika sebuah int perlu menggunakan pointer untuk merubah
	*num = 11
	fmt.Println(*num)
}

func changeNumber(num int) {
	num = 11
	fmt.Println(num)
}

func main() {
    // person := Person {
    //     firstName: "Reza",
    //     lastName: "Fajriansyah",
    // }

    // changeName(&person)

	// fmt.Println(person)
	
	i := 10 
	// initial
	fmt.Println(i)

	// changeNumberPtr(&i)
	changeNumber(i)
	fmt.Println(i)

}