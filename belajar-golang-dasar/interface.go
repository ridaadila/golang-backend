package main

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	person := Person{Name: "Rida"}

	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
