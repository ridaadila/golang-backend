package main

import "fmt"

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	fungsi := sayGoodbye

	fmt.Println(fungsi("Rida"))
}
