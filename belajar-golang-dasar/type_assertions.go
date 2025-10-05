package main

import "fmt"

func random() any {
	return "rida"
}

func main() {
	result := random()
	resultString := result.(string)
	// resultInt := result.(int)

	fmt.Println(resultString)
	// fmt.Println(resultInt)

	switch res := result.(type) {
	case string:
		fmt.Println("String: ", res)
	case int:
		fmt.Println("Int: ", res)
	default:
		fmt.Println("Unknown type")
	}
}
