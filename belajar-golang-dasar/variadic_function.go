package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, val := range numbers {
		total += val
	}

	return total
}

func main() {

	fmt.Println(sumAll(10, 20, 30, 40))

	data := []int{10, 10, 10}

	fmt.Println(sumAll(data...))
}
