package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Angel", "David", "Lucas"}
	values := []int{321, 0, 89, 90, 81}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Rida"))
	fmt.Println(slices.Index(values, 90))
}
