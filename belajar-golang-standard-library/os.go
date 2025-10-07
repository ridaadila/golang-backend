package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for _, val := range args {
		fmt.Println(val)
	}

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println("Hostname: ", hostname)
	} else {
		fmt.Println(err.Error())
	}
}
