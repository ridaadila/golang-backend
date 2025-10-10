package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rida", "id"))
	fmt.Println(strings.ToUpper("rida adila"))
	fmt.Println(strings.ToLower("RIDA ADILA"))
	fmt.Println(strings.Trim("      rida tes adila   ", " "))
	fmt.Println(strings.ReplaceAll("rida adila tes tes nama", "tes", ""))
	fmt.Println(strings.Split("Rida Adila", " "))
}
