package main

import "fmt"

func main() {

	counter := 0

	increment := func() {
		fmt.Println("counter")
		counter++
	}

	increment()
	increment()
	increment()
	increment()
	increment()
	fmt.Println("Final counter:", counter)
}
