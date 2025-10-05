package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.City = "Indonesia"
}

func main() {

	address := Address{}
	ChangeCountryToIndonesia(&address)
	fmt.Println(address)
}
