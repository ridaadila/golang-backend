package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Rida")
	data.PushBack("Adila")
	data.PushBack("Surabaya")

	first := data.Front()
	fmt.Println(first.Value)

	next := first.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
