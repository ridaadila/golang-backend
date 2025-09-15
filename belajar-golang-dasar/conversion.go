package main

import "fmt"

func main() {
	var angka32 int32 = 32768
	var angka64 int64 = int64(angka32)
	var angka16 int16 = int16(angka32)

	fmt.Println(angka32)
	fmt.Println(angka64)
	fmt.Println(angka16)

	var name string = "Rida"
	var r = name[0]
	var isiR = string(r)
	fmt.Println(name)
	fmt.Println(r)
	fmt.Println(isiR)
}
