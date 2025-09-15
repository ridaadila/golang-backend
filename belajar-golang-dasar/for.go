package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke: ", i)
	}

	var names = [...]string{
		"rida",
		"adila",
		"tes",
		"windy",
	}

	for index, value := range names {
		fmt.Println("Index ke: ", index, ", value: ", value)
	}
}
