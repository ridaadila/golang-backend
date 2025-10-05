package main

import "fmt"

type Address struct {
	City, Province string
}

func main() {

	var address1 Address = Address{"Surabaya", "Jatim"}
	var address2 *Address = &address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}
