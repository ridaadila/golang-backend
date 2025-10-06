package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi sama dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(1, 10)

	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error: ", err.Error())
	}
}
