package main

import "fmt"

func main() {
	var data [3]string

	data[0] = "rida"
	data[1] = "adila"
	data[2] = "tes"

	fmt.Println(data[0])
	fmt.Println(data[1])
	fmt.Println(data[2])

	var angka = [...]int{
		1,
		2,
		3,
		4,
	}

	fmt.Println(len(angka))
	fmt.Println(angka[0])
	angka[0] = 100
	fmt.Println(angka[0])

}
