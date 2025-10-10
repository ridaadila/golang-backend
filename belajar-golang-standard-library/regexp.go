package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile(`r[a-z]d`)
	fmt.Println(regex.MatchString("rid"))
	fmt.Println(regex.MatchString("r1d"))
	fmt.Println(regex.MatchString("rad"))

	fmt.Println(regex.FindAllString("rid rud rad r4d r8d", 10))
}
