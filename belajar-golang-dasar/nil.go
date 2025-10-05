package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {

	result := NewMap("rida")

	if result == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(result)
	}
}
