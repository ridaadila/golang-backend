package main

import "fmt"

func main() {

	var person map[string]string = map[string]string{
		"name":    "rida",
		"address": "surabaya",
	}

	fmt.Println(person)
	fmt.Println(person["address"])

	var data = make(map[string]string)

	data["nama"] = "ridates"
	data["usia"] = "25"
	data["ups"] = "salah"

	delete(data, "ups")
	fmt.Println(data)
}
