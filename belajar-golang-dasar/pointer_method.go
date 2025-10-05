package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) ChangeName() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{"Before"}

	man.ChangeName()
	fmt.Println(man)
}
