package main

import "fmt"

func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "rida"
	lastname = "adila"

	return firstname, middlename, lastname
}

func main() {

	a, b, c := getCompleteName()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
