package main

import "fmt"

func main() {

	nama := "rida"

	switch nama {
	case "rida":
		fmt.Println("Benar")
	case "adila":
		fmt.Println("Salah")
	default:
		fmt.Println("out")
	}
}
