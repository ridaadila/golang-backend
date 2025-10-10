package main

import (
	"fmt"
	"strconv"
)

func main() {

	res, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println(res)
	}

	val, err := strconv.Atoi("100")

	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println(val)
	}

	stringVal := strconv.Itoa(250)
	fmt.Println("String val: ", stringVal)
}
