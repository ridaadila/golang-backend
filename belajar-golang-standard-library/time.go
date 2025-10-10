package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2024, time.October, 19, 1, 29, 1, 0, time.UTC)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2025-09-10 09:18:17"

	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println(valueTime.Local())
		fmt.Println(valueTime.Year())
		fmt.Println(valueTime.Month())
		fmt.Println(valueTime.Day())
	}

}
