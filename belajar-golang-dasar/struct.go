package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var rida Customer
	rida.Name = "Rida"
	rida.Address = "Indonesia"
	rida.Age = 20

	fmt.Println(rida)
	fmt.Println(rida.Name)
	fmt.Println(rida.Address)
	fmt.Println(rida.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)

	joko.sayHello("Budi")
}
