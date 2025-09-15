package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRida NoKTP = "111111111"
	var nomor string = "22222222"
	var nomorKtp NoKTP = NoKTP(nomor)

	fmt.Println(ktpRida)
	fmt.Println(nomor)
	fmt.Println(nomorKtp)
}
